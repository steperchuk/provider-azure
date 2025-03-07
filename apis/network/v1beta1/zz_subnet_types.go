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

type DelegationObservation struct {
}

type DelegationParameters struct {

	// A name for this delegation.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// A service_delegation block as defined below.
	// +kubebuilder:validation:Required
	ServiceDelegation []ServiceDelegationParameters `json:"serviceDelegation" tf:"service_delegation,omitempty"`
}

type ServiceDelegationObservation struct {
}

type ServiceDelegationParameters struct {

	// A list of Actions which should be delegated. This list is specific to the service to delegate to. Possible values are Microsoft.Network/networkinterfaces/*, Microsoft.Network/publicIPAddresses/join/action, Microsoft.Network/publicIPAddresses/read, Microsoft.Network/virtualNetworks/read, Microsoft.Network/virtualNetworks/subnets/action, Microsoft.Network/virtualNetworks/subnets/join/action, Microsoft.Network/virtualNetworks/subnets/prepareNetworkPolicies/action, and Microsoft.Network/virtualNetworks/subnets/unprepareNetworkPolicies/action.
	// +kubebuilder:validation:Optional
	Actions []*string `json:"actions,omitempty" tf:"actions,omitempty"`

	// The name of service to delegate to. Possible values are Microsoft.ApiManagement/service, Microsoft.AzureCosmosDB/clusters, Microsoft.BareMetal/AzureVMware, Microsoft.BareMetal/CrayServers, Microsoft.Batch/batchAccounts, Microsoft.ContainerInstance/containerGroups, Microsoft.ContainerService/managedClusters, Microsoft.Databricks/workspaces, Microsoft.DBforMySQL/flexibleServers, Microsoft.DBforMySQL/serversv2, Microsoft.DBforPostgreSQL/flexibleServers, Microsoft.DBforPostgreSQL/serversv2, Microsoft.DBforPostgreSQL/singleServers, Microsoft.HardwareSecurityModules/dedicatedHSMs, Microsoft.Kusto/clusters, Microsoft.Logic/integrationServiceEnvironments, Microsoft.LabServices/labplans, Microsoft.MachineLearningServices/workspaces, Microsoft.Netapp/volumes, Microsoft.Network/dnsResolvers, Microsoft.Network/managedResolvers, Microsoft.PowerPlatform/vnetaccesslinks, Microsoft.ServiceFabricMesh/networks, Microsoft.Sql/managedInstances, Microsoft.Sql/servers, Microsoft.StoragePool/diskPools, Microsoft.StreamAnalytics/streamingJobs, Microsoft.Synapse/workspaces, Microsoft.Web/hostingEnvironments, Microsoft.Web/serverFarms, Microsoft.Orbital/orbitalGateways, NGINX.NGINXPLUS/nginxDeployments, PaloAltoNetworks.Cloudngfw/firewalls, and Qumulo.Storage/fileSystems.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type SubnetObservation struct {

	// The subnet ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SubnetParameters struct {

	// The address prefixes to use for the subnet.
	// +kubebuilder:validation:Required
	AddressPrefixes []*string `json:"addressPrefixes" tf:"address_prefixes,omitempty"`

	// One or more delegation blocks as defined below.
	// +kubebuilder:validation:Optional
	Delegation []DelegationParameters `json:"delegation,omitempty" tf:"delegation,omitempty"`

	// +kubebuilder:validation:Optional
	EnforcePrivateLinkEndpointNetworkPolicies *bool `json:"enforcePrivateLinkEndpointNetworkPolicies,omitempty" tf:"enforce_private_link_endpoint_network_policies,omitempty"`

	// +kubebuilder:validation:Optional
	EnforcePrivateLinkServiceNetworkPolicies *bool `json:"enforcePrivateLinkServiceNetworkPolicies,omitempty" tf:"enforce_private_link_service_network_policies,omitempty"`

	// Enable or Disable network policies for the private endpoint on the subnet. Setting this to true will Enable the policy and setting this to false will Disable the policy. Defaults to true.
	// +kubebuilder:validation:Optional
	PrivateEndpointNetworkPoliciesEnabled *bool `json:"privateEndpointNetworkPoliciesEnabled,omitempty" tf:"private_endpoint_network_policies_enabled,omitempty"`

	// Enable or Disable network policies for the private link service on the subnet. Setting this to true will Enable the policy and setting this to false will Disable the policy. Defaults to true.
	// +kubebuilder:validation:Optional
	PrivateLinkServiceNetworkPoliciesEnabled *bool `json:"privateLinkServiceNetworkPoliciesEnabled,omitempty" tf:"private_link_service_network_policies_enabled,omitempty"`

	// The name of the resource group in which to create the subnet. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The list of IDs of Service Endpoint Policies to associate with the subnet.
	// +kubebuilder:validation:Optional
	ServiceEndpointPolicyIds []*string `json:"serviceEndpointPolicyIds,omitempty" tf:"service_endpoint_policy_ids,omitempty"`

	// The list of Service endpoints to associate with the subnet. Possible values include: Microsoft.AzureActiveDirectory, Microsoft.AzureCosmosDB, Microsoft.ContainerRegistry, Microsoft.EventHub, Microsoft.KeyVault, Microsoft.ServiceBus, Microsoft.Sql, Microsoft.Storage, and Microsoft.Web.
	// +kubebuilder:validation:Optional
	ServiceEndpoints []*string `json:"serviceEndpoints,omitempty" tf:"service_endpoints,omitempty"`

	// The name of the virtual network to which to attach the subnet. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=VirtualNetwork
	// +kubebuilder:validation:Optional
	VirtualNetworkName *string `json:"virtualNetworkName,omitempty" tf:"virtual_network_name,omitempty"`

	// Reference to a VirtualNetwork to populate virtualNetworkName.
	// +kubebuilder:validation:Optional
	VirtualNetworkNameRef *v1.Reference `json:"virtualNetworkNameRef,omitempty" tf:"-"`

	// Selector for a VirtualNetwork to populate virtualNetworkName.
	// +kubebuilder:validation:Optional
	VirtualNetworkNameSelector *v1.Selector `json:"virtualNetworkNameSelector,omitempty" tf:"-"`
}

// SubnetSpec defines the desired state of Subnet
type SubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetParameters `json:"forProvider"`
}

// SubnetStatus defines the observed state of Subnet.
type SubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subnet is the Schema for the Subnets API. Manages a subnet. Subnets represent network segments within the IP space defined by the virtual network.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetSpec   `json:"spec"`
	Status            SubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetList contains a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

// Repository type metadata.
var (
	Subnet_Kind             = "Subnet"
	Subnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subnet_Kind}.String()
	Subnet_KindAPIVersion   = Subnet_Kind + "." + CRDGroupVersion.String()
	Subnet_GroupVersionKind = CRDGroupVersion.WithKind(Subnet_Kind)
)

func init() {
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
}
