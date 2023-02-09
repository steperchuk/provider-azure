/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-azure/apis/azure/v1beta1"
	v1beta11 "github.com/upbound/provider-azure/apis/network/v1beta1"
	rconfig "github.com/upbound/provider-azure/apis/rconfig"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this AuthorizationRule.
func (mg *AuthorizationRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHubName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EventHubNameRef,
		Selector:     mg.Spec.ForProvider.EventHubNameSelector,
		To: reference.To{
			List:    &EventHubList{},
			Managed: &EventHub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventHubName")
	}
	mg.Spec.ForProvider.EventHubName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventHubNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NamespaceNameRef,
		Selector:     mg.Spec.ForProvider.NamespaceNameSelector,
		To: reference.To{
			List:    &EventHubNamespaceList{},
			Managed: &EventHubNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceName")
	}
	mg.Spec.ForProvider.NamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ConsumerGroup.
func (mg *ConsumerGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EventHubName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.EventHubNameRef,
		Selector:     mg.Spec.ForProvider.EventHubNameSelector,
		To: reference.To{
			List:    &EventHubList{},
			Managed: &EventHub{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EventHubName")
	}
	mg.Spec.ForProvider.EventHubName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EventHubNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NamespaceNameRef,
		Selector:     mg.Spec.ForProvider.NamespaceNameSelector,
		To: reference.To{
			List:    &EventHubNamespaceList{},
			Managed: &EventHubNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceName")
	}
	mg.Spec.ForProvider.NamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EventHub.
func (mg *EventHub) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NamespaceNameRef,
		Selector:     mg.Spec.ForProvider.NamespaceNameSelector,
		To: reference.To{
			List:    &EventHubNamespaceList{},
			Managed: &EventHubNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceName")
	}
	mg.Spec.ForProvider.NamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this EventHubNamespace.
func (mg *EventHubNamespace) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkRulesets); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.NetworkRulesets[i3].VirtualNetworkRule); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkRulesets[i3].VirtualNetworkRule[i4].SubnetID),
				Extract:      rconfig.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.NetworkRulesets[i3].VirtualNetworkRule[i4].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.NetworkRulesets[i3].VirtualNetworkRule[i4].SubnetIDSelector,
				To: reference.To{
					List:    &v1beta11.SubnetList{},
					Managed: &v1beta11.Subnet{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.NetworkRulesets[i3].VirtualNetworkRule[i4].SubnetID")
			}
			mg.Spec.ForProvider.NetworkRulesets[i3].VirtualNetworkRule[i4].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.NetworkRulesets[i3].VirtualNetworkRule[i4].SubnetIDRef = rsp.ResolvedReference

		}
	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NamespaceAuthorizationRule.
func (mg *NamespaceAuthorizationRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NamespaceNameRef,
		Selector:     mg.Spec.ForProvider.NamespaceNameSelector,
		To: reference.To{
			List:    &EventHubNamespaceList{},
			Managed: &EventHubNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceName")
	}
	mg.Spec.ForProvider.NamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NamespaceDisasterRecoveryConfig.
func (mg *NamespaceDisasterRecoveryConfig) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NamespaceNameRef,
		Selector:     mg.Spec.ForProvider.NamespaceNameSelector,
		To: reference.To{
			List:    &EventHubNamespaceList{},
			Managed: &EventHubNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceName")
	}
	mg.Spec.ForProvider.NamespaceName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PartnerNamespaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PartnerNamespaceIDRef,
		Selector:     mg.Spec.ForProvider.PartnerNamespaceIDSelector,
		To: reference.To{
			List:    &EventHubNamespaceList{},
			Managed: &EventHubNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PartnerNamespaceID")
	}
	mg.Spec.ForProvider.PartnerNamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PartnerNamespaceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceGroupName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceGroupNameRef,
		Selector:     mg.Spec.ForProvider.ResourceGroupNameSelector,
		To: reference.To{
			List:    &v1beta1.ResourceGroupList{},
			Managed: &v1beta1.ResourceGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceGroupName")
	}
	mg.Spec.ForProvider.ResourceGroupName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceGroupNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NamespaceSchemaGroup.
func (mg *NamespaceSchemaGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NamespaceID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.NamespaceIDRef,
		Selector:     mg.Spec.ForProvider.NamespaceIDSelector,
		To: reference.To{
			List:    &EventHubNamespaceList{},
			Managed: &EventHubNamespace{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.NamespaceID")
	}
	mg.Spec.ForProvider.NamespaceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NamespaceIDRef = rsp.ResolvedReference

	return nil
}
