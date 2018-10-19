package gosingleton

import (
	"reflect"
	"sync"
)

type TypeList map[string]reflect.Type

var instance interface{}
var once sync.Once
var typeList TypeList

func init() {
	typeList = make(TypeList)
}

func RegisterType(typeName string, typeInstance reflect.Type) {
	typeList[typeName] = typeInstance
}

func GetTypeInfo(object interface{}) reflect.Type {
	return reflect.TypeOf(object)
}

func GetInstance(typeName string) interface{} {
	once.Do(func() {
		v := reflect.New(typeList[typeName]).Elem()
		if !v.CanAddr() {
			return
		}
		instance = v.Addr().Interface()
	})

	return instance
}
