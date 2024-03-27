package services

import (
	"exam/api-gateway/config"
	pbp "exam/api-gateway/genproto/product-service"

	pbu "exam/api-gateway/genproto/user-service"
	"exam/api-gateway/mock"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	UserMockService() mock.UserServiceClient
	ProductService() pbp.ProductServiceClient
}

type serviceManager struct {
	userService     pbu.UserServiceClient
	userMockService mock.UserServiceClient
	productService  pbp.ProductServiceClient
}

func (s *serviceManager) UserService() pbu.UserServiceClient {
	return s.userService
}

func (s *serviceManager) UserMockService() mock.UserServiceClient {
	return s.userMockService
}

func (s *serviceManager) ProductService() pbp.ProductServiceClient {
	return s.productService
}

func NewServiceManager(cfg *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	//User service
	// connUser, err := grpc.Dial(
	// 	fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
	// 	grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	return nil, err
	// }

	//Product service
	connProduct, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.ProductServiceHost, cfg.ProductServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		userMockService: mock.NewUserServiceClient(),
		// userService: pbu.NewUserServiceClient(connUser),
		productService: pbp.NewProductServiceClient(connProduct),
	}
	return serviceManager, nil
}
