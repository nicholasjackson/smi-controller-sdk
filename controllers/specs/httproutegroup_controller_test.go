package specs

import (
	"context"
	"testing"
	"time"

	"github.com/jinzhu/copier"
	"github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha4"
	specsv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha4"
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

func testCreateHTTPRouteGroup(t *testing.T) {
	hrg := &specsv1alpha4.HTTPRouteGroup{}
	copier.Copy(hrg, &v4HTTPRoute)

	ctx := context.Background()
	err := k8sClient.Create(ctx, hrg)
	require.NoError(t, err)

	// check that the create method on the API was called
	helpers.AssertCalledEventually(t, mockAPI, "UpsertHTTPRouteGroup", timeout, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func testDeleteHTTProuteGroup(t *testing.T) {
	hrg := &specsv1alpha4.HTTPRouteGroup{}
	copier.Copy(hrg, &v4HTTPRoute)

	ctx := context.Background()
	err := k8sClient.Delete(ctx, hrg)
	require.NoError(t, err)

	// check that the create method on the API was called
	helpers.AssertCalledEventually(t, mockAPI, "DeleteHTTPRouteGroup", timeout, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

var v4HTTPRoute = v1alpha4.HTTPRouteGroup{
	TypeMeta: v1.TypeMeta{
		Kind:       "HTTPRouteGroup",
		APIVersion: v1alpha4.GroupVersion.Identifier(),
	},
	ObjectMeta: v1.ObjectMeta{
		Name:      "httproutegroup",
		Namespace: "default",
	},
	Spec: v1alpha4.HTTPRouteGroupSpec{
		Matches: []v1alpha4.HTTPMatch{
			v1alpha4.HTTPMatch{
				Name:      "testing1",
				Methods:   []string{"GET", "POST"},
				PathRegex: ".*",
				Headers:   map[string]string{"Foo": "Bar", "One": "Two"},
			},
			v1alpha4.HTTPMatch{
				Name:      "testing2",
				Methods:   []string{"DELETE", "POST"},
				PathRegex: "/post",
				Headers:   map[string]string{"abc": "123", "Mario": "Luigi"},
			},
		},
	},
}
