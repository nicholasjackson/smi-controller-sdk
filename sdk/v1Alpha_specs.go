package sdk

import (
	"context"

	"github.com/go-logr/logr"
	specsv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha4"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// V1AlphaSpec defines an interface containing callback methods for the v1alpha4 API
// traffic access objects
type v1AlphaSpecs interface {
	UpsertHTTPRouteGroup
	DeleteHTTPRouteGroup
	UpsertTCPRoute
	DeleteTCPRoute
	UpsertUDPRoute
	DeleteUDPRoute
}

// UpsertHTTPRouteGroup defines a callback function for updating or
// inserting a new HTTPRouteGroup
type UpsertHTTPRouteGroup interface {
	UpsertHTTPRouteGroup(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		rg *specsv1alpha4.HTTPRouteGroup) (ctrl.Result, error)
}

// DeleteHTTPRouteGroup defines a callback function for deleting
// a new HTTPRouteGroup
type DeleteHTTPRouteGroup interface {
	DeleteHTTPRouteGroup(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tr *specsv1alpha4.HTTPRouteGroup) (ctrl.Result, error)
}

// UpsertTCPRoute defines a callback function for updating or
// inserting a new TCPRoute
type UpsertTCPRoute interface {
	UpsertTCPRoute(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tr *specsv1alpha4.TCPRoute) (ctrl.Result, error)
}

// DeleteTCPRoute defines a callback function for deleting
// a new TCPRoute
type DeleteTCPRoute interface {
	DeleteTCPRoute(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		rg *specsv1alpha4.TCPRoute) (ctrl.Result, error)
}

// UpsertUDPRoute defines a callback function for updating or
// inserting a new UDPRoute
type UpsertUDPRoute interface {
	UpsertUDPRoute(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tr *specsv1alpha4.UDPRoute) (ctrl.Result, error)
}

// DeleteTCPRoute defines a callback function for deleting
// a new TCPRoute
type DeleteUDPRoute interface {
	DeleteUDPRoute(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		rg *specsv1alpha4.UDPRoute) (ctrl.Result, error)
}

// UpsertHTTPRouteGroup will call the user defined UpsertHTTPRouteGroup callback
// when defined
func (a *v1AlphaImpl) UpsertHTTPRouteGroup(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	rg *specsv1alpha4.HTTPRouteGroup,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(UpsertHTTPRouteGroup)

	if !ok {
		l.Info("Client code does not implement UpsertHTTPRouteGroup")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.UpsertHTTPRouteGroup(ctx, r, l, rg)
}

// DeleteHTTPRouteGroup will call the user defined DeleteHTTPRouteGroup callback
// when defined
func (a *v1AlphaImpl) DeleteHTTPRouteGroup(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	rg *specsv1alpha4.HTTPRouteGroup,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(DeleteHTTPRouteGroup)

	if !ok {
		l.Info("Client code does not implement DeleteTrafficTarget")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.DeleteHTTPRouteGroup(ctx, r, l, rg)
}

// UpsertTCPRoute will call the user defined UpsertTCPRoute callback
// when defined
func (a *v1AlphaImpl) UpsertTCPRoute(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tr *specsv1alpha4.TCPRoute,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(UpsertTCPRoute)

	if !ok {
		l.Info("Client code does not implement UpsertTCPRoute")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.UpsertTCPRoute(ctx, r, l, tr)
}

// DeleteTCPRoute will call the user defined DeleteTCPRoute callback
// when defined
func (a *v1AlphaImpl) DeleteTCPRoute(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tr *specsv1alpha4.TCPRoute,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(DeleteTCPRoute)

	if !ok {
		l.Info("Client code does not implement DeleteTCPRoute")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.DeleteTCPRoute(ctx, r, l, tr)
}

// UpsertTCPRoute will call the user defined UpsertTCPRoute callback
// when defined
func (a *v1AlphaImpl) UpsertUDPRoute(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	ur *specsv1alpha4.UDPRoute,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(UpsertUDPRoute)

	if !ok {
		l.Info("Client code does not implement UpsertUDPRoute")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.UpsertUDPRoute(ctx, r, l, ur)
}

// DeleteUDPRoute will call the user defined DeleteUDPRoute callback
// when defined
func (a *v1AlphaImpl) DeleteUDPRoute(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	ur *specsv1alpha4.UDPRoute,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(DeleteUDPRoute)

	if !ok {
		l.Info("Client code does not implement DeleteUDPRoute")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.DeleteUDPRoute(ctx, r, l, ur)
}
