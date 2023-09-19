// Copyright 2022 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package health_test

import (
	"context"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"

	"github.com/gardener/gardener/pkg/logger"
	"github.com/gardener/gardener/pkg/resourcemanager/apis/config"
	resourcemanagerclient "github.com/gardener/gardener/pkg/resourcemanager/client"
	"github.com/gardener/gardener/pkg/resourcemanager/controller/health/health"
	"github.com/gardener/gardener/pkg/resourcemanager/controller/health/progressing"
	resourcemanagerpredicate "github.com/gardener/gardener/pkg/resourcemanager/predicate"
	. "github.com/gardener/gardener/pkg/utils/test/matchers"
	thirdpartyapiutil "github.com/gardener/gardener/third_party/controller-runtime/pkg/apiutil"
)

func TestHealth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Integration ResourceManager Health Suite")
}

const testID = "health-controller-test"

var (
	ctx = context.Background()
	log logr.Logger

	restConfig *rest.Config
	testEnv    *envtest.Environment
	testClient client.Client

	testNamespace *corev1.Namespace
)

var _ = BeforeSuite(func() {
	logf.SetLogger(logger.MustNewZapLogger(logger.DebugLevel, logger.FormatJSON, zap.WriteTo(GinkgoWriter)))
	log = logf.Log.WithName(testID)

	By("Start test environment")
	testEnv = &envtest.Environment{
		CRDInstallOptions: envtest.CRDInstallOptions{
			Paths: []string{filepath.Join("..", "..", "..", "..", "example", "resource-manager", "10-crd-resources.gardener.cloud_managedresources.yaml")},
		},
		ErrorIfCRDPathMissing: true,
	}

	var err error
	restConfig, err = testEnv.Start()
	Expect(err).NotTo(HaveOccurred())
	Expect(restConfig).NotTo(BeNil())

	DeferCleanup(func() {
		By("Stop test environment")
		Expect(testEnv.Stop()).To(Succeed())
	})

	By("Create test client")
	testClient, err = client.New(restConfig, client.Options{Scheme: resourcemanagerclient.CombinedScheme})
	Expect(err).NotTo(HaveOccurred())

	By("Create test Namespace")
	testNamespace = &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			// create dedicated namespace for each test run, so that we can run multiple tests concurrently for stress tests
			GenerateName: testID + "-",
		},
	}
	Expect(testClient.Create(ctx, testNamespace)).To(Succeed())
	log.Info("Created Namespace for test", "namespaceName", testNamespace.Name)

	DeferCleanup(func() {
		By("Delete test Namespace")
		Expect(testClient.Delete(ctx, testNamespace)).To(Or(Succeed(), BeNotFoundError()))
	})

	By("Setup manager")

	mgr, err := manager.New(restConfig, manager.Options{
		Scheme:             resourcemanagerclient.CombinedScheme,
		MetricsBindAddress: "0",
		Namespace:          testNamespace.Name,
		MapperProvider: func(config *rest.Config, httpClient *http.Client) (meta.RESTMapper, error) {
			return thirdpartyapiutil.NewDynamicRESTMapper(config)
		},
	})
	Expect(err).NotTo(HaveOccurred())

	By("Register controllers")
	cfg := config.HealthControllerConfig{
		ConcurrentSyncs: pointer.Int(5),
		SyncPeriod:      &metav1.Duration{Duration: 500 * time.Millisecond},
	}
	classFilter := resourcemanagerpredicate.NewClassFilter("")

	Expect((&health.Reconciler{
		Config:      cfg,
		ClassFilter: classFilter,
	}).AddToManager(ctx, mgr, mgr, mgr, "")).To(Succeed())
	Expect((&progressing.Reconciler{
		Config:      cfg,
		ClassFilter: classFilter,
	}).AddToManager(ctx, mgr, mgr, mgr, "")).To(Succeed())

	By("Start manager")
	mgrContext, mgrCancel := context.WithCancel(ctx)

	go func() {
		defer GinkgoRecover()
		Expect(mgr.Start(mgrContext)).To(Succeed())
	}()

	DeferCleanup(func() {
		By("Stop manager")
		mgrCancel()
	})
})
