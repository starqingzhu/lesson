package myreflect

import (
	"bytes"
	"errors"
	"fmt"
	"lession/log"
	"reflect"
)

const (
	moduleName = "reflect"
)

func init() {
	log.PrintInit(moduleName)
}

func ConstructQueryStmt(obj interface{}) (stmt string, err error) {
	//仅支持struct或struct指针类型
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		err = errors.New("only struct is supported")
		return
	}

	buffer := bytes.NewBufferString("")
	buffer.WriteString("SELECT ")

	if typ.NumField() == 0 {
		err = fmt.Errorf("tye type[%s] has no fields", typ.Name())
		return
	}

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if i != 0 {
			buffer.WriteString(", ")
		}
		column := field.Name
		if tag := field.Tag.Get("orm"); tag != "" {
			column = tag
		}

		buffer.WriteString(column)
	}

	stmt = fmt.Sprintf("%s FROM %s", buffer.String(), typ.Name())
	return
}

func PrintReflect() {
	stmt, err := ConstructQueryStmt(&Producet{})
	if err != nil {
		fmt.Println("construct query stmt for product error:", err)
		return
	}
	fmt.Println(stmt)
	stmt, err = ConstructQueryStmt(&Person{})
	if err != nil {
		fmt.Println("construct query stmt for product error:", err)
		return
	}
	fmt.Println(stmt)
}

func PrintReflect2() {
	var temp = 5
	typ := reflect.TypeOf(temp)
	val := reflect.ValueOf(temp)

	fmt.Printf("%s[%d]\n", typ.Kind(), val.Int())
}
