package sdk

import (
	"testing"

	"github.com/go-logr/logr"
	"github.com/stretchr/testify/mock"
	ctrl "sigs.k8s.io/controller-runtime"
)

func setupSDKTests(t *testing.T) (*V1AlphaMock, *api, logr.Logger) {
	v1 := &V1AlphaMock{}
	v1.On("UpsertTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("DeleteTrafficTarget", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("UpsertTrafficSplit", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("DeleteTrafficSplit", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("UpsertHTTPRouteGroup", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("DeleteHTTPRouteGroup", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("UpsertTCPRoute", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("DeleteTCPRoute", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("UpsertUDPRoute", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
	v1.On("DeleteUDPRoute", mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	a := &api{&v1AlphaImpl{}}
	a.RegisterV1Alpha(v1)

	l := ctrl.Log.WithName("controllers").WithName("TrafficTarget")

	return v1, a, l
}
