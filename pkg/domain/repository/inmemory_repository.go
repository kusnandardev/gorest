package repository

type Inmemory interface {
	Get(string) (interface{}, error)
	Set(string, interface{}) error
	Delete(string) error
}
