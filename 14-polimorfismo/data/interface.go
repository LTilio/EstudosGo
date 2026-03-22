package data

type DataSource interface {
	Save(name string)
	GetAll() []string
}
