package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//Producto contiene datos de un articulo
type Producto struct {
	gorm.Model   //id, CreatedAt, UpdatedAt, DeletedAt
	CodigoBarras string
	Precio       uint
}

func main() {
	db, err := gorm.Open("mysql", "root:@/edCurso?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Error en la conexion a la base de datos")
	}
	defer db.Close()
	fmt.Println("Se conectó a la base de datos")
	//--------------------------------------------------------------
	// db.CreateTable(&Producto{})
	// fmt.Println("Se creó la tabla productos en la base de datos")

	//crear u producto en la base de datos

	// db.Create(&Producto{
	// 	CodigoBarras:"85749857484",
	// 	Precio:420000})
	// 	fmt.Println("Se creó un producto en la base de datos")

	//consultar datos de la basae de datos

	var p Producto
	// db.First(&p)
	// fmt.Println("Precio del producto consultado: ", p.Precio)
	db.First(&p, 2) //trae el primer registro q coincida con el id 2
	fmt.Println("Precio del producto consultado: ", p.Precio)

}
