package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func SetName(o interface{}, newName string) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("Can not be set!")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString(newName)
	}

}

func SetId(o interface{}, newId int64) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("Can not be set!")
		return
	} else {
		v = v.Elem()
		fmt.Println("v.Elem(", v, ")")
	}

	f := v.FieldByName("Id")
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}

	if f.Kind() == reflect.Int {
		f.SetInt(newId)
	}

}

func main() {

	u := User{121212, "bo", 23}
	SetName(&u, "ok")
	SetId(&u, 123456)
	fmt.Println(u)

}
