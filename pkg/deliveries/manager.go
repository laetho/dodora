package deliveries

import (
	"fmt"
	"os"

	dtv1 "github.com/laetho/deliverytracker/apis/deliverytracker/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(dtv1.AddToScheme(scheme))
}

// Setup and run the controller for keeping the in-memory Deliveries{} in
// sync with what is in kubernets.
func Run() {

	ctrl.SetLogger(zap.New())
	ctrl.Log.WithName("controller")

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		LeaderElection:         false,
		Scheme:                 scheme,
		Port:                   9443,
		MetricsBindAddress:     ":8081",
		HealthProbeBindAddress: ":8082",
	})
	if err != nil {
		fmt.Println(err)
		panic("unable to start state reconcile controller")
	}

	if err = (&DeliveriesReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {

		os.Exit(1)
	}

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		os.Exit(1)
	}
}
