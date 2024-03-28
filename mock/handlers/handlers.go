package handlers

import (
	"context"
	"exam/api-gateway/mock"
	"net/http"

	pb "exam/api-gateway/genproto/user-service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	UserService *mock.UserServiceClient
}

func NewHandler(userService *mock.UserServiceClient)  *Handler {
	return &Handler{UserService: userService}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var newUser pb.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.UserService.CreateUser(context.Background(), &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var newUser pb.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.UserService.UpdateUser(context.Background(), &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUser(c *gin.Context) {
	userID := c.Query("id")

	user, err := h.UserService.GetUserById(context.Background(), &pb.GetUserId{UserId: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userId := c.Query("id")

	_, err := h.UserService.DeleteUser(context.Background(), &pb.GetUserId{UserId: userId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user was deleted successfully",
	})
}

func (h *Handler) ListUsers(c *gin.Context) {
	users, err := h.UserService.ListUsers(context.Background(), &pb.GetListRequest{Page: 1, Limit: 10})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) CheckField(c *gin.Context) {
	var body pb.CheckFieldRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	status, err := h.UserService.CheckField(context.Background(), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if status.Status {
		c.JSON(http.StatusOK, gin.H{
			"message": "user exists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user does not exist",
	})
}

func (h *Handler) Check(c *gin.Context) {
	email := c.Query("email")

	user, err := h.UserService.Check(context.Background(), &pb.IfExists{Email: email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateRefreshToken(c *gin.Context) {
	var body pb.UpdateRefreshTokenReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	status, err := h.UserService.UpdateRefreshToken(context.Background(), &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if status.Success {
		c.JSON(http.StatusOK, gin.H{
			"message": "updated",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "not updated",
	})
}
