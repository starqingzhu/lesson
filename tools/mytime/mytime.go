package tools

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

const (
	completionTwo	= 2
)


func init()  {

}

func RenameFile(fileName string)string{
	var res string = ""

	reg,err := regexp.Compile("\\.")
	if err != nil{
		return res
	}
	res = reg.ReplaceAllString(fileName,GetFileDateName_Day()+".")

	return res
}

func TestRenameFile(){
	//path := "./log/"
	file := "hello.log"

	 reName := RenameFile(file)

	fmt.Printf("Rename file,str=%s\n",reName)
}

func GetFileDateName_Day() string{
	tm := time.Now()
	var dateStr string
	dateStr += completionTwoNumber(tm.Year())
	dateStr += completionTwoNumber(int(tm.Month()))
	dateStr += completionTwoNumber(int(tm.Day()))

	return dateStr
}

func GetFileDateName_Hour() string{
	tm := time.Now()
	var dateStr string
	dateStr += completionTwoNumber(tm.Year())
	dateStr += completionTwoNumber(int(tm.Month()))
	dateStr += completionTwoNumber(int(tm.Day()))
	dateStr += "_"
	dateStr += completionTwoNumber(int(tm.Hour()))

	return dateStr
}

func GetFileDateName_Minute() string{
	tm := time.Now()
	var dateStr string
	dateStr += completionTwoNumber(tm.Year())
	dateStr += completionTwoNumber(int(tm.Month()))
	dateStr += completionTwoNumber(int(tm.Day()))
	dateStr += "_"
	dateStr += completionTwoNumber(int(tm.Hour()))
	dateStr += completionTwoNumber(tm.Minute())

	return dateStr
}

func GetFileDateName_Second() string{
	tm := time.Now()
	var dateStr string
	dateStr += completionTwoNumber(tm.Year())
	dateStr += completionTwoNumber(int(tm.Month()))
	dateStr += completionTwoNumber(int(tm.Day()))
	dateStr += "_"
	dateStr += completionTwoNumber(int(tm.Hour()))
	dateStr += completionTwoNumber(tm.Minute())
	dateStr += completionTwoNumber(tm.Second())

	return dateStr
}

func completionTwoNumber(num int)string{
	strNum := strconv.Itoa(num)
	if len(strNum) < completionTwo{
		strNum = "0" + strNum
	}

	return strNum
}