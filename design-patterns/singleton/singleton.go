package singleton

import "sync"

type SingleObject struct{}

var instance *SingleObject
var once sync.Once

func GetInstance() *SingleObject {
	once.Do(func() {
		instance = &SingleObject{}
	})
	return instance
}

func (o SingleObject) printMessage() string {
	return "Hello, world!"
}
