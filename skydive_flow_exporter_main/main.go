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

package main

import (
	"flag"
	"os"

	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	skydivev1 "skydive/api/v1"
	"skydive/controllers"
	// +kubebuilder:scaffold:imports
)

var (
	flowExporterScheme   = runtime.NewScheme()
	flowExporterSetuplog = ctrl.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(flowExporterScheme)

	_ = skydivev1.AddToScheme(flowExporterScheme)
	// +kubebuilder:scaffold:flow_exporter_scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8081", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             flowExporterScheme,
		MetricsBindAddress: metricsAddr,
		Port:               9443,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "c6f81ab3.example.com",
	})
	if err != nil {
		flowExporterSetuplog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if err = (&controllers.SkydiveFlowExporterReconciler{
		Client: mgr.GetClient(),
		Log:    ctrl.Log.WithName("controllers").WithName("SkydiveFlowExporter"),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		flowExporterSetuplog.Error(err, "unable to create controller", "controller", "SkydiveFlowExporter")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	flowExporterSetuplog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		flowExporterSetuplog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
