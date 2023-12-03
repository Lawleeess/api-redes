package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

func main() {
	RunServer()
}

// RunServer initialize api server
func RunServer() {
	server := gin.Default()

	server.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	}))

	registerRoutes(server)

	_ = server.Run(":8080")
}

func registerRoutes(e *gin.Engine) {

	authRoutes := e.Group("/api/v1/welcome")
	authRoutes.GET("", Welcome)
	authRoutes.POST("", Hello)
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, "Bienvenido")
}

func Hello(c *gin.Context) {
	var name Name

	if err := c.Bind(&name); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to bind name")
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("Hola %s", name.Name))

}

type Name struct {
	Name string `json:"name"`
}
