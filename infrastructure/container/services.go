package container

import (
	"context"
	"errors"
	"go-microservice/domain/service"
)

var _ Services = &Container{}

// Services defines the methods that our Container must implement
type Services interface {
	GetPublisherReferralService(ctx context.Context) (service.PublisherReferralCreator, error)
	GetNetworkAdminAuthService(ctx context.Context) (service.NetworkAdminAuthenticator, error)
	// Add more methods for other services, commands, and queries
}

func (c *Container) GetPublisherReferralService(ctx context.Context) (service.PublisherReferralCreator, error) {
	publisherRepository, err := c.GetPublisherRepository(ctx) // Assume this method exists in Container
	if err != nil {
		return nil, errors.New("error on getting publisher repository: " + err.Error())
	}
	return service.NewPublisherReferralCreator(ctx, publisherRepository), nil
}

func (c *Container) GetNetworkAdminAuthService(ctx context.Context) (service.NetworkAdminAuthenticator, error) {
	adminRepository, err := c.GetAdminRepository(ctx) // Assume this method exists in Container
	if err != nil {
		return nil, errors.New("error on getting admin repository: " + err.Error())
	}
	return service.NewNetworkAdminAuthenticator(ctx, adminRepository), nil
}
