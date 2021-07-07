package sdk

import (
	"testing"

	specsv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha4"
	"github.com/stretchr/testify/mock"
)

func TestCallsUserDefinedUpsertHTTPRouteGroupWhenSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)

	a.V1Alpha().UpsertHTTPRouteGroup(nil, nil, l, &specsv1alpha4.HTTPRouteGroup{})

	v1.AssertCalled(t, "UpsertHTTPRouteGroup", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestDoesNotCallsUserDefinedUpsertHTTPRouteGroupWhenNotSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)
	a.RegisterV1Alpha(nil)

	a.V1Alpha().UpsertHTTPRouteGroup(nil, nil, l, &specsv1alpha4.HTTPRouteGroup{})

	v1.AssertNotCalled(t, "UpsertHTTPRouteGroup", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestCallsUserDefinedDeleteHTTPRouteGroupWhenSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)

	a.V1Alpha().DeleteHTTPRouteGroup(nil, nil, l, &specsv1alpha4.HTTPRouteGroup{})

	v1.AssertCalled(t, "DeleteHTTPRouteGroup", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestDoesNotCallsUserDefinedDeleteHTTPRouteGroupWhenNotSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)
	a.RegisterV1Alpha(nil)

	a.V1Alpha().DeleteHTTPRouteGroup(nil, nil, l, &specsv1alpha4.HTTPRouteGroup{})

	v1.AssertNotCalled(t, "DeleteHTTPRouteGroup", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}
