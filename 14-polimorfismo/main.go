package main

import (
	"fmt"

	"github.com/ltilio/estudosGo/14-polimorfismo/data"
)

func main() {

	// db := data.NewMongodb("localhost", "teste-polimorfismo")
	db := data.NewPostgress("localhost", "teste-polimorfismo", 5432)
	saveNewName(&db, "Leandro")
	saveNewName(&db, "Gabriela14")
	getNames(&db)
}

func saveNewName(d data.DataSource, name string) {
	d.Save(name)
}

func getNames(d data.DataSource) {
	fmt.Println(d.GetAll())
}
