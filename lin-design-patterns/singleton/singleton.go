package singleton

import "sync"

var once sync.Once
var instance *DB

type DB struct {
}

func ConnectDB() *DB {
	return &DB{}
}

func GetDB() *DB {
	once.Do(func() {
		instance = ConnectDB()
	})
	return instance
}
