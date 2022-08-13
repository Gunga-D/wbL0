package storage

type Memory interface {
	Get(key string) (interface{}, bool)
	Insert(key string, value interface{})
	Recover(interface{}) error
}
