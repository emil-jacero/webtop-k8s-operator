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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	// EnvoyConfig includes configuration details for Envoy Proxy
	EnvoyConfig EnvoyConfig `json:"envoyConfig"`

	// OPAConfig includes configuration details for Open Policy Agent
	OPAConfig OPAConfig `json:"opaConfig"`
}

// EnvoyConfig stores configuration for Envoy Proxy
type EnvoyConfig struct {
	// Image is the Docker image for Envoy Proxy
	Image string `json:"image"`

	// ConfigMapName is the name of the ConfigMap containing the Envoy configuration
	ConfigMapName string `json:"configMapName"`
}

// OPAConfig stores configuration for Open Policy Agent
type OPAConfig struct {
	// Image is the Docker image for OPA
	Image string `json:"image"`

	// PolicyConfigMap is the name of the ConfigMap containing OPA policies
	PolicyConfigMap string `json:"policyConfigMap"`
}

// GatewayStatus defines the observed state of Gateway
type GatewayStatus struct {
	// Conditions represent the latest available observations of an object's state
	Conditions []metav1.Condition `json:"conditions"`

	// Active indicates whether the gateway is active and all components are running
	Active bool `json:"active"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Gateway is the Schema for the gateways API
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GatewaySpec   `json:"spec,omitempty"`
	Status GatewayStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// GatewayList contains a list of Gateway
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}
