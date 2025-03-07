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

type SQLPoolExtendedAuditingPolicyObservation struct {

	// The ID of the Synapse SQL Pool Extended Auditing Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SQLPoolExtendedAuditingPolicyParameters struct {

	// Enable audit events to Azure Monitor? To enable server audit events to Azure Monitor, please enable its master database audit events to Azure Monitor. Defaults to true.
	// +kubebuilder:validation:Optional
	LogMonitoringEnabled *bool `json:"logMonitoringEnabled,omitempty" tf:"log_monitoring_enabled,omitempty"`

	// The number of days to retain logs for in the storage account. Defaults to 0.
	// +kubebuilder:validation:Optional
	RetentionInDays *float64 `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`

	// The ID of the Synapse SQL pool to set the extended auditing policy. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/synapse/v1beta1.SQLPool
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SQLPoolID *string `json:"sqlPoolId,omitempty" tf:"sql_pool_id,omitempty"`

	// Reference to a SQLPool in synapse to populate sqlPoolId.
	// +kubebuilder:validation:Optional
	SQLPoolIDRef *v1.Reference `json:"sqlPoolIdRef,omitempty" tf:"-"`

	// Selector for a SQLPool in synapse to populate sqlPoolId.
	// +kubebuilder:validation:Optional
	SQLPoolIDSelector *v1.Selector `json:"sqlPoolIdSelector,omitempty" tf:"-"`

	// Is storage_account_access_key value the storage's secondary key?
	// +kubebuilder:validation:Optional
	StorageAccountAccessKeyIsSecondary *bool `json:"storageAccountAccessKeyIsSecondary,omitempty" tf:"storage_account_access_key_is_secondary,omitempty"`

	// The access key to use for the auditing storage account.
	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// The blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all extended auditing logs.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Account
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("primary_blob_endpoint",true)
	// +kubebuilder:validation:Optional
	StorageEndpoint *string `json:"storageEndpoint,omitempty" tf:"storage_endpoint,omitempty"`

	// Reference to a Account in storage to populate storageEndpoint.
	// +kubebuilder:validation:Optional
	StorageEndpointRef *v1.Reference `json:"storageEndpointRef,omitempty" tf:"-"`

	// Selector for a Account in storage to populate storageEndpoint.
	// +kubebuilder:validation:Optional
	StorageEndpointSelector *v1.Selector `json:"storageEndpointSelector,omitempty" tf:"-"`
}

// SQLPoolExtendedAuditingPolicySpec defines the desired state of SQLPoolExtendedAuditingPolicy
type SQLPoolExtendedAuditingPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLPoolExtendedAuditingPolicyParameters `json:"forProvider"`
}

// SQLPoolExtendedAuditingPolicyStatus defines the observed state of SQLPoolExtendedAuditingPolicy.
type SQLPoolExtendedAuditingPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLPoolExtendedAuditingPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPoolExtendedAuditingPolicy is the Schema for the SQLPoolExtendedAuditingPolicys API. Manages a Synapse SQL Pool Extended Auditing Policy.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLPoolExtendedAuditingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SQLPoolExtendedAuditingPolicySpec   `json:"spec"`
	Status            SQLPoolExtendedAuditingPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPoolExtendedAuditingPolicyList contains a list of SQLPoolExtendedAuditingPolicys
type SQLPoolExtendedAuditingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLPoolExtendedAuditingPolicy `json:"items"`
}

// Repository type metadata.
var (
	SQLPoolExtendedAuditingPolicy_Kind             = "SQLPoolExtendedAuditingPolicy"
	SQLPoolExtendedAuditingPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLPoolExtendedAuditingPolicy_Kind}.String()
	SQLPoolExtendedAuditingPolicy_KindAPIVersion   = SQLPoolExtendedAuditingPolicy_Kind + "." + CRDGroupVersion.String()
	SQLPoolExtendedAuditingPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SQLPoolExtendedAuditingPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLPoolExtendedAuditingPolicy{}, &SQLPoolExtendedAuditingPolicyList{})
}
