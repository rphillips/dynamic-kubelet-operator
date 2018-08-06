package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DynamicKubeletOperatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []DynamicKubeletOperator `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DynamicKubeletOperator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              DynamicKubeletOperatorSpec   `json:"spec"`
	Status            DynamicKubeletOperatorStatus `json:"status,omitempty"`
}

type DynamicKubeletOperatorSpec struct {
	// Fill me
}
type DynamicKubeletOperatorStatus struct {
	// Fill me
}
