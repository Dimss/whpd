/*
Copyright 2018 The Kubernetes Authors.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


type WhPlatformSpec struct {
	Foo string `json:"foo,omitempty"`
}


type WhPlatformStatus struct {
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Foo",type=string,JSONPath=`.spec.foo`
type WhPlatform struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Spec   WhPlatformSpec   `json:"spec,omitempty"`
	Status WhPlatformStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type WhPlatformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WhPlatform `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WhPlatform{}, &WhPlatformList{})
}