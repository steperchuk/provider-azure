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

type SentinelLogAnalyticsWorkspaceOnboardingObservation struct {

	// The ID of the Security Insights Sentinel Onboarding States.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SentinelLogAnalyticsWorkspaceOnboardingParameters struct {

	// Specifies if the Workspace is using Customer managed key. Defaults to false.
	// +kubebuilder:validation:Optional
	CustomerManagedKeyEnabled *bool `json:"customerManagedKeyEnabled,omitempty" tf:"customer_managed_key_enabled,omitempty"`

	// Specifies the name of the Resource Group where the Security Insights Sentinel Onboarding States should exist. Changing this forces the Log Analytics Workspace off the board and onboard again.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// Specifies the Workspace Name. Changing this forces the Log Analytics Workspace off the board and onboard again.
	// +kubebuilder:validation:Required
	WorkspaceName *string `json:"workspaceName" tf:"workspace_name,omitempty"`
}

// SentinelLogAnalyticsWorkspaceOnboardingSpec defines the desired state of SentinelLogAnalyticsWorkspaceOnboarding
type SentinelLogAnalyticsWorkspaceOnboardingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SentinelLogAnalyticsWorkspaceOnboardingParameters `json:"forProvider"`
}

// SentinelLogAnalyticsWorkspaceOnboardingStatus defines the observed state of SentinelLogAnalyticsWorkspaceOnboarding.
type SentinelLogAnalyticsWorkspaceOnboardingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SentinelLogAnalyticsWorkspaceOnboardingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelLogAnalyticsWorkspaceOnboarding is the Schema for the SentinelLogAnalyticsWorkspaceOnboardings API. Manages a Security Insights Sentinel Onboarding States.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SentinelLogAnalyticsWorkspaceOnboarding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SentinelLogAnalyticsWorkspaceOnboardingSpec   `json:"spec"`
	Status            SentinelLogAnalyticsWorkspaceOnboardingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SentinelLogAnalyticsWorkspaceOnboardingList contains a list of SentinelLogAnalyticsWorkspaceOnboardings
type SentinelLogAnalyticsWorkspaceOnboardingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SentinelLogAnalyticsWorkspaceOnboarding `json:"items"`
}

// Repository type metadata.
var (
	SentinelLogAnalyticsWorkspaceOnboarding_Kind             = "SentinelLogAnalyticsWorkspaceOnboarding"
	SentinelLogAnalyticsWorkspaceOnboarding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SentinelLogAnalyticsWorkspaceOnboarding_Kind}.String()
	SentinelLogAnalyticsWorkspaceOnboarding_KindAPIVersion   = SentinelLogAnalyticsWorkspaceOnboarding_Kind + "." + CRDGroupVersion.String()
	SentinelLogAnalyticsWorkspaceOnboarding_GroupVersionKind = CRDGroupVersion.WithKind(SentinelLogAnalyticsWorkspaceOnboarding_Kind)
)

func init() {
	SchemeBuilder.Register(&SentinelLogAnalyticsWorkspaceOnboarding{}, &SentinelLogAnalyticsWorkspaceOnboardingList{})
}
