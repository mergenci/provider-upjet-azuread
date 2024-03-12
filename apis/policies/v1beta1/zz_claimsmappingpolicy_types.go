// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ClaimsMappingPolicyInitParameters struct {

	// The claims mapping policy. This is a JSON formatted string, for which the jsonencode() function can be used.
	// A string collection containing a JSON string that defines the rules and settings for this policy
	Definition []*string `json:"definition,omitempty" tf:"definition,omitempty"`

	// The display name for this Claims Mapping Policy.
	// Display name for this policy
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`
}

type ClaimsMappingPolicyObservation struct {

	// The claims mapping policy. This is a JSON formatted string, for which the jsonencode() function can be used.
	// A string collection containing a JSON string that defines the rules and settings for this policy
	Definition []*string `json:"definition,omitempty" tf:"definition,omitempty"`

	// The display name for this Claims Mapping Policy.
	// Display name for this policy
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The ID of the Claims Mapping Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ClaimsMappingPolicyParameters struct {

	// The claims mapping policy. This is a JSON formatted string, for which the jsonencode() function can be used.
	// A string collection containing a JSON string that defines the rules and settings for this policy
	// +kubebuilder:validation:Optional
	Definition []*string `json:"definition,omitempty" tf:"definition,omitempty"`

	// The display name for this Claims Mapping Policy.
	// Display name for this policy
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`
}

// ClaimsMappingPolicySpec defines the desired state of ClaimsMappingPolicy
type ClaimsMappingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClaimsMappingPolicyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ClaimsMappingPolicyInitParameters `json:"initProvider,omitempty"`
}

// ClaimsMappingPolicyStatus defines the observed state of ClaimsMappingPolicy.
type ClaimsMappingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClaimsMappingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ClaimsMappingPolicy is the Schema for the ClaimsMappingPolicys API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type ClaimsMappingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.definition) || (has(self.initProvider) && has(self.initProvider.definition))",message="spec.forProvider.definition is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || (has(self.initProvider) && has(self.initProvider.displayName))",message="spec.forProvider.displayName is a required parameter"
	Spec   ClaimsMappingPolicySpec   `json:"spec"`
	Status ClaimsMappingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClaimsMappingPolicyList contains a list of ClaimsMappingPolicys
type ClaimsMappingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClaimsMappingPolicy `json:"items"`
}

// Repository type metadata.
var (
	ClaimsMappingPolicy_Kind             = "ClaimsMappingPolicy"
	ClaimsMappingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClaimsMappingPolicy_Kind}.String()
	ClaimsMappingPolicy_KindAPIVersion   = ClaimsMappingPolicy_Kind + "." + CRDGroupVersion.String()
	ClaimsMappingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(ClaimsMappingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&ClaimsMappingPolicy{}, &ClaimsMappingPolicyList{})
}
