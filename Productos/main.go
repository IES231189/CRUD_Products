package main

import(
	"github.com/gin-gonic/gin"
	"productos/Productos/controllers"
)

func Main(){
	r := gin.Default()

	productos := r.Group("api/productos"){
		productos.GET("/" , controllers.GetProductos)
		productos.GET("/" , controllers.GetProductoBYID)
		productos.POST("/" , controllers.CreateProducto)
		productos.PUT("/:id" , controllers.UpdateProducto)
		productos.DELETE("/:id" , controllers.DeleteProducto)

	}


}