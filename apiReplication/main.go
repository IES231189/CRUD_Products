package main

import (
	"github.com/gin-gonic/gin"
	"replica/apiReplication/controllers"
)

func main() {
	r := gin.Default()

	productos := r.Group("/api/productos")
	{
		productos.GET("/", controllers.GetProductos)
		productos.POST("/", controllers.CreateProducto)
		productos.PUT("/:id", controllers.UpdateProducto)
		productos.DELETE("/:id", controllers.DeleteProducto)
	}

	r.Run(":8081")
}
