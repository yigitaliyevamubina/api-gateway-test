package microservicetesting

import (
	"bytes"
	"encoding/json"
	"exam/api-gateway/config"
	pbp "exam/api-gateway/genproto/product-service"
	pbu "exam/api-gateway/genproto/user-service"
	"exam/api-gateway/microservice-testing/handlers"
	"exam/api-gateway/microservice-testing/models.go"
	"exam/api-gateway/pkg/logger"
	"exam/api-gateway/services"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func RunApiTest(t *testing.T) {
	// New IServiceManager
	cfg := config.Load()
	service, err := services.NewServiceManager(&cfg)
	require.NoError(t, err)
	sm := handlers.NewMockServiceManager(service, logger.New("", ""))

	//tests

	//Create User
	user := &pbu.User{
		Id:           uuid.NewString(),
		FirstName:    "Mubina",
		LastName:     "Yigitaliyeva",
		Age:          17,
		Email:        "mubinayigitaliyeva00@gmail.com",
		Password:     "mubina",
		RefreshToken: "refresh_token",
	}
	payloadBytes, err := json.Marshal(user)
	require.NoError(t, err)
	//response
	r := gin.Default()
	r.POST("/users/create", sm.CreateUser)
	req, err := http.NewRequest(http.MethodPost, "/users/create", bytes.NewReader(payloadBytes))
	require.NoError(t, err)
	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	var respUser *pbu.User
	assert.NoError(t, json.Unmarshal(res.Body.Bytes(), &respUser))
	assert.Equal(t, user.Id, respUser.Id)
	assert.Equal(t, user.FirstName, respUser.FirstName)
	assert.Equal(t, user.LastName, respUser.LastName)
	assert.Equal(t, user.Age, respUser.Age)
	assert.Equal(t, user.Email, respUser.Email)

	// Get User by id
	getReq, err := http.NewRequest(http.MethodGet, "/users/get", nil)
	require.NoError(t, err)
	q := getReq.URL.Query()
	q.Add("id", user.Id)
	getReq.URL.RawQuery = q.Encode()
	getRes := httptest.NewRecorder()
	//response
	r = gin.Default()
	r.GET("/users/get", sm.GetUser)
	r.ServeHTTP(getRes, getReq)
	require.Equal(t, http.StatusOK, getRes.Code)
	var getUserResp *pbu.User
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getUserResp))
	assert.Equal(t, user.Id, getUserResp.Id)
	assert.Equal(t, user.FirstName, getUserResp.FirstName)
	assert.Equal(t, user.LastName, getUserResp.LastName)
	assert.Equal(t, user.Age, getUserResp.Age)
	assert.Equal(t, user.Email, getUserResp.Email)

	user.FirstName = "Updated first name"
	user.LastName = "Updated last name"
	user.Age = 20
	payloadBytes, err = json.Marshal(user)
	require.NoError(t, err)

	// Update User
	updateReq, err := http.NewRequest(http.MethodPut, "/users/update", bytes.NewBuffer(payloadBytes))
	require.NoError(t, err)
	updateRes := httptest.NewRecorder()
	//response
	r = gin.Default()
	r.PUT("/users/update", sm.UpdateUser)
	r.ServeHTTP(updateRes, updateReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, updateRes.Code)
	var updatedUser *pbu.User
	assert.NoError(t, json.Unmarshal(updateRes.Body.Bytes(), &updatedUser))
	assert.Equal(t, user.FirstName, updatedUser.FirstName)
	assert.Equal(t, user.LastName, updatedUser.LastName)
	assert.Equal(t, user.Age, updatedUser.Age)

	// Delete User by id
	delReq, err := http.NewRequest(http.MethodDelete, "/users/delete", nil)
	require.NoError(t, err)
	q = delReq.URL.Query()
	q.Add("id", user.Id)
	delReq.URL.RawQuery = q.Encode()
	delRes := httptest.NewRecorder()
	//response
	r = gin.Default()
	r.DELETE("/users/delete", sm.DeleteUser)
	r.ServeHTTP(delRes, delReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var message models.Message
	bodyBytes, err := io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &message))
	require.Equal(t, "user was deleted", message.Message)

	// Product
	product := &pbp.Product{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       133.34,
		Amount:      10,
	}
	// Create product
	productBytes, err := json.Marshal(product)
	require.NoError(t, err)
	//response
	r = gin.Default()
	r.POST("/products/create", sm.CreateProduct)
	createProdReq, err := http.NewRequest(http.MethodPost, "/products/create", bytes.NewReader(productBytes))
	require.NoError(t, err)
	createProdRes := httptest.NewRecorder()
	r.ServeHTTP(createProdRes, createProdReq)
	assert.Equal(t, http.StatusOK, createProdRes.Code)
	var respProduct *pbp.Product
	assert.NoError(t, json.Unmarshal(createProdRes.Body.Bytes(), &respProduct))
	assert.Equal(t, product.Name, respProduct.Name)
	assert.Equal(t, product.Description, respProduct.Description)
	assert.Equal(t, product.Price, respProduct.Price)
	assert.Equal(t, product.Amount, respProduct.Amount)

	// Get Product by id
	getReq, err = http.NewRequest(http.MethodGet, "/products/get", nil)
	require.NoError(t, err)
	q = getReq.URL.Query()
	q.Add("id", cast.ToString(respProduct.Id))
	getReq.URL.RawQuery = q.Encode()
	getRes = httptest.NewRecorder()
	//response
	r = gin.Default()
	r.GET("/products/get", sm.GetProduct)
	r.ServeHTTP(getRes, getReq)
	require.Equal(t, http.StatusOK, getRes.Code)
	var getProdResp *pbp.Product
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getProdResp))
	assert.Equal(t, product.Name, getProdResp.Name)
	assert.Equal(t, product.Description, getProdResp.Description)
	assert.Equal(t, product.Price, getProdResp.Price)
	assert.Equal(t, product.Amount, getProdResp.Amount)

	product.Id = respProduct.Id
	product.Name = "Updated product name"
	product.Description = "Updated product description"
	product.Amount = 5
	payloadBytes, err = json.Marshal(product)
	require.NoError(t, err)

	// Update Product
	updateReq, err = http.NewRequest(http.MethodPut, "/products/update", bytes.NewBuffer(payloadBytes))
	require.NoError(t, err)
	updateRes = httptest.NewRecorder()
	//response
	r = gin.Default()
	r.PUT("/products/update", sm.UpdateProduct)
	r.ServeHTTP(updateRes, updateReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, updateRes.Code)
	var updatedProduct *pbp.Product
	assert.NoError(t, json.Unmarshal(updateRes.Body.Bytes(), &updatedProduct))
	assert.Equal(t, product.Name, updatedProduct.Name)
	assert.Equal(t, product.Description, updatedProduct.Description)
	assert.Equal(t, product.Amount, updatedProduct.Amount)

	// Delete Product by id
	delReq, err = http.NewRequest(http.MethodDelete, "/products/delete", nil)
	require.NoError(t, err)
	q = delReq.URL.Query()
	q.Add("id", cast.ToString(respProduct.Id))
	delReq.URL.RawQuery = q.Encode()
	delRes = httptest.NewRecorder()
	//response
	r = gin.Default()
	r.DELETE("/products/delete", sm.DeleteProduct)
	r.ServeHTTP(delRes, delReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var delMessage models.Message
	bodyBytes, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &delMessage))
	require.Equal(t, "product was deleted", delMessage.Message)
}
