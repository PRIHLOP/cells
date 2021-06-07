package registry

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/selector"
)

type registrySelector struct {
	so selector.Options
}

func NewSelector(opts ...selector.Option) selector.Selector {
	r := &registrySelector{}

	sopts := selector.Options{}
	for _, opt := range opts {
		opt(&sopts)
	}

	r.so = sopts

	return r
}

func (r *registrySelector) Init(opts ...selector.Option) error {
	for _, o := range opts {
		o(&r.so)
	}
	return nil
}
func (r *registrySelector) Options() selector.Options {
	return r.so

}

// Select returns a function which should return the next node
func (r *registrySelector) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	return func() (*registry.Node, error) {
		return &registry.Node{
			Id:      "pydio.grpc.registry",
			Address: "127.0.0.1",
			Port:    8000,
		}, nil
	}, nil
}

// Mark sets the success/error against a node
func (r *registrySelector) Mark(service string, node *registry.Node, err error) {
}

// Reset returns state back to zero for a service
func (r *registrySelector) Reset(service string) {
}

// Close renders the selector unusable
func (r *registrySelector) Close() error {
	return nil
}

// Name of the selector
func (r *registrySelector) String() string {
	return "registry"
}