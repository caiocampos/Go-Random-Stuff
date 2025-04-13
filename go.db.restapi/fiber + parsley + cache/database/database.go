package database

type Database[T any] interface {
	Connect() error
	Disconnect() error
	Get() *T
}
