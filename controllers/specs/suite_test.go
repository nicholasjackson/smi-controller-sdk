/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package specs

import (
	"path/filepath"
	"testing"

	"github.com/servicemeshinterface/smi-controller-sdk/controllers/helpers"
	"github.com/servicemeshinterface/smi-controller-sdk/sdk"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"k8s.io/client-go/rest"
	"k8s.io/kubectl/pkg/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"

	ctrl "sigs.k8s.io/controller-runtime"

	specsv1alpha1 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha1"
	specsv1alpha2 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha2"
	specsv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha3"
	specsv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha4"
	//+kubebuilder:scaffold:imports
)

// These tests use Ginkgo (BDD-style Go testing framework). Refer to
// http://onsi.github.io/ginkgo/ to learn more about Ginkgo.

var cfg *rest.Config
var k8sClient client.Client
var testEnv *envtest.Environment
var mockAPI *helpers.MockAPI

func TestAPIs(t *testing.T) {
	t.Cleanup(func() {
		err := testEnv.Stop()
		require.NoError(t, err)
	})

	setupSuite(t)

	// execute tests
	t.Run("Create HTTPRouteGroup", testCreateHTTPRouteGroup)
	t.Run("Delete HTTPRouteGroup", testDeleteHTTProuteGroup)

	t.Run("Create TCPRoute", testCreateTCPRoute)
	t.Run("Delete TCPRoute", testDeleteTCPRoute)

	t.Run("Create UDPRoute", testCreateUDPRoute)
	t.Run("Delete UDPRoute", testDeleteUDPRoute)
}

func setupSuite(t *testing.T) {
	testEnv = &envtest.Environment{
		CRDDirectoryPaths:        []string{filepath.Join("..", "..", "config", "crd", "bases")},
		ErrorIfCRDPathMissing:    true,
		AttachControlPlaneOutput: false,
	}

	var err error
	cfg, err = testEnv.Start()
	require.NoError(t, err)
	require.NotNil(t, cfg)

	err = specsv1alpha1.AddToScheme(scheme.Scheme)
	require.NoError(t, err)

	err = specsv1alpha2.AddToScheme(scheme.Scheme)
	require.NoError(t, err)

	err = specsv1alpha3.AddToScheme(scheme.Scheme)
	require.NoError(t, err)

	err = specsv1alpha4.AddToScheme(scheme.Scheme)
	require.NoError(t, err)

	// +kubebuilder:scaffold:scheme

	k8sManager, err := ctrl.NewManager(cfg, ctrl.Options{
		Scheme: scheme.Scheme,
	})
	require.NoError(t, err)

	err = (&HTTPRouteGroupReconciler{
		Client: k8sManager.GetClient(),
	}).SetupWithManager(k8sManager)
	require.NoError(t, err)

	err = (&TCPRouteReconciler{
		Client: k8sManager.GetClient(),
	}).SetupWithManager(k8sManager)
	require.NoError(t, err)

	err = (&UDPRouteReconciler{
		Client: k8sManager.GetClient(),
	}).SetupWithManager(k8sManager)
	require.NoError(t, err)

	go func() {
		err = k8sManager.Start(ctrl.SetupSignalHandler())
		require.NoError(t, err)
	}()

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	require.NoError(t, err)
	require.NotNil(t, k8sClient)

	// create the mocks and register it with the SDK
	mockAPI = &helpers.MockAPI{}
	sdk.API().RegisterV1Alpha(mockAPI)

	mockAPI.On("UpsertHTTPRouteGroup", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(ctrl.Result{}, nil)
	mockAPI.On("DeleteHTTPRouteGroup", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(ctrl.Result{}, nil)

	mockAPI.On("UpsertTCPRoute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(ctrl.Result{}, nil)
	mockAPI.On("DeleteTCPRoute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(ctrl.Result{}, nil)

	mockAPI.On("UpsertUDPRoute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(ctrl.Result{}, nil)
	mockAPI.On("DeleteUDPRoute", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(ctrl.Result{}, nil)
}
