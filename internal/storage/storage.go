package storage

type Memory interface {
	Get(key string) interface{}
	Insert(key string, value interface{})
}
