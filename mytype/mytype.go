package mytype

import "fmt"

/*

引用类型:切片、映射、通道、接口和函数
内置类型:字符串、bool、数值类型

用户自定义类型
方法
参数: 指针和值
接口多态:指代码根据类型的具体实现采取不同行为的能力
组合+扩展类型
标识符
*/

type (
	//创建自定义类型
	User struct {
		name          string
		email         string
		ext           int
		intprivileged bool
	}

	//使用已有自定义类型组合新类型
	Admin struct {
		person User
		level  string
	}

	//基于已有类型重新定义类型
	Duration int64
)

func (ad Admin) Print(){
	fmt.Printf("Admin:---->\npserson:%+v\nlevel:%s\n",ad.person,ad.level)
}
func (ad *Admin) Printp(){
	fmt.Printf("*Admin:---->\npserson:%+v\nlevel:%s\n",ad.person,ad.level)
}

func Init() {
	fmt.Println("---------type init----------")
	var user User = User{
		name:          "Mr.sun",
		email:         "starqingzhu@163.com",
		ext:           0,
		intprivileged: true,
	}
	fmt.Printf("user实例:\n%+v\n", user)

	var user1 User
	fmt.Printf("user实例0值:\n%+v\n", user1)

	ad := Admin{
		person: user,
		level:  "super",
	}
	fmt.Printf("admin实例:\n%+v\n", ad)

	//var tm Duration
	//tm = int64(1000)
	//fmt.Printf("redefine system type int64,tm=%d\n",tm)
}

func Func(){
	var user User = User{
		name:          "Mr.sun",
		email:         "starqingzhu@163.com",
		ext:           0,
		intprivileged: true,
	}
	ad := Admin{
		person: user,
		level:  "super",
	}
	ad.Print()
	ad.Printp()
}

func Interface(){

}