//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by client-gen/saf1u. DO NOT EDIT.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/saf1u/haus.com/pkg/apis/haus.com/v1"
	scheme "github.com/saf1u/haus.com/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HausesGetter has a method to return a HausInterface.
// A group's client should implement this interface.
type HausesGetter interface {
	Hauses(namespace string) HausInterface
}

// HausInterface has methods to work with Haus resources.
type HausInterface interface {
	Create(ctx context.Context, haus *v1.Haus, opts metav1.CreateOptions) (*v1.Haus, error)
	Update(ctx context.Context, haus *v1.Haus, opts metav1.UpdateOptions) (*v1.Haus, error)
	UpdateStatus(ctx context.Context, haus *v1.Haus, opts metav1.UpdateOptions) (*v1.Haus, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Haus, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.HausList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Haus, err error)
	HausExpansion
}

// hauses implements HausInterface
type hauses struct {
	client rest.Interface
	ns     string
}

// newHauses returns a Hauses
func newHauses(c *SamplecontrollerV1Client, namespace string) *hauses {
	return &hauses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the haus, and returns the corresponding haus object, and an error if there is any.
func (c *hauses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Haus, err error) {
	result = &v1.Haus{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hauses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Hauses that match those selectors.
func (c *hauses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.HausList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.HausList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hauses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hauses.
func (c *hauses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hauses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a haus and creates it.  Returns the server's representation of the haus, and an error, if there is any.
func (c *hauses) Create(ctx context.Context, haus *v1.Haus, opts metav1.CreateOptions) (result *v1.Haus, err error) {
	result = &v1.Haus{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hauses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(haus).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a haus and updates it. Returns the server's representation of the haus, and an error, if there is any.
func (c *hauses) Update(ctx context.Context, haus *v1.Haus, opts metav1.UpdateOptions) (result *v1.Haus, err error) {
	result = &v1.Haus{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hauses").
		Name(haus.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(haus).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *hauses) UpdateStatus(ctx context.Context, haus *v1.Haus, opts metav1.UpdateOptions) (result *v1.Haus, err error) {
	result = &v1.Haus{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hauses").
		Name(haus.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(haus).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the haus and deletes it. Returns an error if one occurs.
func (c *hauses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hauses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hauses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hauses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched haus.
func (c *hauses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Haus, err error) {
	result = &v1.Haus{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hauses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}