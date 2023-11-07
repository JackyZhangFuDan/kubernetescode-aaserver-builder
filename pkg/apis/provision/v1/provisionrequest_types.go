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

package v1

import (
	"context"

	provisionv1alpha1 "github.com/kubernetescode-aaserver-builder/pkg/apis/provision/v1alpha1"
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
	IngressEntrance    string   `json:"ingressEntrance" `
	BusinessDbVolume   DbVolume `json:"businessDbVolume"`
	BusinessDbCpuLimit string   `json:"businessDbCpuLimit"`
	NamespaceName      string   `json:"namespaceName"`
}

var _ resource.Object = &ProvisionRequest{}
var _ resourcestrategy.Validater = &ProvisionRequest{}

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
		Version:  "v1",
		Resource: "provisionrequests",
	}
}

func (in *ProvisionRequest) IsStorageVersion() bool {
	return false
}

func (in *ProvisionRequest) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

var _ resource.ObjectList = &ProvisionRequestList{}

func (in *ProvisionRequestList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}

// ProvisionRequestStatus defines the observed state of ProvisionRequest
type ProvisionRequestStatus struct {
	metav1.TypeMeta `json:",inline" `
	IngressReady    bool `json:"ingressReady" `
	DbReady         bool `json:"dbReady"`
}

func (in ProvisionRequestStatus) SubResourceName() string {
	return "status"
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

// NewStorageVersionObject returns a new empty instance of storage version.
func (in *ProvisionRequest) NewStorageVersionObject() runtime.Object {
	return &provisionv1alpha1.ProvisionRequest{}
}

// ConvertToStorageVersion receives an new instance of storage version object as the conversion target
// and overwrites it to the equal form of the current resource version.
func (in *ProvisionRequest) ConvertToStorageVersion(storageObj runtime.Object) error {
	storageObj.(*provisionv1alpha1.ProvisionRequest).ObjectMeta = in.ObjectMeta
	//storageObj.(*provisionv1alpha1.ProvisionRequest).TypeMeta = in.TypeMeta
	storageObj.(*provisionv1alpha1.ProvisionRequest).Spec.BusinessDbCpuLimit = in.Spec.BusinessDbCpuLimit
	storageObj.(*provisionv1alpha1.ProvisionRequest).Spec.BusinessDbVolume = provisionv1alpha1.DbVolume(in.Spec.BusinessDbVolume)
	storageObj.(*provisionv1alpha1.ProvisionRequest).Spec.IngressEntrance = in.Spec.IngressEntrance
	storageObj.(*provisionv1alpha1.ProvisionRequest).Spec.NamespaceName = in.Spec.NamespaceName
	storageObj.(*provisionv1alpha1.ProvisionRequest).Status.DbReady = in.Status.DbReady
	storageObj.(*provisionv1alpha1.ProvisionRequest).Status.IngressReady = in.Status.IngressReady
	return nil
}

// ConvertFromStorageVersion receives an instance of storage version as the conversion source and
// in-place mutates the current object to the equal form of the storage version object.
func (in *ProvisionRequest) ConvertFromStorageVersion(storageObj runtime.Object) error {
	//in.TypeMeta = storageObj.(*provisionv1alpha1.ProvisionRequest).TypeMeta
	in.ObjectMeta = storageObj.(*provisionv1alpha1.ProvisionRequest).ObjectMeta
	in.Spec.BusinessDbCpuLimit = storageObj.(*provisionv1alpha1.ProvisionRequest).Spec.BusinessDbCpuLimit
	in.Spec.BusinessDbVolume = DbVolume(storageObj.(*provisionv1alpha1.ProvisionRequest).Spec.BusinessDbVolume)
	in.Spec.IngressEntrance = storageObj.(*provisionv1alpha1.ProvisionRequest).Spec.IngressEntrance
	in.Spec.NamespaceName = storageObj.(*provisionv1alpha1.ProvisionRequest).Spec.NamespaceName
	in.Status.DbReady = storageObj.(*provisionv1alpha1.ProvisionRequest).Status.DbReady
	in.Status.IngressReady = storageObj.(*provisionv1alpha1.ProvisionRequest).Status.IngressReady
	return nil
}

var _ resource.MultiVersionObject = &ProvisionRequest{}
