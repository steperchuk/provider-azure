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

type SQLPoolSecurityAlertPolicyObservation struct {

	// The ID of the Synapse SQL Pool Security Alert Policy.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SQLPoolSecurityAlertPolicyParameters struct {

	// Specifies an array of alerts that are disabled. Allowed values are: Sql_Injection, Sql_Injection_Vulnerability, Access_Anomaly, Data_Exfiltration, Unsafe_Action.
	// +kubebuilder:validation:Optional
	DisabledAlerts []*string `json:"disabledAlerts,omitempty" tf:"disabled_alerts,omitempty"`

	// Boolean flag which specifies if the alert is sent to the account administrators or not. Defaults to false.
	// +kubebuilder:validation:Optional
	EmailAccountAdminsEnabled *bool `json:"emailAccountAdminsEnabled,omitempty" tf:"email_account_admins_enabled,omitempty"`

	// Specifies an array of email addresses to which the alert is sent.
	// +kubebuilder:validation:Optional
	EmailAddresses []*string `json:"emailAddresses,omitempty" tf:"email_addresses,omitempty"`

	// Specifies the state of the policy, whether it is enabled or disabled or a policy has not been applied yet on the specific SQL pool. Possible values are Disabled, Enabled and New.
	// +kubebuilder:validation:Required
	PolicyState *string `json:"policyState" tf:"policy_state,omitempty"`

	// Specifies the number of days to keep in the Threat Detection audit logs. Defaults to 0.
	// +kubebuilder:validation:Optional
	RetentionDays *float64 `json:"retentionDays,omitempty" tf:"retention_days,omitempty"`

	// Specifies the ID of the Synapse SQL Pool. Changing this forces a new resource to be created.
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

	// Specifies the identifier key of the Threat Detection audit storage account.
	// +kubebuilder:validation:Optional
	StorageAccountAccessKeySecretRef *v1.SecretKeySelector `json:"storageAccountAccessKeySecretRef,omitempty" tf:"-"`

	// Specifies the blob storage endpoint (e.g. https://example.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs.
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

// SQLPoolSecurityAlertPolicySpec defines the desired state of SQLPoolSecurityAlertPolicy
type SQLPoolSecurityAlertPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SQLPoolSecurityAlertPolicyParameters `json:"forProvider"`
}

// SQLPoolSecurityAlertPolicyStatus defines the observed state of SQLPoolSecurityAlertPolicy.
type SQLPoolSecurityAlertPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SQLPoolSecurityAlertPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPoolSecurityAlertPolicy is the Schema for the SQLPoolSecurityAlertPolicys API. Manages a Security Alert Policy for a Synapse SQL Pool.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SQLPoolSecurityAlertPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SQLPoolSecurityAlertPolicySpec   `json:"spec"`
	Status            SQLPoolSecurityAlertPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SQLPoolSecurityAlertPolicyList contains a list of SQLPoolSecurityAlertPolicys
type SQLPoolSecurityAlertPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SQLPoolSecurityAlertPolicy `json:"items"`
}

// Repository type metadata.
var (
	SQLPoolSecurityAlertPolicy_Kind             = "SQLPoolSecurityAlertPolicy"
	SQLPoolSecurityAlertPolicy_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SQLPoolSecurityAlertPolicy_Kind}.String()
	SQLPoolSecurityAlertPolicy_KindAPIVersion   = SQLPoolSecurityAlertPolicy_Kind + "." + CRDGroupVersion.String()
	SQLPoolSecurityAlertPolicy_GroupVersionKind = CRDGroupVersion.WithKind(SQLPoolSecurityAlertPolicy_Kind)
)

func init() {
	SchemeBuilder.Register(&SQLPoolSecurityAlertPolicy{}, &SQLPoolSecurityAlertPolicyList{})
}
