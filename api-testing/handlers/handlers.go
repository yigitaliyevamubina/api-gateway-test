package handlers

import (
	"encoding/json"
	"exam/api-gateway/api-testing/storage"
	"exam/api-gateway/api-testing/storage/kv"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

// User crud
func RegisterUser(c *gin.Context) {
	var newUser storage.UserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.Id = uuid.NewString()
	newUser.Email = strings.ToLower(newUser.Email)
	err := newUser.Validate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userJson, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(newUser.Id, string(userJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// from := "mubinayigitaliyeva00@gmail.com"
	// password := "iocd vnhb lnvx digm"
	// err = email.SendVerificationCode(email.Params{
	// 	From:     from,
	// 	Password: password,
	// 	To:       newUser.Email,
	// 	Message:  fmt.Sprintf("Hi %s,", newUser.FirstName),
	// 	Code:     "12345",
	// })
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"error": err.Error(),
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "a verification code was sent to your email, please check it.",
	})
}

func Verify(c *gin.Context) {
	userCode := c.Param("code")

	if userCode != "12345" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect code",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Correct code",
	})
}

func CreateUser(c *gin.Context) {
	var newUser storage.UserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newUser.Id = uuid.NewString()

	userJson, err := json.Marshal(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := kv.Set(newUser.Id, string(userJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func GetUser(c *gin.Context) {
	userID := c.Query("id")
	userString, err := kv.Get(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp storage.UserRequest
	if err := json.Unmarshal([]byte(userString), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// pp.Println(resp, "get user")
	c.JSON(http.StatusOK, resp)
}

func DeleteUser(c *gin.Context) {
	userId := c.Query("id")
	if err := kv.Delete(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user was deleted successfully",
	})
}

func ListUsers(c *gin.Context) {
	usersStrings, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var users []*storage.UserRequest
	for _, userString := range usersStrings {
		var user storage.UserRequest
		if err := json.Unmarshal([]byte(userString), &user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		users = append(users, &user)
	}

	c.JSON(http.StatusOK, users)
}

// Product crud
func CreateProduct(c *gin.Context) {
	var newProduct storage.Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		// fmt.Println(err, "111")
		return
	}

	productJson, err := json.Marshal(newProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		// fmt.Println(err, "222")
		return
	}

	if err := kv.Set(cast.ToString(newProduct.Id), string(productJson), 1000); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		// fmt.Println(err, "333")
		return
	}

	c.JSON(http.StatusOK, newProduct)
}

func GetProduct(c *gin.Context) {
	productID := c.Query("id")
	productString, err := kv.Get(productID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var resp storage.Product
	if err := json.Unmarshal([]byte(productString), &resp); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteProduct(c *gin.Context) {
	productId := c.Query("id")
	if err := kv.Delete(productId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product was deleted successfully",
	})
}

func ListProducts(c *gin.Context) {
	productsStrings, err := kv.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var products []*storage.Product
	for _, productString := range productsStrings {
		var product storage.Product
		if err := json.Unmarshal([]byte(productString), &product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		products = append(products, &product)
	}

	c.JSON(http.StatusOK, products)
}
