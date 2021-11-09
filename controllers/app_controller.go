/*
Copyright 2021.

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

package controllers

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appv1 "github.com/ktlcove/app-operator/api/v1"
	// "github.com/fatih/structs"
)

// AppReconciler reconciles a App object
type AppReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=app.ktlcove.io,resources=apps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=app.ktlcove.io,resources=apps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=app.ktlcove.io,resources=apps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the App object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.9.2/pkg/reconcile

func (r *AppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// your logic here
	log.Log.Info("app reconcile")

	app := &appv1.App{}
	err := r.Get(ctx, req.NamespacedName, app)

	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	markDeleted := app.GetDeletionTimestamp() != nil

	if markDeleted && controllerutil.ContainsFinalizer(app, appv1.APP_FINALIZER) {
		err := r.Finialize(app)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	// log.Log.Info("got app", structs.Map(app))
	// app_string, _ := json.Marshal(app)
	// log.Log.Info(string(app_string))
	// ctx_string, _ := json.Marshal(ctx)
	// log.Log.Info(string(ctx_string))
	// req_string, _ := json.Marshal(req)
	// log.Log.Info(string(req_string))
	return ctrl.Result{}, nil
}

func (r *AppReconciler) Finialize(app *appv1.App) error {
	log.Log.Info("app controller finialize app", app.ObjectMeta.Namespace, app.ObjectMeta.Name)
	return nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *AppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appv1.App{}).
		Complete(r)
}
