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

type GalleryApplicationVersionObservation struct {

	// The ID of the Gallery Application Version.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GalleryApplicationVersionParameters struct {

	// Should the Gallery Application reports health. Defaults to false.
	// +kubebuilder:validation:Optional
	EnableHealthCheck *bool `json:"enableHealthCheck,omitempty" tf:"enable_health_check,omitempty"`

	// The end of life date in RFC3339 format of the Gallery Application Version.
	// +kubebuilder:validation:Optional
	EndOfLifeDate *string `json:"endOfLifeDate,omitempty" tf:"end_of_life_date,omitempty"`

	// Should the Gallery Application Version be excluded from the latest filter? If set to true this Gallery Application Version won't be returned for the latest version. Defaults to false.
	// +kubebuilder:validation:Optional
	ExcludeFromLatest *bool `json:"excludeFromLatest,omitempty" tf:"exclude_from_latest,omitempty"`

	// The ID of the Gallery Application. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.GalleryApplication
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GalleryApplicationID *string `json:"galleryApplicationId,omitempty" tf:"gallery_application_id,omitempty"`

	// Reference to a GalleryApplication in compute to populate galleryApplicationId.
	// +kubebuilder:validation:Optional
	GalleryApplicationIDRef *v1.Reference `json:"galleryApplicationIdRef,omitempty" tf:"-"`

	// Selector for a GalleryApplication in compute to populate galleryApplicationId.
	// +kubebuilder:validation:Optional
	GalleryApplicationIDSelector *v1.Selector `json:"galleryApplicationIdSelector,omitempty" tf:"-"`

	// The Azure Region where the Gallery Application Version exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// A manage_action block as defined below.
	// +kubebuilder:validation:Required
	ManageAction []ManageActionParameters `json:"manageAction" tf:"manage_action,omitempty"`

	// The version name of the Gallery Application Version, such as 1.0.0. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A source block as defined below.
	// +kubebuilder:validation:Required
	Source []SourceParameters `json:"source" tf:"source,omitempty"`

	// A mapping of tags to assign to the Gallery Application Version.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// One or more target_region blocks as defined below.
	// +kubebuilder:validation:Required
	TargetRegion []TargetRegionParameters `json:"targetRegion" tf:"target_region,omitempty"`
}

type ManageActionObservation struct {
}

type ManageActionParameters struct {

	// The command to install the Gallery Application. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Install *string `json:"install" tf:"install,omitempty"`

	// The command to remove the Gallery Application. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Required
	Remove *string `json:"remove" tf:"remove,omitempty"`

	// The command to update the Gallery Application. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Update *string `json:"update,omitempty" tf:"update,omitempty"`
}

type SourceObservation struct {
}

type SourceParameters struct {

	// The Storage Blob URI of the default configuration. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	DefaultConfigurationLink *string `json:"defaultConfigurationLink,omitempty" tf:"default_configuration_link,omitempty"`

	// The Storage Blob URI of the source application package. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Blob
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	MediaLink *string `json:"mediaLink,omitempty" tf:"media_link,omitempty"`

	// Reference to a Blob in storage to populate mediaLink.
	// +kubebuilder:validation:Optional
	MediaLinkRef *v1.Reference `json:"mediaLinkRef,omitempty" tf:"-"`

	// Selector for a Blob in storage to populate mediaLink.
	// +kubebuilder:validation:Optional
	MediaLinkSelector *v1.Selector `json:"mediaLinkSelector,omitempty" tf:"-"`
}

type TargetRegionObservation struct {
}

type TargetRegionParameters struct {

	// The Azure Region in which the Gallery Application Version exists.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.GalleryApplication
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("location",false)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Reference to a GalleryApplication in compute to populate name.
	// +kubebuilder:validation:Optional
	NameRef *v1.Reference `json:"nameRef,omitempty" tf:"-"`

	// Selector for a GalleryApplication in compute to populate name.
	// +kubebuilder:validation:Optional
	NameSelector *v1.Selector `json:"nameSelector,omitempty" tf:"-"`

	// The number of replicas of the Gallery Application Version to be created per region. Possible values are between 1 and 10.
	// +kubebuilder:validation:Required
	RegionalReplicaCount *float64 `json:"regionalReplicaCount" tf:"regional_replica_count,omitempty"`

	// The storage account type for the Gallery Application Version. Possible values are Standard_LRS, Premium_LRS and Standard_ZRS. Defaults to Standard_LRS.
	// +kubebuilder:validation:Optional
	StorageAccountType *string `json:"storageAccountType,omitempty" tf:"storage_account_type,omitempty"`
}

// GalleryApplicationVersionSpec defines the desired state of GalleryApplicationVersion
type GalleryApplicationVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GalleryApplicationVersionParameters `json:"forProvider"`
}

// GalleryApplicationVersionStatus defines the observed state of GalleryApplicationVersion.
type GalleryApplicationVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GalleryApplicationVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GalleryApplicationVersion is the Schema for the GalleryApplicationVersions API. Manages a Gallery Application Version.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type GalleryApplicationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GalleryApplicationVersionSpec   `json:"spec"`
	Status            GalleryApplicationVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GalleryApplicationVersionList contains a list of GalleryApplicationVersions
type GalleryApplicationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GalleryApplicationVersion `json:"items"`
}

// Repository type metadata.
var (
	GalleryApplicationVersion_Kind             = "GalleryApplicationVersion"
	GalleryApplicationVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GalleryApplicationVersion_Kind}.String()
	GalleryApplicationVersion_KindAPIVersion   = GalleryApplicationVersion_Kind + "." + CRDGroupVersion.String()
	GalleryApplicationVersion_GroupVersionKind = CRDGroupVersion.WithKind(GalleryApplicationVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&GalleryApplicationVersion{}, &GalleryApplicationVersionList{})
}
