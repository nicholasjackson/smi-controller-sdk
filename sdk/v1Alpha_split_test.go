package sdk

import (
	"testing"

	splitv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	"github.com/stretchr/testify/mock"
)

func TestCallsUserDefinedUpsertTrafficSplitWhenSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)

	a.V1Alpha().UpsertTrafficSplit(nil, nil, l, &splitv1alpha3.TrafficSplit{})

	v1.AssertCalled(t, "UpsertTrafficSplit", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}

func TestDoesNotCallsUserDefinedUpsertTrafficSplitWhenNotSet(t *testing.T) {
	v1, a, l := setupSDKTests(t)
	a.RegisterV1Alpha(nil)

	a.V1Alpha().UpsertTrafficSplit(nil, nil, l, &splitv1alpha3.TrafficSplit{})

	v1.AssertNotCalled(t, "UpsertTrafficSplit", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
}
