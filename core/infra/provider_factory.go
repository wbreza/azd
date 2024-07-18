package infra

import (
	"context"

	"github.com/wbreza/azd/core/ioc"
	extinfra "github.com/wbreza/azd/ext/infra"
)

type ProviderFactory struct {
	serviceLocator ioc.ServiceLocator
}

func NewProviderFactory(serviceLocator ioc.ServiceLocator) *ProviderFactory {
	return &ProviderFactory{
		serviceLocator: serviceLocator,
	}
}

func (f *ProviderFactory) Create(ctx context.Context, providerName string) (extinfra.Provider, error) {
	var provider extinfra.Provider

	if err := f.serviceLocator.ResolveNamed(ctx, providerName, &provider); err != nil {
		return nil, err
	}

	return provider, nil
}
