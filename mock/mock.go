package mock

import (
	"context"
	pbu "exam/api-gateway/genproto/user-service"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type UserServiceClientI interface {
	CreateUser(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error)
	GetUserById(ctx context.Context, in *pbu.GetUserId, opts ...grpc.CallOption) (*pbu.User, error)
	UpdateUser(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error)
	DeleteUser(ctx context.Context, in *pbu.GetUserId, opts ...grpc.CallOption) (*pbu.Status, error)
	ListUsers(ctx context.Context, in *pbu.GetListRequest, opts ...grpc.CallOption) (*pbu.GetListResponse, error)
	CheckField(ctx context.Context, in *pbu.CheckFieldRequest, opts ...grpc.CallOption) (*pbu.CheckFieldResponse, error)
	Check(ctx context.Context, in *pbu.IfExists, opts ...grpc.CallOption) (*pbu.User, error)
	UpdateRefreshToken(ctx context.Context, in *pbu.UpdateRefreshTokenReq, opts ...grpc.CallOption) (*pbu.Status, error)
}

type UserServiceClient struct {
}

func NewUserServiceClient() UserServiceClientI {
	return &UserServiceClient{}
}

func (c *UserServiceClient) CreateUser(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error) {
	return in, nil
}

func (c *UserServiceClient) GetUserById(ctx context.Context, in *pbu.GetUserId, opts ...grpc.CallOption) (*pbu.User, error) {
	return &pbu.User{
		Id:           "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf",
		FirstName:    "Test FirstName",
		LastName:     "Test Lastname",
		Age:          18,
		Email:        "testemail@gmail.com",
		Password:     "**kkw##knrtest",
		RefreshToken: "refresh token test",
		CreatedAt:    time.Now().String(),
	}, nil
}

func (c *UserServiceClient) UpdateUser(ctx context.Context, in *pbu.User, opts ...grpc.CallOption) (*pbu.User, error) {
	return in, nil
}

func (c *UserServiceClient) DeleteUser(ctx context.Context, in *pbu.GetUserId, opts ...grpc.CallOption) (*pbu.Status, error) {
	return &pbu.Status{
		Success: true,
	}, nil
}

func (c *UserServiceClient) ListUsers(ctx context.Context, in *pbu.GetListRequest, opts ...grpc.CallOption) (*pbu.GetListResponse, error) {
	return &pbu.GetListResponse{
		Count: 3,
		Users: []*pbu.User{
			{
				Id:           "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf",
				FirstName:    "Test FirstName",
				LastName:     "Test Lastname",
				Age:          18,
				Email:        "testemail@gmail.com",
				Password:     "**kkw##knrtest",
				RefreshToken: "refresh token test",
				CreatedAt:    time.Now().String(),
			},
			{
				Id:           uuid.NewString(),
				FirstName:    "Test FirstName 2",
				LastName:     "Test Lastname 2",
				Age:          20,
				Email:        "testemail@gmail.com",
				Password:     "**kkw##knrtest 2",
				RefreshToken: "refresh token test 2",
				CreatedAt:    time.Now().String(),
			},
			{
				Id:           uuid.NewString(),
				FirstName:    "Test FirstName 3",
				LastName:     "Test Lastname 3",
				Age:          20,
				Email:        "testemail@gmail.com",
				Password:     "**kkw##knrtest 3",
				RefreshToken: "refresh token test 3",
				CreatedAt:    time.Now().String(),
			},
		},
		}, nil
}

func (c *UserServiceClient) CheckField(ctx context.Context, in *pbu.CheckFieldRequest, opts ...grpc.CallOption) (*pbu.CheckFieldResponse, error) {
	return &pbu.CheckFieldResponse{Status: true}, nil
}

func (c *UserServiceClient) Check(ctx context.Context, in *pbu.IfExists, opts ...grpc.CallOption) (*pbu.User, error) {
	return &pbu.User{
		Id:           "d4f3f3ce-15f8-48da-9938-e5d9e0bb2aaf",
		FirstName:    "Test FirstName",
		LastName:     "Test Lastname",
		Age:          18,
		Email:        "testemail@gmail.com",
		Password:     "**kkw##knrtest",
		RefreshToken: "refresh token test",
		CreatedAt:    time.Now().String(),
	}, nil
}

func (c *UserServiceClient) UpdateRefreshToken(ctx context.Context, in *pbu.UpdateRefreshTokenReq, opts ...grpc.CallOption) (*pbu.Status, error) {
	return &pbu.Status{Success: true}, nil
}
