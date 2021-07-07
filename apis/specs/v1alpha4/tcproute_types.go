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

package v1alpha4

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TCPRouteSpec defines the desired state of TCPRoute
type TCPRouteSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of TCPRoute. Edit tcproute_types.go to remove/update
	Matches TCPMatch `json:"matches,omitempty"`
}

// TCPMatch defines an individual route for TCP traffic
type TCPMatch struct {
	// Name is the name of the match for referencing in a TrafficTarget
	Name string `json:"name,omitempty"`

	// Ports to allow inbound traffic on
	Ports []int `json:"ports,omitempty"`
}

// TCPRouteStatus defines the observed state of TCPRoute
type TCPRouteStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// TCPRoute is the Schema for the tcproutes API
type TCPRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TCPRouteSpec   `json:"spec,omitempty"`
	Status TCPRouteStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// TCPRouteList contains a list of TCPRoute
type TCPRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TCPRoute `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TCPRoute{}, &TCPRouteList{})
}
