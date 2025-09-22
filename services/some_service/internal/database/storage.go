package database

type Storage interface {
	AutoMigrate() error
}
