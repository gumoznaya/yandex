package repository

type Storage struct{
	ID int
	Long string
	Short string
}
type InMemory struct {
	InMemory []Storage
}

var InMemoryStorage InMemory 
