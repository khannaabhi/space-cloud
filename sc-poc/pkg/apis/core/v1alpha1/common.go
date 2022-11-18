package v1alpha1

import "k8s.io/apimachinery/pkg/runtime"

// SecretSource represents a source for a confidential value.
// Taken from apiv1.EnvVarSource.
type SecretSource struct {
	// SecretKeyRef selects a key of a secret in the pod's namespace
	// +kubebuilder:validation:Optional
	SecretKeyRef *SecretKeyRef `json:"secretKeyRef,omitempty"`

	// EnvRef selects value from an environment variable
	// +kubebuilder:validation:Optional
	EnvRef *EnvRef `json:"envRef,omitempty"`
}

type SecretKeyRef struct {
	// Name of the secret
	Name string `json:"name"`

	// Key of the secret
	Key string `json:"key"`

	// Source describes which secret manager to use. By default its k8s
	// +kubebuilder:validation:Optional
	Source string `json:"source,omitempty"`
}

// EnvRef describes a reference to an environment variable
type EnvRef struct {
	// Name of the environment variable
	Name string `json:"name"`
}

// ResourceRef describes a reference to a resource object
type ResourceRef struct {
	// Name of the resource
	Name string `json:"name"`
}

// AuthSecret describes the state of common properties required in every auth secret
type AuthSecret struct {
	// IsPrimary denotes if this secret is to be used as the default secret
	IsPrimary bool `json:"isPrimary"`

	// The kid value of this secret
	KID string `json:"kid"`

	// AllowedAudiences is describes the allowed values in the "aud" field of the jwt token.
	// +kubebuilder:validation:Optional
	AllowedAudiences []string `json:"allowedAudiences,omitempty"`

	// AllowedIssuers describes the allowed values in the "iss" field of the jwt token.
	// +kubebuilder:validation:Optional
	AllowedIssuers []string `json:"allowedIssuers,omitempty"`
}

// HTTPPlugin describes a plugin to be used in an HTTP endpoint
type HTTPPlugin struct {
	// Name describes a name of the plugin
	Name string `json:"name"`

	// Driver describes the driver to use for the plugin
	Driver string `json:"driver"`

	// Params describes the additional properties which are required by the driver
	Params runtime.RawExtension `json:"params,omitempty"`
}