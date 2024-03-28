package tests

import (
	"encoding/json"
	"exam/api-gateway/mock/handlers"
	"exam/api-gateway/api-testing/storage"
	pb "exam/api-gateway/genproto/user-service"
	"exam/api-gateway/mock"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPI(t *testing.T) {
	buffer, err := mock.OpenFile("user.json")
	require.NoError(t, err)

	h := handlers.NewHandler(&mock.UserServiceClient{})

	// User Create
	req := mock.NewRequest(http.MethodPost, "/users/create", buffer)
	res := httptest.NewRecorder()
	r := gin.Default()
	r.POST("/users/create", h.CreateUser)
	r.ServeHTTP(res, req)
	assert.Equal(t, http.StatusOK, res.Code)

	var user pb.User
	require.NoError(t, json.Unmarshal(res.Body.Bytes(), &user))
	fmt.Println(user, "-----------")
	require.Equal(t, user.Email, "testemail@gmail.com")
	require.Equal(t, user.Age, int64(18))
	require.Equal(t, user.FirstName, "Test FirstName")
	require.Equal(t, user.LastName, "Test Lastname")
	require.Equal(t, user.Password, "**kkw##knrtest")
	require.Equal(t, user.RefreshToken, "refresh token test")

	// User Get
	getReq := mock.NewRequest(http.MethodGet, "/users/get", buffer)
	q := getReq.URL.Query()
	q.Add("id", user.Id)
	getReq.URL.RawQuery = q.Encode()
	getRes := httptest.NewRecorder()
	r = gin.Default()
	r.GET("/users/get", h.GetUser)
	r.ServeHTTP(getRes, getReq)
	require.Equal(t, http.StatusOK, getRes.Code)
	var getUserResp pb.User
	require.NoError(t, json.Unmarshal(getRes.Body.Bytes(), &getUserResp))
	assert.Equal(t, user.Id, getUserResp.Id)
	assert.Equal(t, user.FirstName, getUserResp.FirstName)
	assert.Equal(t, user.LastName, getUserResp.LastName)
	assert.Equal(t, user.Age, getUserResp.Age)
	assert.Equal(t, user.Email, getUserResp.Email)

	// User List
	listReq := mock.NewRequest(http.MethodGet, "/users", buffer)
	listRes := httptest.NewRecorder()

	r.GET("/users", h.ListUsers)
	r.ServeHTTP(listRes, listReq)
	assert.Equal(t, http.StatusOK, listRes.Code)
	bodyBytes, err := io.ReadAll(listRes.Body)
	assert.NoError(t, err)
	assert.NotNil(t, bodyBytes)

	// User Delete
	delReq := mock.NewRequest(http.MethodDelete, "/users/delete", buffer)
	q = delReq.URL.Query()
	q.Add("id", user.Id)
	delReq.URL.RawQuery = q.Encode()
	delRes := httptest.NewRecorder()
	r.DELETE("/users/delete", h.DeleteUser)
	r.ServeHTTP(delRes, delReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, delRes.Code)
	var resMessage storage.Message
	bodyBytes, err = io.ReadAll(delRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &resMessage))
	require.Equal(t, "user was deleted successfully", resMessage.Message)

	// Check field
	body := pb.CheckFieldRequest{
		Field: "email",
		Data:  "testemail@gmail.com",
	}
	buffer, err = json.Marshal(body)
	assert.NoError(t, err)
	checkFieldReq := mock.NewRequest(http.MethodPost, "/users/checkfield", buffer)
	checkFieldRes := httptest.NewRecorder()
	r.POST("/users/checkfield", h.CheckField)
	r.ServeHTTP(checkFieldRes, checkFieldReq)
	assert.NoError(t, err)
	fmt.Println(err, "<--------------")
	assert.Equal(t, http.StatusOK, checkFieldRes.Code)
	var message storage.Message
	bodyBytes, err = io.ReadAll(checkFieldRes.Body)
	require.NoError(t, err)
	require.NoError(t, json.Unmarshal(bodyBytes, &message))
	require.Equal(t, "user exists", message.Message)
	fmt.Println(message, "==============")

	//Check
	checkReq := mock.NewRequest(http.MethodDelete, "/users/check", buffer)
	q = delReq.URL.Query()
	q.Add("email", user.Email)
	checkReq.URL.RawQuery = q.Encode()
	checkRes := httptest.NewRecorder()
	r.DELETE("/users/check", h.Check)
	r.ServeHTTP(checkRes, checkReq)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, checkRes.Code)
	var checkUser pb.User
	require.NoError(t, json.Unmarshal(checkRes.Body.Bytes(), &checkUser))
	assert.Equal(t, user.Id, checkUser.Id)
	assert.Equal(t, user.FirstName, checkUser.FirstName)
	assert.Equal(t, user.LastName, checkUser.LastName)
	assert.Equal(t, user.Age, checkUser.Age)
	assert.Equal(t, user.Email, checkUser.Email)

	// Update refresh token
	bodyReq := pb.UpdateRefreshTokenReq{
		UserId:       user.Id,
		RefreshToken: "refresh token",
	}
	buffer, err = json.Marshal(bodyReq)
	assert.NoError(t, err)
	updateRefreshReq := mock.NewRequest(http.MethodPost, "/users/update/refreshtoken", buffer)
	updateRefreshRes := httptest.NewRecorder()
	r.POST("/users/update/refreshtoken", h.UpdateRefreshToken)
	r.ServeHTTP(updateRefreshRes, updateRefreshReq)
	assert.NoError(t, err)
	fmt.Println(err, "-----------")
	assert.Equal(t, http.StatusOK, updateRefreshRes.Code)
	var updateMess storage.Message
	bodyBytes, err = io.ReadAll(updateRefreshRes.Body)
	require.NoError(t, err)
	fmt.Println(err)
	require.NoError(t, json.Unmarshal(bodyBytes, &updateMess))
	require.Equal(t, "updated", updateMess.Message)
}
