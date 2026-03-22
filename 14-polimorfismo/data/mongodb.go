package data

import "fmt"

var namesMongoDb []string

type Mongodb struct {
	host         string
	organization string
}

func NewMongodb(host, organization string) Mongodb {
	return Mongodb{
		host:         host,
		organization: organization,
	}
}

func (m *Mongodb) Save(name string) {
	namesMongoDb = append(namesMongoDb, name)
	fmt.Println("Dado persistido no MONGO: ", name)
}

func (m Mongodb) GetAll() []string {
	return namesMongoDb
}
