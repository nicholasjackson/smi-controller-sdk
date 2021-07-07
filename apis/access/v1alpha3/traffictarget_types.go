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

package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TrafficTargetSpec defines the desired state of TrafficTarget
// It is the specification of a TrafficTarget
type TrafficTargetSpec struct {
	// Selector is the pod or group of pods to allow ingress traffic
	Destination IdentityBindingSubject `json:"destination"`

	// Sources are the pod or group of pods to allow ingress traffic
	Sources []IdentityBindingSubject `json:"sources,omitempty"`

	// Rules are the traffic rules to allow (HTTPRoutes | TCPRoute)
	Rules []TrafficTargetRule `json:"rules,omitempty"`
}

// TrafficTargetRule is the TrafficSpec to allow for a TrafficTarget
type TrafficTargetRule struct {
	// Kind is the kind of TrafficSpec to allow
	Kind string `json:"kind"`

	// Name of the TrafficSpec to use
	Name string `json:"name"`

	// Matches is a list of TrafficSpec routes to allow traffic for
	// +optional
	Matches []string `json:"matches,omitempty"`
}

// IdentityBindingSubject is a Kubernetes objects which should be allowed access to the TrafficTarget
type IdentityBindingSubject struct {
	// Kind is the type of Subject to allow ingress (ServiceAccount | Group)
	Kind string `json:"kind"`

	// Name of the Subject, i.e. ServiceAccountName
	Name string `json:"name"`

	// Namespace where the Subject is deployed
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// TrafficTargetStatus defines the observed state of UDPRoute
type TrafficTargetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// TrafficTarget associates a set of traffic definitions (rules) with a service identity which is allocated to a group of pods.
// Access is controlled via referenced TrafficSpecs and by a list of source service identities.
// * If a pod which holds the referenced service identity makes a call to the destination on one of the defined routes then access
//   will be allowed
// * Any pod which attempts to connect and is not in the defined list of sources will be denied
// * Any pod which is in the defined list, but attempts to connect on a route which is not in the list of the
//   TrafficSpecs will be denied
type TrafficTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TrafficTargetSpec `json:"spec"`
}

//+kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficTargetList contains a list of TrafficTarget
type TrafficTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrafficTarget     `json:"items"`
	Status          TrafficTargetStatus `json:"status,omitempty"`
}

func init() {
	SchemeBuilder.Register(&TrafficTarget{}, &TrafficTargetList{})
}
