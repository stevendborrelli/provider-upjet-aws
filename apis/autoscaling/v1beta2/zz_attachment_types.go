// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AttachmentInitParameters struct {

	// Name of ASG to associate with the ELB.
	// +crossplane:generate:reference:type=AutoscalingGroup
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	// Reference to a AutoscalingGroup to populate autoscalingGroupName.
	// +kubebuilder:validation:Optional
	AutoscalingGroupNameRef *v1.Reference `json:"autoscalingGroupNameRef,omitempty" tf:"-"`

	// Selector for a AutoscalingGroup to populate autoscalingGroupName.
	// +kubebuilder:validation:Optional
	AutoscalingGroupNameSelector *v1.Selector `json:"autoscalingGroupNameSelector,omitempty" tf:"-"`

	// Name of the ELB.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ELB *string `json:"elb,omitempty" tf:"elb,omitempty"`

	// Reference to a ELB in elb to populate elb.
	// +kubebuilder:validation:Optional
	ELBRef *v1.Reference `json:"elbRef,omitempty" tf:"-"`

	// Selector for a ELB in elb to populate elb.
	// +kubebuilder:validation:Optional
	ELBSelector *v1.Selector `json:"elbSelector,omitempty" tf:"-"`

	// ARN of a load balancer target group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBTargetGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	LBTargetGroupArn *string `json:"lbTargetGroupArn,omitempty" tf:"lb_target_group_arn,omitempty"`

	// Reference to a LBTargetGroup in elbv2 to populate lbTargetGroupArn.
	// +kubebuilder:validation:Optional
	LBTargetGroupArnRef *v1.Reference `json:"lbTargetGroupArnRef,omitempty" tf:"-"`

	// Selector for a LBTargetGroup in elbv2 to populate lbTargetGroupArn.
	// +kubebuilder:validation:Optional
	LBTargetGroupArnSelector *v1.Selector `json:"lbTargetGroupArnSelector,omitempty" tf:"-"`
}

type AttachmentObservation struct {

	// Name of ASG to associate with the ELB.
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	// Name of the ELB.
	ELB *string `json:"elb,omitempty" tf:"elb,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ARN of a load balancer target group.
	LBTargetGroupArn *string `json:"lbTargetGroupArn,omitempty" tf:"lb_target_group_arn,omitempty"`
}

type AttachmentParameters struct {

	// Name of ASG to associate with the ELB.
	// +crossplane:generate:reference:type=AutoscalingGroup
	// +kubebuilder:validation:Optional
	AutoscalingGroupName *string `json:"autoscalingGroupName,omitempty" tf:"autoscaling_group_name,omitempty"`

	// Reference to a AutoscalingGroup to populate autoscalingGroupName.
	// +kubebuilder:validation:Optional
	AutoscalingGroupNameRef *v1.Reference `json:"autoscalingGroupNameRef,omitempty" tf:"-"`

	// Selector for a AutoscalingGroup to populate autoscalingGroupName.
	// +kubebuilder:validation:Optional
	AutoscalingGroupNameSelector *v1.Selector `json:"autoscalingGroupNameSelector,omitempty" tf:"-"`

	// Name of the ELB.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elb/v1beta1.ELB
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ELB *string `json:"elb,omitempty" tf:"elb,omitempty"`

	// Reference to a ELB in elb to populate elb.
	// +kubebuilder:validation:Optional
	ELBRef *v1.Reference `json:"elbRef,omitempty" tf:"-"`

	// Selector for a ELB in elb to populate elb.
	// +kubebuilder:validation:Optional
	ELBSelector *v1.Selector `json:"elbSelector,omitempty" tf:"-"`

	// ARN of a load balancer target group.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBTargetGroup
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	LBTargetGroupArn *string `json:"lbTargetGroupArn,omitempty" tf:"lb_target_group_arn,omitempty"`

	// Reference to a LBTargetGroup in elbv2 to populate lbTargetGroupArn.
	// +kubebuilder:validation:Optional
	LBTargetGroupArnRef *v1.Reference `json:"lbTargetGroupArnRef,omitempty" tf:"-"`

	// Selector for a LBTargetGroup in elbv2 to populate lbTargetGroupArn.
	// +kubebuilder:validation:Optional
	LBTargetGroupArnSelector *v1.Selector `json:"lbTargetGroupArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// AttachmentSpec defines the desired state of Attachment
type AttachmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AttachmentParameters `json:"forProvider"`
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
	InitProvider AttachmentInitParameters `json:"initProvider,omitempty"`
}

// AttachmentStatus defines the observed state of Attachment.
type AttachmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AttachmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Attachment is the Schema for the Attachments API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Attachment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AttachmentSpec   `json:"spec"`
	Status            AttachmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AttachmentList contains a list of Attachments
type AttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Attachment `json:"items"`
}

// Repository type metadata.
var (
	Attachment_Kind             = "Attachment"
	Attachment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Attachment_Kind}.String()
	Attachment_KindAPIVersion   = Attachment_Kind + "." + CRDGroupVersion.String()
	Attachment_GroupVersionKind = CRDGroupVersion.WithKind(Attachment_Kind)
)

func init() {
	SchemeBuilder.Register(&Attachment{}, &AttachmentList{})
}