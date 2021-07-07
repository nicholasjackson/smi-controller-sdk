package sdk

import (
	"testing"

	accessv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha3"
	"github.com/stretchr/testify/mock"
)

func TestCallsUserDefinedUpsertTrafficTargetWhenSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)

	a.V1Alpha().UpsertTrafficTarget(nil, nil, l, &accessv1alpha3.TrafficTarget{})

	v1.AssertCalled(t, "UpsertTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestDoesNotCallsUserDefinedUpsertTrafficTargetWhenNotSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)
	a.RegisterV1Alpha(nil)

	a.V1Alpha().UpsertTrafficTarget(nil, nil, l, &accessv1alpha3.TrafficTarget{})

	v1.AssertNotCalled(t, "UpsertTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}
