package main

import (
	"exam/api-gateway/api-testing/handlers"
	"exam/api-gateway/api-testing/storage/kv"
	"log"
	"net/http"

	// "github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
)

func main() {
	// client := redis.NewClient(&redis.Options{
	// 	Addr: "localhost:6379",
	// })
	// kv.Init(kv.NewRedisClient(client))
	kv.Init(kv.NewInMemoryInst())

	router := gin.New()

	router.POST("/users/register", handlers.RegisterUser)
	router.GET("/users/verify/:code", handlers.Verify)
	router.GET("/users/get", handlers.GetUser)
	router.POST("/users/create", handlers.CreateUser)
	router.DELETE("/users/delete", handlers.DeleteUser)
	router.GET("/users", handlers.ListUsers)

	router.GET("/products/get", handlers.GetProduct)
	router.POST("/products/create", handlers.CreateProduct)
	router.DELETE("/products/delete", handlers.DeleteProduct)
	router.GET("/products", handlers.ListProducts)

	log.Fatal(http.ListenAndServe(":9191", router))
}
