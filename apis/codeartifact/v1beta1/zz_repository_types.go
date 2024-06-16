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

type ExternalConnectionsInitParameters struct {

	// The name of the external connection associated with a repository.
	ExternalConnectionName *string `json:"externalConnectionName,omitempty" tf:"external_connection_name,omitempty"`
}

type ExternalConnectionsObservation struct {

	// The name of the external connection associated with a repository.
	ExternalConnectionName *string `json:"externalConnectionName,omitempty" tf:"external_connection_name,omitempty"`

	PackageFormat *string `json:"packageFormat,omitempty" tf:"package_format,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ExternalConnectionsParameters struct {

	// The name of the external connection associated with a repository.
	// +kubebuilder:validation:Optional
	ExternalConnectionName *string `json:"externalConnectionName" tf:"external_connection_name,omitempty"`
}

type RepositoryInitParameters struct {

	// The description of the repository.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain that contains the created repository.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/codeartifact/v1beta1.Domain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("domain",false)
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The account number of the AWS account that owns the domain.
	DomainOwner *string `json:"domainOwner,omitempty" tf:"domain_owner,omitempty"`

	// Reference to a Domain in codeartifact to populate domain.
	// +kubebuilder:validation:Optional
	DomainRef *v1.Reference `json:"domainRef,omitempty" tf:"-"`

	// Selector for a Domain in codeartifact to populate domain.
	// +kubebuilder:validation:Optional
	DomainSelector *v1.Selector `json:"domainSelector,omitempty" tf:"-"`

	// An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
	ExternalConnections []ExternalConnectionsInitParameters `json:"externalConnections,omitempty" tf:"external_connections,omitempty"`

	// The name of the repository to create.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
	Upstream []UpstreamInitParameters `json:"upstream,omitempty" tf:"upstream,omitempty"`
}

type RepositoryObservation struct {

	// The account number of the AWS account that manages the repository.
	AdministratorAccount *string `json:"administratorAccount,omitempty" tf:"administrator_account,omitempty"`

	// The ARN of the repository.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The description of the repository.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain that contains the created repository.
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The account number of the AWS account that owns the domain.
	DomainOwner *string `json:"domainOwner,omitempty" tf:"domain_owner,omitempty"`

	// An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
	ExternalConnections []ExternalConnectionsObservation `json:"externalConnections,omitempty" tf:"external_connections,omitempty"`

	// The ARN of the repository.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the repository to create.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
	Upstream []UpstreamObservation `json:"upstream,omitempty" tf:"upstream,omitempty"`
}

type RepositoryParameters struct {

	// The description of the repository.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The domain that contains the created repository.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/codeartifact/v1beta1.Domain
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("domain",false)
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// The account number of the AWS account that owns the domain.
	// +kubebuilder:validation:Optional
	DomainOwner *string `json:"domainOwner,omitempty" tf:"domain_owner,omitempty"`

	// Reference to a Domain in codeartifact to populate domain.
	// +kubebuilder:validation:Optional
	DomainRef *v1.Reference `json:"domainRef,omitempty" tf:"-"`

	// Selector for a Domain in codeartifact to populate domain.
	// +kubebuilder:validation:Optional
	DomainSelector *v1.Selector `json:"domainSelector,omitempty" tf:"-"`

	// An array of external connections associated with the repository. Only one external connection can be set per repository. see External Connections.
	// +kubebuilder:validation:Optional
	ExternalConnections []ExternalConnectionsParameters `json:"externalConnections,omitempty" tf:"external_connections,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The name of the repository to create.
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A list of upstream repositories to associate with the repository. The order of the upstream repositories in the list determines their priority order when AWS CodeArtifact looks for a requested package version. see Upstream
	// +kubebuilder:validation:Optional
	Upstream []UpstreamParameters `json:"upstream,omitempty" tf:"upstream,omitempty"`
}

type UpstreamInitParameters struct {

	// The name of an upstream repository.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/codeartifact/v1beta1.Repository
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("repository",false)
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Reference to a Repository in codeartifact to populate repositoryName.
	// +kubebuilder:validation:Optional
	RepositoryNameRef *v1.Reference `json:"repositoryNameRef,omitempty" tf:"-"`

	// Selector for a Repository in codeartifact to populate repositoryName.
	// +kubebuilder:validation:Optional
	RepositoryNameSelector *v1.Selector `json:"repositoryNameSelector,omitempty" tf:"-"`
}

type UpstreamObservation struct {

	// The name of an upstream repository.
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`
}

type UpstreamParameters struct {

	// The name of an upstream repository.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/codeartifact/v1beta1.Repository
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("repository",false)
	// +kubebuilder:validation:Optional
	RepositoryName *string `json:"repositoryName,omitempty" tf:"repository_name,omitempty"`

	// Reference to a Repository in codeartifact to populate repositoryName.
	// +kubebuilder:validation:Optional
	RepositoryNameRef *v1.Reference `json:"repositoryNameRef,omitempty" tf:"-"`

	// Selector for a Repository in codeartifact to populate repositoryName.
	// +kubebuilder:validation:Optional
	RepositoryNameSelector *v1.Selector `json:"repositoryNameSelector,omitempty" tf:"-"`
}

// RepositorySpec defines the desired state of Repository
type RepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RepositoryParameters `json:"forProvider"`
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
	InitProvider RepositoryInitParameters `json:"initProvider,omitempty"`
}

// RepositoryStatus defines the observed state of Repository.
type RepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Repository is the Schema for the Repositorys API. Provides a CodeArtifact Repository resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.repository) || (has(self.initProvider) && has(self.initProvider.repository))",message="spec.forProvider.repository is a required parameter"
	Spec   RepositorySpec   `json:"spec"`
	Status RepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RepositoryList contains a list of Repositorys
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Repository `json:"items"`
}

// Repository type metadata.
var (
	Repository_Kind             = "Repository"
	Repository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Repository_Kind}.String()
	Repository_KindAPIVersion   = Repository_Kind + "." + CRDGroupVersion.String()
	Repository_GroupVersionKind = CRDGroupVersion.WithKind(Repository_Kind)
)

func init() {
	SchemeBuilder.Register(&Repository{}, &RepositoryList{})
}