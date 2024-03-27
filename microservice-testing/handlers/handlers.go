package handlers

import (
	"context"
	"exam/api-gateway/pkg/logger"
	"exam/api-gateway/services"
	"net/http"
	"strconv"

	pbp "exam/api-gateway/genproto/product-service"
	pbu "exam/api-gateway/genproto/user-service"

	"github.com/gin-gonic/gin"
)

type MockServiceManager struct {
	sm  services.IServiceManager
	log logger.Logger
}

func NewMockServiceManager(sm services.IServiceManager, log logger.Logger) *MockServiceManager {
	return &MockServiceManager{sm: sm, log: log}
}

func (sm *MockServiceManager) CreateUser(c *gin.Context) {
	var user pbu.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("invalid json", logger.Error(err))
		return
	}

	resp, err := sm.sm.UserService().CreateUser(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while creating user", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) GetUser(c *gin.Context) {
	id := c.Query("id")
	resp, err := sm.sm.UserService().GetUserById(context.Background(), &pbu.GetUserId{UserId: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while getting user", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) DeleteUser(c *gin.Context) {
	id := c.Query("id")
	resp, err := sm.sm.UserService().DeleteUser(context.Background(), &pbu.GetUserId{UserId: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while deleting user", logger.Error(err))
		return
	}

	if resp.Success {
		c.JSON(http.StatusOK, gin.H{
			"message": "user was deleted",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error while deleting user, status false",
		})
	}
}

func (sm *MockServiceManager) UpdateUser(c *gin.Context) {
	var user pbu.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("invalid json", logger.Error(err))
		return
	}

	resp, err := sm.sm.UserService().UpdateUser(context.Background(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while updating user", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}


// Product
func (sm *MockServiceManager) CreateProduct(c *gin.Context) {
	var product pbp.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("invalid json", logger.Error(err))
		return
	}

	resp, err := sm.sm.ProductService().CreateProduct(context.Background(), &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while creating product", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) GetProduct(c *gin.Context) {
	id := c.Query("id")
	int32Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while converting id to int32", logger.Error(err))
		return
	}

	resp, err := sm.sm.ProductService().GetProductById(context.Background(), &pbp.GetProductId{ProductId: int32(int32Id)})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while getting product", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (sm *MockServiceManager) DeleteProduct(c *gin.Context) {
	id := c.Query("id")
	int32Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while converting id to int", logger.Error(err))
		return
	}
	resp, err := sm.sm.ProductService().DeleteProduct(context.Background(), &pbp.GetProductId{ProductId: int32(int32Id)})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while deleting product", logger.Error(err))
		return
	}

	if resp.Success {
		c.JSON(http.StatusOK, gin.H{
			"message": "product was deleted",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error while deleting product, status false",
		})
	}
}

func (sm *MockServiceManager) UpdateProduct(c *gin.Context) {
	var product pbp.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("invalid json", logger.Error(err))
		return
	}

	resp, err := sm.sm.ProductService().UpdateProduct(context.Background(), &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		sm.log.Error("error while updating product", logger.Error(err))
		return
	}

	c.JSON(http.StatusOK, resp)
}
