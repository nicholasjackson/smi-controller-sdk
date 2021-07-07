/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package access

import (
	"context"

	"github.com/servicemeshinterface/smi-controller-sdk/controllers/helpers"
	"github.com/servicemeshinterface/smi-controller-sdk/sdk"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	accessv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha3"
)

// TrafficTargetReconciler reconciles a TrafficTarget object
type TrafficTargetReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=access.smi-spec.io,resources=traffictargets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=access.smi-spec.io,resources=traffictargets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=access.smi-spec.io,resources=traffictargets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TrafficTarget object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *TrafficTargetReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	tt := &accessv1alpha3.TrafficTarget{}
	if err := r.Get(ctx, req.NamespacedName, tt); err != nil {
		logger.Info("unable to fetch TrafficTarget, most likely deleted")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	ttFinalizerName := "traffictarget.finalizers.smi-controller"

	// examine DeletionTimestamp to determine if object is under deletion
	if tt.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object. This is equivalent
		// registering our finalizer.
		if !helpers.ContainsString(tt.ObjectMeta.Finalizers, ttFinalizerName) {
			tt.ObjectMeta.Finalizers = append(tt.ObjectMeta.Finalizers, ttFinalizerName)
			if err := r.Update(context.Background(), tt); err != nil {
				return ctrl.Result{}, err
			}

			return ctrl.Result{}, nil
		}
	} else {
		// The object is being deleted
		if helpers.ContainsString(tt.ObjectMeta.Finalizers, ttFinalizerName) {
			// our finalizer is present, so lets handle any external dependency
			sdk.API().V1Alpha().DeleteTrafficTarget(ctx, r.Client, logger, tt)

			// remove our finalizer from the list and update it.
			tt.ObjectMeta.Finalizers = helpers.RemoveString(tt.ObjectMeta.Finalizers, ttFinalizerName)
			if err := r.Update(context.Background(), tt); err != nil {
				return ctrl.Result{}, err
			}
		}

		// Stop reconciliation as the item is being deleted
		return ctrl.Result{}, nil
	}

	return sdk.API().V1Alpha().UpsertTrafficTarget(ctx, r.Client, logger, tt)
}

func (r *TrafficTargetReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&accessv1alpha3.TrafficTarget{}).
		Complete(r)
}
