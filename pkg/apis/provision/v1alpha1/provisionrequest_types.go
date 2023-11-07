/*
Copyright 2023.

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

package v1alpha1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

type DbVolume string

const DbVolumeBig DbVolume = "BIG"
const DbVolumeSmall DbVolume = "SMALL"
const DbVolumeMedium DbVolume = "MEDIUM"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ProvisionRequest
// +k8s:openapi-gen=true
type ProvisionRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProvisionRequestSpec   `json:"spec,omitempty"`
	Status ProvisionRequestStatus `json:"status,omitempty"`
}

// ProvisionRequestList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ProvisionRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ProvisionRequest `json:"items"`
}

// ProvisionRequestSpec defines the desired state of ProvisionRequest
type ProvisionRequestSpec struct {
	IngressEntrance  string   `json:"ingressEntrance" `
	BusinessDbVolume DbVolume `json:"businessDbVolume"`
	NamespaceName    string   `json:"namespaceName"`

	BusinessDbCpuLimit string `json:"businessDbCpuLimit"`
}

var _ resource.Object = &ProvisionRequest{}

func (in *ProvisionRequest) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *ProvisionRequest) NamespaceScoped() bool {
	return true
}

func (in *ProvisionRequest) New() runtime.Object {
	return &ProvisionRequest{}
}

func (in *ProvisionRequest) NewList() runtime.Object {
	return &ProvisionRequestList{}
}

func (in *ProvisionRequest) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "provision.mydomain.com",
		Version:  "v1alpha1",
		Resource: "provisionrequests",
	}
}

func (in *ProvisionRequest) IsStorageVersion() bool {
	return true
}

var _ resourcestrategy.Validater = &ProvisionRequest{}

func (in *ProvisionRequest) Validate(ctx context.Context) field.ErrorList {
	errs := field.ErrorList{}

	if len(in.Spec.NamespaceName) == 0 {
		errs = append(errs,
			field.Required(field.NewPath("spec").Key("namespaceName"),
				"namespace name is required"))
	}
	if len(errs) > 0 {
		return errs
	} else {
		return nil
	}
}

var _ resourcestrategy.Defaulter = &ProvisionRequest{}

func (in *ProvisionRequest) Default() {
	if in.Spec.BusinessDbVolume == "" {
		in.Spec.BusinessDbVolume = DbVolumeMedium
	}
}

var _ resource.ObjectList = &ProvisionRequestList{}

func (in *ProvisionRequestList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

// ProvisionRequest implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &ProvisionRequest{}

func (in *ProvisionRequest) GetStatus() resource.StatusSubResource {
	return in.Status
}

// ProvisionRequestStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &ProvisionRequestStatus{}

func (in ProvisionRequestStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*ProvisionRequest).Status = in
}

var _ resource.ObjectWithArbitrarySubResource = &ProvisionRequest{}

func (in *ProvisionRequest) GetArbitrarySubResources() []resource.ArbitrarySubResource {
	return []resource.ArbitrarySubResource{
		// +kubebuilder:scaffold:subresource
		&ProvisionRequestStatus{},
	}
}
