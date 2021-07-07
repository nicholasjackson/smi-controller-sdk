package sdk

import (
	"context"

	"github.com/go-logr/logr"
	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// V1AlphaAccess defines an interface containing callback methods for the v1alpha4 API
// traffic access objects
type v1AlphaSplit interface {
	UpsertTrafficSplit
	DeleteTrafficSplit
}

// UpsertTrafficSplit defines a callback function for updating or
// inserting a new TrafficSplit
type UpsertTrafficSplit interface {
	UpsertTrafficSplit(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tt *splitv1alpha4.TrafficSplit) (ctrl.Result, error)
}

// DeleteTrafficSplit defines a callback function for deleting
// a new TrafficSplit
type DeleteTrafficSplit interface {
	DeleteTrafficSplit(
		ctx context.Context,
		r client.Client,
		l logr.Logger,
		tt *splitv1alpha4.TrafficSplit) (ctrl.Result, error)
}

// UpsertTrafficSplit will call the user defined UpsertTrafficSplit callback
// when defined
func (a *v1AlphaImpl) UpsertTrafficSplit(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *splitv1alpha4.TrafficSplit,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(UpsertTrafficSplit)

	if !ok {
		l.Info("Client code does not implement UpsertTrafficSplit")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.UpsertTrafficSplit(ctx, r, l, tt)
}

func (a *v1AlphaImpl) DeleteTrafficSplit(
	ctx context.Context,
	r client.Client,
	l logr.Logger,
	tt *splitv1alpha4.TrafficSplit,
) (ctrl.Result, error) {

	// does the user api have this callback?
	v, ok := a.userV1alpha.(DeleteTrafficSplit)

	if !ok {
		l.Info("Client code does not implement DeleteTrafficSplit")
		return ctrl.Result{}, nil
	}

	// call the interface method
	return v.DeleteTrafficSplit(ctx, r, l, tt)
}
