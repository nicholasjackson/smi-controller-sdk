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

const (
	// HTTPRouteMethodAll is a wildcard for all HTTP methods
	HTTPRouteMethodAll HTTPRouteMethod = "*"
	// HTTPRouteMethodGet HTTP GET method
	HTTPRouteMethodGet HTTPRouteMethod = "GET"
	// HTTPRouteMethodHead HTTP HEAD method
	HTTPRouteMethodHead HTTPRouteMethod = "HEAD"
	// HTTPRouteMethodPut HTTP PUT method
	HTTPRouteMethodPut HTTPRouteMethod = "PUT"
	// HTTPRouteMethodPost HTTP POST method
	HTTPRouteMethodPost HTTPRouteMethod = "POST"
	// HTTPRouteMethodDelete HTTP DELETE method
	HTTPRouteMethodDelete HTTPRouteMethod = "DELETE"
	// HTTPRouteMethodConnect HTTP CONNECT method
	HTTPRouteMethodConnect HTTPRouteMethod = "CONNECT"
	// HTTPRouteMethodOptions HTTP OPTIONS method
	HTTPRouteMethodOptions HTTPRouteMethod = "OPTIONS"
	// HTTPRouteMethodTrace HTTP TRACE method
	HTTPRouteMethodTrace HTTPRouteMethod = "TRACE"
	// HTTPRouteMethodPatch HTTP PATCH method
	HTTPRouteMethodPatch HTTPRouteMethod = "PATCH"
)

// HTTPHeaders is a map of key/value pairs which match HTTP header name and value
type HTTPHeaders map[string]string

// HTTPRouteMethod are methods allowed by the route
type HTTPRouteMethod string

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// HTTPRouteGroupSpec defines the desired state of HTTPRouteGroup
type HTTPRouteGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Routes for inbound traffic
	Matches []HTTPMatch `json:"matches,omitempty"`
}

// HTTPMatch defines an individual route for HTTP traffic
type HTTPMatch struct {
	// Name is the name of the match for referencing in a TrafficTarget
	Name string `json:"name,omitempty"`

	// Methods for inbound traffic as defined in RFC 7231
	// https://tools.ietf.org/html/rfc7231#section-4
	Methods []string `json:"methods,omitempty"`

	// PathRegex is a regular expression defining the route
	PathRegex string `json:"pathRegex,omitempty"`

	// Headers is a list of headers used to match HTTP traffic
	Headers HTTPHeaders `json:"headers,omitempty"`
}

// HTTPRouteGroupStatus defines the observed state of HTTPRouteGroup
type HTTPRouteGroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:storageversion

// HTTPRouteGroup is the Schema for the httproutegroups API
type HTTPRouteGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HTTPRouteGroupSpec   `json:"spec,omitempty"`
	Status HTTPRouteGroupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// HTTPRouteGroupList contains a list of HTTPRouteGroup
type HTTPRouteGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HTTPRouteGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HTTPRouteGroup{}, &HTTPRouteGroupList{})
}
