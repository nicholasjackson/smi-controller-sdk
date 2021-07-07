package controller

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	"github.com/go-logr/logr"
	accessv1alpha1 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha1"
	accessv1alpha2 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha2"
	accessv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/access/v1alpha3"

	splitv1alpha1 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha1"
	splitv1alpha2 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha2"
	splitv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha3"
	splitv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/split/v1alpha4"

	specsv1alpha1 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha1"
	specsv1alpha2 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha2"
	specsv1alpha3 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha3"
	specsv1alpha4 "github.com/servicemeshinterface/smi-controller-sdk/apis/specs/v1alpha4"

	accesscontrollers "github.com/servicemeshinterface/smi-controller-sdk/controllers/access"
	specscontrollers "github.com/servicemeshinterface/smi-controller-sdk/controllers/specs"
	splitcontrollers "github.com/servicemeshinterface/smi-controller-sdk/controllers/split"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
	stopChan chan (struct{})
)

// Config holds the configuration options for the controller
type Config struct {
	HealthProbeBindAddress string
	MetricsBindAddress     string
	WebhookBindAddress     string
	WebhooksEnabled        bool
	LeaderElectionID       string
	LeaderElectionEnabled  bool
	Logger                 logr.Logger
}

// DefaultConfig returns an instance of Config with the default settings
func DefaultConfig() Config {
	return Config{
		HealthProbeBindAddress: ":9444",
		WebhookBindAddress:     ":9443",
		MetricsBindAddress:     ":9102",
		LeaderElectionID:       "4ede023f.smi-spec.io",
		WebhooksEnabled:        true,
		Logger:                 zap.New(zap.UseDevMode(true)),
	}
}

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	_ = accessv1alpha1.AddToScheme(scheme)
	_ = accessv1alpha2.AddToScheme(scheme)
	_ = accessv1alpha3.AddToScheme(scheme)

	_ = splitv1alpha1.AddToScheme(scheme)
	_ = splitv1alpha2.AddToScheme(scheme)
	_ = splitv1alpha3.AddToScheme(scheme)
	_ = splitv1alpha4.AddToScheme(scheme)

	_ = specsv1alpha1.AddToScheme(scheme)
	_ = specsv1alpha2.AddToScheme(scheme)
	_ = specsv1alpha3.AddToScheme(scheme)
	_ = specsv1alpha4.AddToScheme(scheme)
	// +kubebuilder:scaffold:scheme
}

// Start the controller
func Start(config Config) {
	webhookAddress := strings.Split(config.WebhookBindAddress, ":")[0]
	webhookPortString := strings.Split(config.WebhookBindAddress, ":")[1]
	webhookPort, _ := strconv.Atoi(webhookPortString)

	if webhookAddress == "" {
		webhookAddress = "0.0.0.0"
	}

	ctrl.SetLogger(config.Logger)

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:                 scheme,
		MetricsBindAddress:     config.MetricsBindAddress,
		HealthProbeBindAddress: config.HealthProbeBindAddress,
		Port:                   webhookPort,
		Host:                   webhookAddress,
		LeaderElection:         config.LeaderElectionEnabled,
		LeaderElectionID:       config.LeaderElectionID,
	})

	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	// add the health checks
	mgr.AddHealthzCheck("healthz", func(r *http.Request) error {
		setupLog.Info("Health check called")
		//	TODO expose this functionality to the consumer of the SDK
		return nil
	})

	// add the readyness checks
	mgr.AddReadyzCheck("readyz", func(r *http.Request) error {
		setupLog.Info("Ready check called")
		//	TODO expose this functionality to the consumer of the SDK
		return nil
	})

	// Traffic Access Controllers
	if err = (&accesscontrollers.TrafficTargetReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "TrafficTarget")
		os.Exit(1)
	}

	// Traffic Split Controllers
	if err = (&splitcontrollers.TrafficSplitReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "TrafficSplit")
		os.Exit(1)
	}

	// Traffic Spec Controllers
	if err = (&specscontrollers.HTTPRouteGroupReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "HTTPRouteGroup")
		os.Exit(1)
	}

	if err = (&specscontrollers.TCPRouteReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "TCPRoute")
		os.Exit(1)
	}

	if err = (&specscontrollers.UDPRouteReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "UDPRoute")
		os.Exit(1)
	}

	if config.WebhooksEnabled {
		// Access conversion webhooks
		if err = (&accessv1alpha1.TrafficTarget{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "accessv1alpha1.TrafficTarget")
			os.Exit(1)
		}
		if err = (&accessv1alpha2.TrafficTarget{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "accessv1alpha1.TrafficTarget")
			os.Exit(1)
		}

		// Splits conversion webhooks
		if err = (&splitv1alpha1.TrafficSplit{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "accessv1alpha1.TrafficSplit")
			os.Exit(1)
		}
		if err = (&splitv1alpha2.TrafficSplit{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "accessv1alpha2.TrafficSplit")
			os.Exit(1)
		}
		if err = (&splitv1alpha3.TrafficSplit{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "accessv1alpha3.TrafficSplit")
			os.Exit(1)
		}

		// Specs conversion webhooks
		if err = (&specsv1alpha1.HTTPRouteGroup{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "accessv1alpha1.HTTPRouteGroup")
			os.Exit(1)
		}
		if err = (&specsv1alpha2.HTTPRouteGroup{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "accessv1alpha2.HTTPRouteGroup")
			os.Exit(1)
		}
		if err = (&specsv1alpha3.HTTPRouteGroup{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "accessv1alpha3.HTTPRouteGroup")
			os.Exit(1)
		}
		if err = (&specsv1alpha1.TCPRoute{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "specsv1alpha1.TCPRoute")
			os.Exit(1)
		}
		if err = (&specsv1alpha2.TCPRoute{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "specsv1alpha2.TCPRoute")
			os.Exit(1)
		}
		if err = (&specsv1alpha3.TCPRoute{}).SetupWebhookWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create webhook", "webhook", "specsv1alpha3.TCPRoute")
			os.Exit(1)
		}
	}
	// +kubebuilder:scaffold:builder

	stopChan := ctrl.SetupSignalHandler()

	setupLog.Info("starting manager")
	if err := mgr.Start(stopChan); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

// Stop the controller and shutdown gracefully
// TODO this seems to block on the channe, investigate
func Stop() {
	fmt.Println("stopping")
	stopChan <- struct{}{}
	fmt.Println("stopped")
}
