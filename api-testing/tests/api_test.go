package tests

import (
	"encoding/json"
	"exam/api-gateway/api-testing/handlers"
	"exam/api-gateway/api-testing/storage"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApi(t *testing.T) {
	require.NoError(t, SetupMinimumInstance(""))
	buffer, err := OpenFile("user.json")
	// fmt.Println(err, " +++++  user json")
	require.NoError(t, err)

	// User Create
	req := NewRequest(http.MethodPost, "/users/create", buffer)
	res := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/users/create", handlers.CreateUser)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

	var user storage.UserRequest
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &user))
	fmt.Println(user ,"-----------")
	require.Equal(t, user.Email, "mubinayigitaliyeva00@gmail.com")
	require.Equal(t, user.Age, int64(17))
	require.Equal(t, user.FirstName, "Mubina")
	require.Equal(t, user.LastName, "Yigitaliyeva")
	require.Equal(t, user.Password, "test")
	require.NotNil(t, user.Id)

	// User Get
	getReq := NewRequest(http.MethodGet, "/users/get", buffer)
	q := getReq.URL.Query()
	q.Add("id", user.Id)
	getReq.URL.RawQuery = q.Encode()
	getRes := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users/get", handlers.GetUser)
	r.ServeHTTP(getRes, getReq)
	require.Equal(t, http.StatusOK, getRes.Code)
	var getUserResp storage.UserRequest
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getUserResp))
	assert.Equal(t, user.Id, getUserResp.Id)
	assert.Equal(t, user.FirstName, getUserResp.FirstName)
	assert.Equal(t, user.LastName, getUserResp.LastName)
	assert.Equal(t, user.Age, getUserResp.Age)
	assert.Equal(t, user.Email, getUserResp.Email)

	// User List
	listReq := NewRequest(http.MethodGet, "/users", buffer)
	listRes := httptest.NewRecorder()
	
	r.GET("/users", handlers.ListUsers)
	r.ServeHTTP(listRes, listReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bodyBytes, err := io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)

	// User Delete
	delReq := NewRequest(http.MethodDelete, "/users/delete", buffer)
	q = delReq.URL.Query()
	q.Add("id", user.Id)
	delReq.URL.RawQuery = q.Encode()
	delRes := httptest.NewRecorder()
	r.DELETE("/users/delete", handlers.DeleteUser)
	r.ServeHTTP(delRes, delReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var message storage.Message
	bodyBytes, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &message))
	require.Equal(t, "user was deleted successfully", message.Message)

	// User Register
	regReq := NewRequest(http.MethodPost, "/users/register", buffer)
	regRes := httptest.NewRecorder()
	r.POST("/users/register", handlers.RegisterUser)
	r.ServeHTTP(regRes, regReq)
	// fmt.Println(string(buffer))
	assert.Equal(t, http.StatusOK, regRes.Code)
	var resp storage.Message
	bodyBytes, err = io.ReadAll(regRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &resp))
	require.NotNil(t, resp.Message)
	require.Equal(t, "a verification code was sent to your email, please check it.", resp.Message)

	// User Verify
	uri := fmt.Sprintf("/users/verify/%s", "12345")
	verReq := NewRequest(http.MethodGet, uri, buffer)
	verRes := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users/verify/:code", handlers.Verify)
	r.ServeHTTP(verRes, verReq)
	assert.Equal(t, http.StatusOK, verRes.Code)
	var response *storage.Message
	bodyBytes, err = io.ReadAll(verRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &response))
	require.Equal(t, "Correct code", response.Message)

	//User Verify with incorrect code
	incorrectURI := fmt.Sprintf("/users/verify/%s", "11111")
	incorrectVerReq := NewRequest(http.MethodGet, incorrectURI, buffer)
	incorrectVerRes := httptest.NewRecorder()
	r.ServeHTTP(incorrectVerRes, incorrectVerReq)
	assert.Equal(t, http.StatusBadRequest, incorrectVerRes.Code)
	var incorrectResponse storage.Message
	bodyBytes, err = io.ReadAll(incorrectVerRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &incorrectResponse))
	require.Equal(t, "Incorrect code", incorrectResponse.Message)

	gin.SetMode(gin.TestMode)
	require.NoError(t, SetupMinimumInstance(""))
	buffer, err = OpenFile("product.json")
	require.NoError(t, err)

	// Product Create
	req = NewRequest(http.MethodPost, "/products/create", buffer)
	// fmt.Println(string(buffer))
	res = httptest.NewRecorder()
	r.POST("products/create", handlers.CreateProduct)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)
	var product storage.Product
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &product))
	require.Equal(t, product.Amount, int32(12))
	require.Equal(t, product.Description, "Test Description")
	require.Equal(t, product.Name, "Test Product")
	require.Equal(t, product.Price, float32(127.9))

	// Product Get
	getReq = NewRequest(http.MethodGet, "/products/get", buffer)
	q = getReq.URL.Query()
	q.Add("id", string(product.Id))
	getReq.URL.RawQuery = q.Encode()
	getRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/products/get", handlers.GetProduct)
	r.ServeHTTP(getRes, getReq)
	assert.Equal(t, http.StatusOK, getRes.Code)

	var getProduct storage.Product
	bodyBytes, err = io.ReadAll(getRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &getProduct))
	// pp.Println(string(bodyBytes))
	require.Equal(t, product.Amount, getProduct.Amount)
	require.Equal(t, product.Description, getProduct.Description)
	require.Equal(t, product.Name, getProduct.Name)

	// Product List
	listReq = NewRequest(http.MethodGet, "/products", buffer)
	listRes = httptest.NewRecorder()
	r = gin.Default()
	r.GET("/products", handlers.ListProducts)
	r.ServeHTTP(listRes, listReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bodyBytes, err = io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)
	// pp.Println(string(bodyBytes))

	// Product Delete
	delReq = NewRequest(http.MethodDelete, "/products/delete", buffer)
	q = delReq.URL.Query()
	q.Add("id", string(product.Id))
	delReq.URL.RawQuery = q.Encode()
	r.DELETE("products/delete", handlers.DeleteProduct)
	r.ServeHTTP(delRes, delReq)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var productMessage storage.Message
	bodyBytes, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &productMessage))
	require.Equal(t, "product was deleted successfully", productMessage.Message)
}
