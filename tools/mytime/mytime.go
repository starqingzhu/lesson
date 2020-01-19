package tools

import (
	"regexp"
	"strconv"
	"time"
)

//常量
const (
	completionTwo	= 2					//补全日期字符长度为2，eg：1补全为01
)

//时间级别范围
const (
	Day = iota + 1
	Hour
	Minute
	Second
)

//存放时间级别
var DateLevel map[int64]string

//函数别名
type (
	renameFunc func()string
)



func init()  {
	DateLevel = map[int64]string{
		Day		:	"Day",
		Hour	:	"Hour",
		Minute	:	"Minute",
		Second	:	"Second",
	}
}

func RenameFile(fileName string,fc renameFunc)string{
	var res string = ""

	reg,err := regexp.Compile("\\.")
	if err != nil{
		return res
	}
	res = reg.ReplaceAllString(fileName,"_"+fc()+".")

	return res
}

//func TestRenameFile(){
//	path := "./log/"
//	file := "hello.log"
//
//	 reName := RenameFile(file,GetFileDateName_Day)
//
//	fmt.Printf("Rename file,str=%s\n",path+reName)
//}

func GetFileDateName(dateLevel int64)string{
	tm := time.Now()
	var dateStr string
	dateStr += completionTwoNumber(tm.Year())
	dateStr += completionTwoNumber(int(tm.Month()))
	dateStr += completionTwoNumber(int(tm.Day()))
	if dateLevel == Day{
		return dateStr
	}
	dateStr += "_"
	dateStr += completionTwoNumber(int(tm.Hour()))
	if dateLevel == Hour{
		return dateStr
	}
	dateStr += completionTwoNumber(tm.Minute())
	if dateLevel == Minute{
		return dateStr
	}
	dateStr += completionTwoNumber(tm.Second())
	if dateLevel == Second{
		return dateStr
	}
	return dateStr

}

func GetFileDateName_Day() string{
	return GetFileDateName(Day)
}

func GetFileDateName_Hour() string{
	return GetFileDateName(Hour)
}

func GetFileDateName_Minute() string{
	return GetFileDateName(Minute)
}

func GetFileDateName_Second() string{
	return GetFileDateName(Second)
}

func completionTwoNumber(num int)string{
	strNum := strconv.Itoa(num)
	if len(strNum) < completionTwo{
		strNum = "0" + strNum
	}

	return strNum
}