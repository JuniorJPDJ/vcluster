// Code generated by generator. DO NOT EDIT.

package v1

import (
	"github.com/loft-sh/api/v3/pkg/apis/virtualcluster"
	"github.com/loft-sh/apiserver/pkg/builders"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func addKnownTypes(scheme *runtime.Scheme) error {
	// TODO this will get cleaned up with the scheme types are fixed
	scheme.AddKnownTypes(SchemeGroupVersion,
		&HelmRelease{},
		&HelmReleaseList{},
	)
	return nil
}

var (
	ApiVersion = builders.NewApiVersion("virtualcluster.loft.sh", "v1").WithResources(
		virtualcluster.VirtualclusterHelmReleaseStorage,
	)

	// Required by code generated by go2idl
	AddToScheme = (&runtime.SchemeBuilder{
		ApiVersion.SchemeBuilder.AddToScheme,
		RegisterDefaults,
		RegisterConversions,
		addKnownTypes,
		func(scheme *runtime.Scheme) error {
			metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
			return nil
		},
	}).AddToScheme

	SchemeBuilder      = ApiVersion.SchemeBuilder
	localSchemeBuilder = SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type HelmReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HelmRelease `json:"items"`
}