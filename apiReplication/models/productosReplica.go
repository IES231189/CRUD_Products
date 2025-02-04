package models

import (
	"errors"
)


type Producto struct {
	ID       int    
	Nombre   string 
	Cantidad int    
}

var productos = []Producto{}
var currentID = 1


func CreateProducto(producto Producto) Producto {
	producto.ID = currentID
	currentID++
	productos = append(productos, producto)
	return producto
}


func GetProductos() []Producto {
	return productos
}


func GetProductoByID(id int) (Producto, error) {
	for _, p := range productos {
		if p.ID == id {
			return p, nil
		}
	}
	return Producto{}, errors.New("Producto no encontrado")
}


func UpdateProducto(id int, updated Producto) (Producto, error) {
	for i, p := range productos {
		if p.ID == id {
			productos[i].Nombre = updated.Nombre
			productos[i].Cantidad = updated.Cantidad
			return productos[i], nil
		}
	}
	return Producto{}, errors.New("Producto no encontrado")
}


func DeleteProducto(id int) error {
	for i, p := range productos {
		if p.ID == id {
			productos = append(productos[:i], productos[i+1:]...)
			return nil
		}
	}
	return errors.New("Producto no encontrado")
}
