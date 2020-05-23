package mystring

import (
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

func init() {
	fmt.Println("---------welcome mystring init----------")
}

func TestMystring() {
	var str = "/111/222/333/444"
	var strSep = "/"

	res := strings.Split(str, strSep)


	res2 := res[1:len(res)-1]
	fmt.Printf("mystring len:%d res %v\n", len(res2), res2)



	str = "18501961767"
	rep := regexp.MustCompile(`^(\d{3})(\d{4})(\d{4})`)
	res1 := rep.ReplaceAllString(str,"$1****$3")
	fmt.Printf("mystring %s replace %s\n",str,res1)


}

func TestMystringParse(){
	type Response2 struct {
		No     int    `json:"errno" mcpack:"errno"`
		Result interface{} `json:"result" mcpack:"result"`
	}
	//var errStr = "{\"errno\":-1,\"result\":{\"value\":\"please do not repeat the same request!\"}}"
	var errStr = "{\"errno\":0,\"result\":{\"62\":\"default\",\"100\":\"del\"}}"
	var res Response2 = Response2{}


	err := json.Unmarshal([]byte(errStr),&res)
	if err != nil {
		fmt.Printf("mystring Unmarshal err: %v ,errStr: %s\n", err, errStr)
		return
	}

	if res.No != 0{
		return
	}

	value,ok := res.Result.(map[int]interface{})
	if !ok {
		fmt.Printf("mystring assert fail type: %v\n",reflect.TypeOf(res.Result))
		return
	}

	if len(value) > 0 {
		for _, value := range value {
			strValue,ok := value.(string)
			if !ok{
				fmt.Printf("mystring assert fail value type: %v\n",reflect.TypeOf(value))
				continue
			}
			if strValue == "del" {
				fmt.Printf("mystring value has del\n")
				break
			}
		}

	}

	fmt.Printf("mystring value: %v\n",value)



}

func TestMystring2Byte(){
	var str string = "helloworld"
	var strByte string = string([]byte(str)[:5])
	var strByte2 string  = string([]byte(str)[5:])
	fmt.Printf("mystring child byte: %v %v\n",strByte,strByte2)
}

func TestMystring3(){
	var str string= "{\"62\":\"default\",\"100\":\"del\"}"
	var strMap map[int]string

	err := json.Unmarshal([]byte(str),&strMap)
	if err != nil {
		fmt.Printf("mystring err: %v\n",err)
		return
	}
	fmt.Printf("mystring map:%+v\n",strMap)
}
