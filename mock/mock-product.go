package mock

import (
	"context"
	pbp "exam/api-gateway/genproto/product-service"

	"google.golang.org/grpc"
)

type ProductServiceClientI interface {
	CreateProduct(ctx context.Context, in *pbp.Product, opts ...grpc.CallOption) (*pbp.Product, error)
	GetProductById(ctx context.Context, in *pbp.GetProductId, opts ...grpc.CallOption) (*pbp.Product, error)
	UpdateProduct(ctx context.Context, in *pbp.Product, opts ...grpc.CallOption) (*pbp.Product, error)
	DeleteProduct(ctx context.Context, in *pbp.GetProductId, opts ...grpc.CallOption) (*pbp.Status, error)
	ListProducts(ctx context.Context, in *pbp.GetListRequest, opts ...grpc.CallOption) (*pbp.GetListResponse, error)
}

type ProductServiceClient struct {
}

func NewProductServiceClient() ProductServiceClientI {
	return &ProductServiceClient{}
}

func (c *ProductServiceClient) CreateProduct(ctx context.Context, in *pbp.Product, opts ...grpc.CallOption) (*pbp.Product, error) {
	return in, nil
}

func (c *ProductServiceClient) GetProductById(ctx context.Context, in *pbp.GetProductId, opts ...grpc.CallOption) (*pbp.Product, error) {

	return &pbp.Product{
		Id:          1,
		Name:        "Test Product name",
		Description: "Product description",
		Price:       10,
		Amount:      13,
	}, nil
}

func (c *ProductServiceClient) UpdateProduct(ctx context.Context, in *pbp.Product, opts ...grpc.CallOption) (*pbp.Product, error) {
	return &pbp.Product{
		Id:          1,
		Name:        "Test Product name",
		Description: "Product description",
		Price:       10,
		Amount:      13,
	}, nil
}

func (c *ProductServiceClient) DeleteProduct(ctx context.Context, in *pbp.GetProductId, opts ...grpc.CallOption) (*pbp.Status, error) {
	return &pbp.Status{
		Success: true,
	}, nil
}

func (c *ProductServiceClient) ListProducts(ctx context.Context, in *pbp.GetListRequest, opts ...grpc.CallOption) (*pbp.GetListResponse, error) {
	pr := pbp.Product{
		Id:          1,
		Name:        "Test Product name",
		Description: "Product description",
		Price:       10,
		Amount:      13,
	}
	return &pbp.GetListResponse{
		Count: 3,
		Products: []*pbp.Product{
			&pr,
			&pr,
			&pr,
		},
	}, nil
}
