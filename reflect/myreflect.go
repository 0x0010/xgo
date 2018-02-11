package main

import (
	"fmt"
	"reflect"
)

type MessageBody struct {
	msgType string
	uid     string
	level   string
	nn      string
	Txt     string
	bnn     string
	bl      string
}

func main() {
	s := "this is string"
	sType := reflect.TypeOf(s)
	fmt.Println(sType)
	fmt.Println(sType.Kind())
	fmt.Println(reflect.String)

	msg := MessageBody{}
	msg.Txt = "Hello World"
	msgType := reflect.TypeOf(msg)
	msgTypeKind := msgType.Kind();
	fmt.Println(msgType.Name())
	fmt.Println(msgTypeKind.String())
	fmt.Println(msgTypeKind)
	fmt.Println(msgType)

	fmt.Println("=========")

	fmt.Println(msgType.NumField())

	fieldNum := msgType.NumField();
	for i := 0; i < fieldNum; i++ {
		fmt.Println("Field", i, msgType.Field(i).Name)
	}

	fmt.Println("=========")
	fmt.Println(msg)
	reflect.Indirect(reflect.ValueOf(&msg)).FieldByName("Txt").SetString("World Hello")
	fmt.Println(msg)

}
