/*
Copyright 2021 ysicing <i@ysicing.me>.

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

package devops

import (
	"context"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"time"

	devopsv1beta1 "github.com/ysicing/k3s-autoscaler/apis/devops/v1beta1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AutoScalerReconciler reconciles a AutoScaler object
type AutoScalerReconciler struct {
	client.Client
	Scheme     *runtime.Scheme
	Recorder   record.EventRecorder
	PodLister  corelisters.PodLister
	NodeLister corelisters.NodeLister
}

//+kubebuilder:rbac:groups=core,resources=nodes,verbs=get;list;watch;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups=core,resources=events,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=devops.ysicing.me,resources=autoscalers,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=devops.ysicing.me,resources=autoscalers/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=devops.ysicing.me,resources=autoscalers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the AutoScaler object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *AutoScalerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (res ctrl.Result, err error) {
	start := time.Now()
	klog.V(5).Infof("Starting to process K3sAutoScaler %v", req.Name)
	defer func() {
		if err != nil {
			if err != nil {
				klog.Warningf("Failed to process K3sAutoScaler %v err %v, elapsedTime %v", req.Name, time.Since(start), err)
			} else if res.RequeueAfter > 0 {
				klog.Infof("Finish to process K3sAutoScaler %v, elapsedTime %v, RetryAfter %v", req.Name, time.Since(start), res.RequeueAfter)
			} else {
				klog.Infof("Finish to process K3sAutoScaler %v, elapsedTime %v", req.Name, time.Since(start))
			}
		}
	}()
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AutoScalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	c, err := controller.New("k3s-autoscaler", mgr, controller.Options{
		MaxConcurrentReconciles: 3,
		Reconciler:              r,
	})
	if err != nil {
		return err
	}
	// watch pods
	err = c.Watch(&source.Kind{Type: &v1.Pod{}}, &eventHandler{reader: mgr.GetCache()})
	if err != nil {
		return err
	}
	// watch nodes
	err = c.Watch(&source.Kind{Type: &v1.Node{}}, &eventHandler{reader: mgr.GetCache()})
	if err != nil {
		return err
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&devopsv1beta1.AutoScaler{}).
		Complete(r)
}
