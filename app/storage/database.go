package storage

type Database interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}
