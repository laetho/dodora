package deliveries

import (
	"context"
	"fmt"
	dtv1 "github.com/laetho/deliverytracker/apis/deliverytracker/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DeliveriesReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

func (r *DeliveriesReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	var delivery dtv1.Delivery

	if err := r.Get(ctx, req.NamespacedName, &delivery); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	fmt.Println(delivery)

	return ctrl.Result{}, nil
}

func (r *DeliveriesReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&dtv1.Delivery{}).
		Complete(r)
}