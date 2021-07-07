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

func testCreateUDPRoute(t *testing.T) {
	tr := &specsv1alpha4.UDPRoute{}
	copier.Copy(tr, &v4UDPRoute)

	ctx := context.Background()
	err := k8sClient.Create(ctx, tr)
	require.NoError(t, err)

	// check that the create method on the API was called
	helpers.AssertCalledEventually(t, mockAPI, "UpsertUDPRoute", timeout, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func testDeleteUDPRoute(t *testing.T) {
	tr := &specsv1alpha4.UDPRoute{}
	copier.Copy(tr, &v4UDPRoute)

	ctx := context.Background()
	err := k8sClient.Delete(ctx, tr)
	require.NoError(t, err)

	// check that the create method on the API was called
	helpers.AssertCalledEventually(t, mockAPI, "DeleteUDPRoute", timeout, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

var v4UDPRoute = &v1alpha4.UDPRoute{
	TypeMeta: v1.TypeMeta{
		Kind:       "UDPRoute",
		APIVersion: v1alpha4.GroupVersion.Identifier(),
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v4udproute",
		Namespace: "default",
	},
	Spec: v1alpha4.UDPRouteSpec{
		Matches: v1alpha4.UDPMatch{
			Name:  "testing",
			Ports: []int{9090, 8080},
		},
	},
}
