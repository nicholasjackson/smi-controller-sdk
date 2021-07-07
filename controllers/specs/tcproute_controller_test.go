package specs

import (
	"context"
	"testing"

	"github.com/jinzhu/copier"
	"github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha4"
	specsv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha4"
	"github.com/servicemeshinterface/smi-controller-sdk/controllers/helpers"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func testCreateTCPRoute(t *testing.T) {
	tr := &specsv1alpha4.TCPRoute{}
	copier.Copy(tr, &v4TCPRoute)

	ctx := context.Background()
	err := k8sClient.Create(ctx, tr)
	require.NoError(t, err)

	// check that the create method on the API was called
	helpers.AssertCalledEventually(t, mockAPI, "UpsertTCPRoute", timeout, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func testDeleteTCPRoute(t *testing.T) {
	tr := &specsv1alpha4.TCPRoute{}
	copier.Copy(tr, &v4TCPRoute)

	ctx := context.Background()
	err := k8sClient.Delete(ctx, tr)
	require.NoError(t, err)

	// check that the create method on the API was called
	helpers.AssertCalledEventually(t, mockAPI, "DeleteTCPRoute", timeout, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

var v4TCPRoute = &v1alpha4.TCPRoute{
	TypeMeta: v1.TypeMeta{
		Kind:       "TCPRoute",
		APIVersion: v1alpha4.GroupVersion.Identifier(),
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v4tcproute",
		Namespace: "default",
	},
	Spec: v1alpha4.TCPRouteSpec{
		Matches: v1alpha4.TCPMatch{
			Name:  "testing",
			Ports: []int{9090, 8080},
		},
	},
}
