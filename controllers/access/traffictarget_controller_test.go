package access

import (
	"context"
	"testing"
	"time"

	"github.com/jinzhu/copier"
	"github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha3"
	"github.com/servicemeshinterface/smi-controller-sdk/controllers/helpers"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	timeout  = time.Second * 10
	duration = time.Second * 10
	interval = time.Millisecond * 250
)

func testCreateTrafficTarget(t *testing.T) {
	tt := &v1alpha3.TrafficTarget{}
	copier.Copy(tt, &v3TrafficTarget)

	ctx := context.Background()
	err := k8sClient.Create(ctx, tt)
	require.NoError(t, err)

	// check that the create method on the API was called
	helpers.AssertCalledEventually(t, mockAPI, "UpsertTrafficTarget", timeout, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func testDeleteTrafficTarget(t *testing.T) {
	tt := &v1alpha3.TrafficTarget{}
	copier.Copy(tt, &v3TrafficTarget)

	ctx := context.Background()
	err := k8sClient.Delete(ctx, tt)
	require.NoError(t, err)

	// check that the create method on the API was called
	helpers.AssertCalledEventually(t, mockAPI, "DeleteTrafficTarget", timeout, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

var v3TrafficTarget = &v1alpha3.TrafficTarget{
	TypeMeta: v1.TypeMeta{
		Kind:       "TrafficTarget",
		APIVersion: v1alpha3.GroupVersion.Identifier(),
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "v3access",
		Namespace: "default",
	},
	Spec: v1alpha3.TrafficTargetSpec{
		Destination: v1alpha3.IdentityBindingSubject{
			Kind:      "ServiceAccount",
			Name:      "myservice",
			Namespace: "default",
		},
		Sources: []v1alpha3.IdentityBindingSubject{
			v1alpha3.IdentityBindingSubject{
				Kind:      "ServiceAccount",
				Name:      "mydestination1",
				Namespace: "default",
			},
			v1alpha3.IdentityBindingSubject{
				Kind:      "ServiceAccount",
				Name:      "mydestination2",
				Namespace: "default",
			},
		},
		Rules: []v1alpha3.TrafficTargetRule{
			v1alpha3.TrafficTargetRule{
				Kind:    "HTTPRouteGroup",
				Name:    "myname1",
				Matches: []string{"abc", "123"},
			},
			v1alpha3.TrafficTargetRule{
				Kind:    "HTTPRouteGroup",
				Name:    "myname2",
				Matches: []string{"123", "abc"},
			},
		},
	},
}
