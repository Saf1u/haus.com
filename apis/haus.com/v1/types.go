// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)



// HausSpec defines the desired state of Haus
type HausSpec struct {
	Count    int    `json:"count"`
	Location string `json:"location"`
	Replicas int `json:"replicas"`
}

// HausStatus defines the observed state of Haus
type HausStatus struct {
	State int `json:"state"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Haus is the Schema for the Haus custom resource
// +genclient
type Haus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HausSpec   `json:"spec,omitempty"`
	Status HausStatus `json:"status,omitempty"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HausList contains a list of Haus
type HausList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Haus `json:"items"`
}
