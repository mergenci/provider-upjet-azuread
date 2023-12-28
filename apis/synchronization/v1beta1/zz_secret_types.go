// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CredentialInitParameters struct {

	// The key of the secret.
	// Name for this key-value pair.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type CredentialObservation struct {

	// The key of the secret.
	// Name for this key-value pair.
	Key *string `json:"key,omitempty" tf:"key,omitempty"`
}

type CredentialParameters struct {

	// The key of the secret.
	// Name for this key-value pair.
	// +kubebuilder:validation:Optional
	Key *string `json:"key" tf:"key,omitempty"`

	// The value of the secret.
	// Value for this key-value pair.
	// +kubebuilder:validation:Required
	ValueSecretRef v1.SecretKeySelector `json:"valueSecretRef" tf:"-"`
}

type SecretInitParameters struct {

	// One or more credential blocks as documented below.
	Credential []CredentialInitParameters `json:"credential,omitempty" tf:"credential,omitempty"`

	// The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
	// The object ID of the service principal for which this synchronization secret should be created
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/serviceprincipals/v1beta1.Principal
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// Reference to a Principal in serviceprincipals to populate servicePrincipalId.
	// +kubebuilder:validation:Optional
	ServicePrincipalIDRef *v1.Reference `json:"servicePrincipalIdRef,omitempty" tf:"-"`

	// Selector for a Principal in serviceprincipals to populate servicePrincipalId.
	// +kubebuilder:validation:Optional
	ServicePrincipalIDSelector *v1.Selector `json:"servicePrincipalIdSelector,omitempty" tf:"-"`
}

type SecretObservation struct {

	// One or more credential blocks as documented below.
	Credential []CredentialObservation `json:"credential,omitempty" tf:"credential,omitempty"`

	// An ID used to uniquely identify this synchronization sec.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
	// The object ID of the service principal for which this synchronization secret should be created
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`
}

type SecretParameters struct {

	// One or more credential blocks as documented below.
	// +kubebuilder:validation:Optional
	Credential []CredentialParameters `json:"credential,omitempty" tf:"credential,omitempty"`

	// The object ID of the service principal for which this synchronization secrets should be stored. Changing this field forces a new resource to be created.
	// The object ID of the service principal for which this synchronization secret should be created
	// +crossplane:generate:reference:type=github.com/upbound/provider-azuread/apis/serviceprincipals/v1beta1.Principal
	// +kubebuilder:validation:Optional
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty" tf:"service_principal_id,omitempty"`

	// Reference to a Principal in serviceprincipals to populate servicePrincipalId.
	// +kubebuilder:validation:Optional
	ServicePrincipalIDRef *v1.Reference `json:"servicePrincipalIdRef,omitempty" tf:"-"`

	// Selector for a Principal in serviceprincipals to populate servicePrincipalId.
	// +kubebuilder:validation:Optional
	ServicePrincipalIDSelector *v1.Selector `json:"servicePrincipalIdSelector,omitempty" tf:"-"`
}

// SecretSpec defines the desired state of Secret
type SecretSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretParameters `json:"forProvider"`
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
	InitProvider SecretInitParameters `json:"initProvider,omitempty"`
}

// SecretStatus defines the observed state of Secret.
type SecretStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Secret is the Schema for the Secrets API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type Secret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecretSpec   `json:"spec"`
	Status            SecretStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretList contains a list of Secrets
type SecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Secret `json:"items"`
}

// Repository type metadata.
var (
	Secret_Kind             = "Secret"
	Secret_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Secret_Kind}.String()
	Secret_KindAPIVersion   = Secret_Kind + "." + CRDGroupVersion.String()
	Secret_GroupVersionKind = CRDGroupVersion.WithKind(Secret_Kind)
)

func init() {
	SchemeBuilder.Register(&Secret{}, &SecretList{})
}
