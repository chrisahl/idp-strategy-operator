// Copyright Contributors to the Open Cluster Management project

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/identitatem/idp-strategy-operator/api/client/clientset/versioned/scheme"
	v1alpha1 "github.com/identitatem/idp-strategy-operator/api/identitatem/v1alpha1"
	rest "k8s.io/client-go/rest"
)

type IdentitatemV1alpha1Interface interface {
	RESTClient() rest.Interface
	StrategiesGetter
}

// IdentitatemV1alpha1Client is used to interact with features provided by the identitatem.io group.
type IdentitatemV1alpha1Client struct {
	restClient rest.Interface
}

func (c *IdentitatemV1alpha1Client) Strategies(namespace string) StrategyInterface {
	return newStrategies(c, namespace)
}

// NewForConfig creates a new IdentitatemV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*IdentitatemV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &IdentitatemV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new IdentitatemV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *IdentitatemV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new IdentitatemV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *IdentitatemV1alpha1Client {
	return &IdentitatemV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *IdentitatemV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
