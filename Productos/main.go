package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"productos/Productos/controllers"
	"productos/Productos/models"
)

const replicaServerURL = "http://localhost:8081"

func main() {
	r := gin.Default()

	productos := r.Group("/api/productos")
	{
		productos.GET("/", controllers.GetProductos)
		productos.GET("/:id", controllers.GetProductoByID)
		productos.POST("/", func(c *gin.Context) {
			controllers.CreateProducto(c)
			replicateToServer("POST", "/api/productos", c)
		})
		productos.PUT("/:id", func(c *gin.Context) {
			controllers.UpdateProducto(c)
			replicateToServer("PUT", "/api/productos/"+c.Param("id"), c)
		})
		
		productos.DELETE("/:id", func(c *gin.Context) {
			controllers.DeleteProducto(c)
			replicateToServer("DELETE", "/api/productos/"+c.Param("id"), c)
		})
	}

	r.Run(":8080")
}

func replicateToServer(method, path string, c *gin.Context) {
	url := replicaServerURL + path
	var body bytes.Buffer

	if method == http.MethodPost || method == http.MethodPut {
		if err := json.NewEncoder(&body).Encode(c.Request.Body); err != nil {
			fmt.Println("Error encoding JSON:", err)
			return
		}
	}

	req, err := http.NewRequest(method, url, &body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error replicating to server:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Replication status:", resp.Status)
}
