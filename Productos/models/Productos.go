package Productos

 
type Producto struct{
	ID int
	Nombre string
	Cantidad int
}


var ProductosDB []Producto
var nextID = 1

func GetAllProductos() []Producto {
	return ProductosDB
}

func GetProductoBYID(id int) (Producto, bool){
	for _, p := range ProductosDB {
		if p.ID == id {
			return p, true
		}
	}
	return Producto{}, false
}

func CreateProducto(producto Producto) Producto {
	producto.ID = nextID
	nextID++
	ProductosDB = append(ProductosDB, producto)
	return producto
}


func UpdateProducto(id int , update Producto)(Producto , bool){
	for i, p := range ProductosDB {
		if p.ID == id {
			ProductosDB[i] = update
			ProductosDB[i].ID = id
			return ProductosDB[i], true
		}
	}
	return Producto{}, false
}


func DeleteProducto(id int) bool {
	for i, p := range ProductosDB {
		if p.ID == id {
			ProductosDB = append(ProductosDB[:i], ProductosDB[i+1:]...)
			return true
		}
	}
	return false
}





