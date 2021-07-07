package sdk

import (
	"context"

	"github.com/go-logr/logr"
	accessv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha3"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// V1AlphaAccess defines an interface containing callback methods for the v1alpha4 API
// traffic access objects
type v1AlphaAccess interface {
	UpsertTrafficTarget
	DeleteTrafficTarget
}

// UpsertTrafficTarget defines a callback function for updating or
// inserting a new TrafficTarget
type UpsertTrafficTarget interface {
	UpsertTrafficTarget(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tt *accessv1alpha3.TrafficTarget) (ctrl.Result, error)
}

// DeleteTrafficTarget defines a callback function for deleting
// a new TrafficTarget
type DeleteTrafficTarget interface {
	DeleteTrafficTarget(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tt *accessv1alpha3.TrafficTarget) (ctrl.Result, error)
}

// UpsertTrafficTarget will call the user defined UpsertTrafficTarget callback
// when defined
func (a *v1AlphaImpl) UpsertTrafficTarget(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha3.TrafficTarget,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(UpsertTrafficTarget)

	if !ok {
		l.Info("Client code does not implement UpsertTrafficTarget")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.UpsertTrafficTarget(ctx, r, l, tt)
}

func (a *v1AlphaImpl) DeleteTrafficTarget(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *accessv1alpha3.TrafficTarget,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(DeleteTrafficTarget)

	if !ok {
		l.Info("Client code does not implement DeleteTrafficTarget")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.DeleteTrafficTarget(ctx, r, l, tt)
}
