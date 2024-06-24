// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	// ResolveReferences of this LB.
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func (mg *LB) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	if mg.Spec.ForProvider.AccessLogs != nil {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessLogs.Bucket),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.AccessLogs.BucketRef,
				Selector:     mg.Spec.ForProvider.AccessLogs.BucketSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.AccessLogs.Bucket")
		}
		mg.Spec.ForProvider.AccessLogs.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.AccessLogs.BucketRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.SecurityGroupRefs,
			Selector:      mg.Spec.ForProvider.SecurityGroupSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupRefs = mrsp.ResolvedReferences

	for i3 := 0; i3 < len(mg.Spec.ForProvider.SubnetMapping); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetMapping[i3].SubnetID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.SubnetMapping[i3].SubnetIDRef,
				Selector:     mg.Spec.ForProvider.SubnetMapping[i3].SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.SubnetMapping[i3].SubnetID")
		}
		mg.Spec.ForProvider.SubnetMapping[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.SubnetMapping[i3].SubnetIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Subnets),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.SubnetRefs,
			Selector:      mg.Spec.ForProvider.SubnetSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnets")
	}
	mg.Spec.ForProvider.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetRefs = mrsp.ResolvedReferences

	if mg.Spec.InitProvider.AccessLogs != nil {
		{
			m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AccessLogs.Bucket),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.AccessLogs.BucketRef,
				Selector:     mg.Spec.InitProvider.AccessLogs.BucketSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.AccessLogs.Bucket")
		}
		mg.Spec.InitProvider.AccessLogs.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.AccessLogs.BucketRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "SecurityGroup", "SecurityGroupList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.SecurityGroups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.SecurityGroupRefs,
			Selector:      mg.Spec.InitProvider.SecurityGroupSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SecurityGroups")
	}
	mg.Spec.InitProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SecurityGroupRefs = mrsp.ResolvedReferences

	for i3 := 0; i3 < len(mg.Spec.InitProvider.SubnetMapping); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SubnetMapping[i3].SubnetID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.SubnetMapping[i3].SubnetIDRef,
				Selector:     mg.Spec.InitProvider.SubnetMapping[i3].SubnetIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.SubnetMapping[i3].SubnetID")
		}
		mg.Spec.InitProvider.SubnetMapping[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.SubnetMapping[i3].SubnetIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "Subnet", "SubnetList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.Subnets),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.InitProvider.SubnetRefs,
			Selector:      mg.Spec.InitProvider.SubnetSelector,
			To:            reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Subnets")
	}
	mg.Spec.InitProvider.Subnets = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.InitProvider.SubnetRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this LBListener.
func (mg *LBListener) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.DefaultAction); i3++ {
		if mg.Spec.ForProvider.DefaultAction[i3].Forward != nil {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.DefaultAction[i3].Forward.TargetGroup); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LBTargetGroup", "LBTargetGroupList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultAction[i3].Forward.TargetGroup[i5].Arn),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.ForProvider.DefaultAction[i3].Forward.TargetGroup[i5].ArnRef,
						Selector:     mg.Spec.ForProvider.DefaultAction[i3].Forward.TargetGroup[i5].ArnSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.DefaultAction[i3].Forward.TargetGroup[i5].Arn")
				}
				mg.Spec.ForProvider.DefaultAction[i3].Forward.TargetGroup[i5].Arn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.DefaultAction[i3].Forward.TargetGroup[i5].ArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.DefaultAction); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LBTargetGroup", "LBTargetGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArn),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArnRef,
				Selector:     mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArn")
		}
		mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.DefaultAction[i3].TargetGroupArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LB", "LBList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.LoadBalancerArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.LoadBalancerArnRef,
			Selector:     mg.Spec.ForProvider.LoadBalancerArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.LoadBalancerArn")
	}
	mg.Spec.ForProvider.LoadBalancerArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.LoadBalancerArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.DefaultAction); i3++ {
		if mg.Spec.InitProvider.DefaultAction[i3].Forward != nil {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.DefaultAction[i3].Forward.TargetGroup); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LBTargetGroup", "LBTargetGroupList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DefaultAction[i3].Forward.TargetGroup[i5].Arn),
						Extract:      reference.ExternalName(),
						Reference:    mg.Spec.InitProvider.DefaultAction[i3].Forward.TargetGroup[i5].ArnRef,
						Selector:     mg.Spec.InitProvider.DefaultAction[i3].Forward.TargetGroup[i5].ArnSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.DefaultAction[i3].Forward.TargetGroup[i5].Arn")
				}
				mg.Spec.InitProvider.DefaultAction[i3].Forward.TargetGroup[i5].Arn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.DefaultAction[i3].Forward.TargetGroup[i5].ArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.DefaultAction); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LBTargetGroup", "LBTargetGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.DefaultAction[i3].TargetGroupArn),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.InitProvider.DefaultAction[i3].TargetGroupArnRef,
				Selector:     mg.Spec.InitProvider.DefaultAction[i3].TargetGroupArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.DefaultAction[i3].TargetGroupArn")
		}
		mg.Spec.InitProvider.DefaultAction[i3].TargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.DefaultAction[i3].TargetGroupArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LB", "LBList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.LoadBalancerArn),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.LoadBalancerArnRef,
			Selector:     mg.Spec.InitProvider.LoadBalancerArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.LoadBalancerArn")
	}
	mg.Spec.InitProvider.LoadBalancerArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.LoadBalancerArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LBListenerRule.
func (mg *LBListenerRule) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Action); i3++ {
		if mg.Spec.ForProvider.Action[i3].AuthenticateCognito != nil {
			{
				m, l, err = apisresolver.GetManagedResource("cognitoidp.aws.upbound.io", "v1beta2", "UserPool", "UserPoolList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolArnRef,
					Selector:     mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolArn")
			}
			mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Action); i3++ {
		if mg.Spec.ForProvider.Action[i3].AuthenticateCognito != nil {
			{
				m, l, err = apisresolver.GetManagedResource("cognitoidp.aws.upbound.io", "v1beta1", "UserPoolClient", "UserPoolClientList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolClientID),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolClientIDRef,
					Selector:     mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolClientIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolClientID")
			}
			mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolClientID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolClientIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Action); i3++ {
		if mg.Spec.ForProvider.Action[i3].AuthenticateCognito != nil {
			{
				m, l, err = apisresolver.GetManagedResource("cognitoidp.aws.upbound.io", "v1beta1", "UserPoolDomain", "UserPoolDomainList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolDomain),
					Extract:      resource.ExtractParamPath("domain", false),
					Reference:    mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolDomainRef,
					Selector:     mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolDomainSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolDomain")
			}
			mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolDomain = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Action[i3].AuthenticateCognito.UserPoolDomainRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Action); i3++ {
		if mg.Spec.ForProvider.Action[i3].Forward != nil {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Action[i3].Forward.TargetGroup); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LBTargetGroup", "LBTargetGroupList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Action[i3].Forward.TargetGroup[i5].Arn),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.ForProvider.Action[i3].Forward.TargetGroup[i5].ArnRef,
						Selector:     mg.Spec.ForProvider.Action[i3].Forward.TargetGroup[i5].ArnSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.Action[i3].Forward.TargetGroup[i5].Arn")
				}
				mg.Spec.ForProvider.Action[i3].Forward.TargetGroup[i5].Arn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.Action[i3].Forward.TargetGroup[i5].ArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Action); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LBTargetGroup", "LBTargetGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Action[i3].TargetGroupArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Action[i3].TargetGroupArnRef,
				Selector:     mg.Spec.ForProvider.Action[i3].TargetGroupArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Action[i3].TargetGroupArn")
		}
		mg.Spec.ForProvider.Action[i3].TargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Action[i3].TargetGroupArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LBListener", "LBListenerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ListenerArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.ListenerArnRef,
			Selector:     mg.Spec.ForProvider.ListenerArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ListenerArn")
	}
	mg.Spec.ForProvider.ListenerArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ListenerArnRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.Action); i3++ {
		if mg.Spec.InitProvider.Action[i3].AuthenticateCognito != nil {
			{
				m, l, err = apisresolver.GetManagedResource("cognitoidp.aws.upbound.io", "v1beta2", "UserPool", "UserPoolList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolArn),
					Extract:      resource.ExtractParamPath("arn", true),
					Reference:    mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolArnRef,
					Selector:     mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolArnSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolArn")
			}
			mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolArn = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolArnRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Action); i3++ {
		if mg.Spec.InitProvider.Action[i3].AuthenticateCognito != nil {
			{
				m, l, err = apisresolver.GetManagedResource("cognitoidp.aws.upbound.io", "v1beta1", "UserPoolClient", "UserPoolClientList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolClientID),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolClientIDRef,
					Selector:     mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolClientIDSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolClientID")
			}
			mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolClientID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolClientIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Action); i3++ {
		if mg.Spec.InitProvider.Action[i3].AuthenticateCognito != nil {
			{
				m, l, err = apisresolver.GetManagedResource("cognitoidp.aws.upbound.io", "v1beta1", "UserPoolDomain", "UserPoolDomainList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolDomain),
					Extract:      resource.ExtractParamPath("domain", false),
					Reference:    mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolDomainRef,
					Selector:     mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolDomainSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolDomain")
			}
			mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolDomain = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.InitProvider.Action[i3].AuthenticateCognito.UserPoolDomainRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Action); i3++ {
		if mg.Spec.InitProvider.Action[i3].Forward != nil {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Action[i3].Forward.TargetGroup); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LBTargetGroup", "LBTargetGroupList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Action[i3].Forward.TargetGroup[i5].Arn),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.InitProvider.Action[i3].Forward.TargetGroup[i5].ArnRef,
						Selector:     mg.Spec.InitProvider.Action[i3].Forward.TargetGroup[i5].ArnSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.Action[i3].Forward.TargetGroup[i5].Arn")
				}
				mg.Spec.InitProvider.Action[i3].Forward.TargetGroup[i5].Arn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.Action[i3].Forward.TargetGroup[i5].ArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Action); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LBTargetGroup", "LBTargetGroupList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Action[i3].TargetGroupArn),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.Action[i3].TargetGroupArnRef,
				Selector:     mg.Spec.InitProvider.Action[i3].TargetGroupArnSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Action[i3].TargetGroupArn")
		}
		mg.Spec.InitProvider.Action[i3].TargetGroupArn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Action[i3].TargetGroupArnRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("elbv2.aws.upbound.io", "v1beta2", "LBListener", "LBListenerList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.ListenerArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.ListenerArnRef,
			Selector:     mg.Spec.InitProvider.ListenerArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.ListenerArn")
	}
	mg.Spec.InitProvider.ListenerArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.ListenerArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this LBTargetGroup.
func (mg *LBTargetGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.VPCIDRef,
			Selector:     mg.Spec.ForProvider.VPCIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VPCID")
	}
	mg.Spec.ForProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VPCIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.InitProvider.VPCIDRef,
			Selector:     mg.Spec.InitProvider.VPCIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.VPCID")
	}
	mg.Spec.InitProvider.VPCID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.VPCIDRef = rsp.ResolvedReference

	return nil
}