package controllers

import (
	"net/http"
	"strconv"

	"productos/Productos/models"

	"github.com/gin-gonic/gin"
)

func GetProductos(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetAllProductos())
}

func GetProductoByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if producto, found := models.GetProductoByID(id); found {
		c.JSON(http.StatusOK, producto)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
	}
}

func CreateProducto(c *gin.Context) {
	var producto models.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	created := models.CreateProducto(producto)
	c.JSON(http.StatusCreated, created)
}

func UpdateProducto(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var producto models.Producto
	if err := c.ShouldBindJSON(&producto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if updated, found := models.UpdateProducto(id, producto); found {
		c.JSON(http.StatusOK, updated)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
	}
}

func DeleteProducto(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if models.DeleteProducto(id) {
		c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Producto no encontrado"})
	}
}
