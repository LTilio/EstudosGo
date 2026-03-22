package data

import "fmt"

var namesPgDb []string

type Postgress struct {
	host     string
	port     int
	database string
}

func NewPostgress(host, database string, port int) Postgress {
	return Postgress{
		host:     host,
		database: database,
		port:     port,
	}
}

func (p *Postgress) Save(name string) {
	namesPgDb = append(namesPgDb, name)
	fmt.Println("Dado persistido no db: ", name)
}

func (p Postgress) GetAll() []string {
	return namesPgDb
}
