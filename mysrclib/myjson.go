package mysrclib

import (
	"encoding/json"
	"fmt"
)

func init(){
	fmt.Println("---------welcome myjson init----------")
}


type (
	Person struct {
		Name	string		`json:"name"`
		Age		int64		`json:"age"`
		Sex		int			`json:"sex"`
	}
)

func TestJsonDecodePerson(){
	var JPerson = `{
		"Name": "Mr.sun",
		"Age":  28,
		"Sex":  0
	}`

	var person Person
	err := json.Unmarshal([]byte(JPerson),&person)
	if err != nil{
		MyError.Printf("decode JPerson:%+v failed,err=%v\n",JPerson,err)
		return
	}
	MyInfo.Printf("decode JPerson success,JPerson=%+v\n",person)
}

func TestJsonEncodePerson(){
	var JPerson = Person{
		Name: "Mr.ren",
		Age:  30,
		Sex:  1,
	}

	person,err := json.Marshal(JPerson)
	//person,err := json.MarshalIndent(JPerson,"","	")
	if err != nil{
		MyError.Printf("JPerson:%+v,Encode err:%v\n",JPerson,err)
		return
	}
	MyInfo.Printf("encode JPerson success,JPerson=%+v\n",string(person))


}