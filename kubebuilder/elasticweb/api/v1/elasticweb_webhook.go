/*
Copyright 2024.

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

package v1

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var elasticweblog = logf.Log.WithName("elasticweb-resource")

// SetupWebhookWithManager will setup the manager to manage the webhooks
func (r *ElasticWeb) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// 和填写默认值有关
// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// 如果要操作资源，需要在此增加权限
// +kubebuilder:webhook:path=/mutate-elasticweb-com-bolingcavalry-v1-elasticweb,mutating=true,failurePolicy=fail,sideEffects=None,groups=elasticweb.com.bolingcavalry,resources=elasticwebs,verbs=create;update,versions=v1,name=melasticweb.kb.io,admissionReviewVersions=v1

// 有此对象才会开启填写默认值的逻辑
var _ webhook.Defaulter = &ElasticWeb{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
// 填写默认值的代码
func (r *ElasticWeb) Default() {
	elasticweblog.Info("default", "name", r.Name)
	// 如果创建的时候没有输入总QPS，就设置个默认值
	if r.Spec.TotalQPS == nil {
		r.Spec.TotalQPS = new(int32)
		*r.Spec.TotalQPS = 1300
		elasticweblog.Info("a. TotalQPS is nil, set default value now", "TotalQPS", *r.Spec.TotalQPS)
	} else {
		elasticweblog.Info("b. TotalQPS exists", "TotalQPS", *r.Spec.TotalQPS)
	}

}

// 和校验有关
// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// 校验时如果要操作其他资源，在此修改
// +kubebuilder:webhook:path=/validate-elasticweb-com-bolingcavalry-v1-elasticweb,mutating=false,failurePolicy=fail,sideEffects=None,groups=elasticweb.com.bolingcavalry,resources=elasticwebs,verbs=create;update,versions=v1,name=velasticweb.kb.io,admissionReviewVersions=v1

// 有此实例，逻辑才会生效
var _ webhook.Validator = &ElasticWeb{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
// 新增时校验逻辑
func (r *ElasticWeb) ValidateCreate() (admission.Warnings, error) {
	elasticweblog.Info("validate create", "name", r.Name)

	return r.validateElasticWeb()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ElasticWeb) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	elasticweblog.Info("validate update", "name", r.Name)

	return r.validateElasticWeb()
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ElasticWeb) ValidateDelete() (admission.Warnings, error) {
	elasticweblog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}

func (r *ElasticWeb) validateElasticWeb() (warnings admission.Warnings, err error) {
	var allErrs field.ErrorList

	if *r.Spec.SinglePodQPS > 1000 {
		elasticweblog.Info("c. Invalid SinglePodQPS")

		err := field.Invalid(field.NewPath("spec").Child("singlePodQPS"),
			*r.Spec.SinglePodQPS,
			"d. must be less than 1000")

		allErrs = append(allErrs, err)

		warnings = append(warnings, "Invalid SinglePodQPS, must be less than 1000")
		return warnings, apierrors.NewInvalid(
			schema.GroupKind{Group: "elasticweb.com.bolingcavalry", Kind: "ElasticWeb"},
			r.Name,
			allErrs)
	} else {
		elasticweblog.Info("e. SinglePodQPS is valid")
		return nil, nil
	}
}
