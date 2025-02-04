package controllers


import (
	"net/http"
	"github.com/gin-gonic/gin"
	"replica/apiReplication/models"
)



func CreateProducto(c *gin.Context) {
	var producto models.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	created := models.CreateProducto(producto)
	c.JSON(http.StatusCreated, created)
}
