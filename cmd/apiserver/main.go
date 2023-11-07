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

package main

import (
	"k8s.io/klog"
	"sigs.k8s.io/apiserver-runtime/pkg/builder"

	// +kubebuilder:scaffold:resource-imports
	provisionv1 "github.com/kubernetescode-aaserver-builder/pkg/apis/provision/v1"

	provisionv1alpha1 "github.com/kubernetescode-aaserver-builder/pkg/apis/provision/v1alpha1"
)

func main() {
	err := builder.APIServer.
		// +kubebuilder:scaffold:resource-register
		WithResource(&provisionv1alpha1.ProvisionRequest{}).
		WithResource(&provisionv1.ProvisionRequest{}).
		WithLocalDebugExtension().
		DisableAuthorization().
		WithOptionsFns(func(options *builder.ServerOptions) *builder.ServerOptions {
			options.RecommendedOptions.CoreAPI = nil
			options.RecommendedOptions.Admission = nil
			return options
		}).
		Execute()
	if err != nil {
		klog.Fatal(err)
	}
}
