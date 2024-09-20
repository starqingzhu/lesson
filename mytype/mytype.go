package mytype

import (
	"fmt"
	"reflect"
)

/*

引用类型:切片、映射、通道、接口和函数
内置类型:字符串、bool、数值类型
结构类型:

用户自定义类型
方法
参数: 指针和值
接口+多态:指代码根据类型的具体实现采取不同行为的能力
组合+扩展类型
标识符
接口:是用来定义行为的类型。这些被定义的行为不由接口直接实现，而是通过方法由用户定义的类型实现。


嵌入类型:目的代码复用
	注意:内部类型+外部类型，内部类型的方法在外部类型没有实现时，自动提升为外部类型，一旦外部类型实现了同样方法，内部类型则不提升为外部类型也不会消失。
	私有属性包内可见，包外不可见

公开和未公开的标识符:
	1.公开或者未公开的标识符，不是一个值
	2.短变量声明操作符，有能力捕获引用的类型，并创建一个未公开的类型变量

问题：
结构体类型的某个字段声明中只有一个类型名，那么该字段代表了什么？
	Go 语言规范规定，如果一个字段的声明中只有字段的类型名而没有字段的名称，那么它就是一个嵌入字段，也可以被称为匿名字段。
	我们可以通过此类型变量的名称后跟“.”，再后跟嵌入字段类型的方式引用到该字段。也就是说，嵌入字段的类型既是类型也是名称。
	注意：只要名称相同，无论这两个方法的签名是否一致，被嵌入类型的方法都会“屏蔽”掉嵌入字段的同名方法。

*/

func init() {
	fmt.Println("---------welcome mytype init----------")
}

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

	//嵌入类型
	Manager struct {
		User
		Level string
	}

	//基于已有类型重新定义类型
	Duration int64
)

func (u *User) Printp() {
	fmt.Printf("mytype User:%+v\n", u)
}

func (ad Admin) Print() {
	fmt.Printf("mytype Admin:---->\npserson:%+v\nlevel:%s\n", ad.person, ad.level)
}
func (ad *Admin) Printp() {
	fmt.Printf("mytype *Admin:---->\npserson:%+v\nlevel:%s\n", ad.person, ad.level)
}

func Init() {
	fmt.Println("---------type init----------")
	var user User = User{
		name:          "Mr.sun",
		email:         "starqingzhu@163.com",
		ext:           0,
		intprivileged: true,
	}
	fmt.Printf("mytype user实例:\n%+v\n", user)

	var user1 User
	fmt.Printf("mytype user实例0值:\n%+v\n", user1)

	ad := Admin{
		person: user,
		level:  "super",
	}
	fmt.Printf("mytype admin实例:\n%+v\n", ad)

	//var tm Duration
	//tm = int64(1000)
	//fmt.Printf("redefine system type int64,tm=%d\n",tm)
}

func Func() {
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

type Notifier interface {
	Notify()
}

type (
	IUser struct {
		name  string
		email string
	}
	IAdmin struct {
		name  string
		email string
	}
)

func (u *IUser) Notify() {
	fmt.Printf("mytype Sending user email to %s<%s>\n", u.name, u.email)
}

func sendNotification(n Notifier) {
	n.Notify()
}

func TestInterfaceNotify() {
	fmt.Println("---------test interface----------")
	u := IUser{
		name:  "sun",
		email: "starqingzhu@163.com",
	}
	fmt.Printf("IUser:\n%+v\n", u)
	sendNotification(&u)
}

func (a *IAdmin) Notify() {
	fmt.Printf("mytype Sending user email to %s<%s>\n", a.name, a.email)
}

func TestMutiInterfaceNotify() {
	fmt.Println("---------test muti interface----------")
	u := IUser{
		name:  "sun",
		email: "starqingzhu@163.com",
	}
	a := IAdmin{
		name:  "a_sun",
		email: "starqingzhu@163.com",
	}
	fmt.Printf("mytype IUser:\n%+v\nIAdmin:\n%+v\n", u, a)
	sendNotification(&u)
	sendNotification(&a)

}

func TestInsertion() {
	fmt.Println("---------test insertion----------")
	m := Manager{
		User: User{
			name:          "m.Mr.sun",
			email:         "m.starqingzhu@163.com",
			ext:           0,
			intprivileged: true,
		},
		Level: "1",
	}
	fmt.Printf("mytype Manager:\n%+v\n", m)
	m.User.Printp()
}

type alertCounter int

func New(value int) alertCounter {
	return alertCounter(value)
}

type (
	user1 struct {
		Name  string
		email string
	}
	Admin1 struct {
		user1
		Level int
	}

	user2 struct {
		Name  string
		Email string
	}
	Admin2 struct {
		user2
		Level int
	}
)

func TestInsertionFloat() {
	fmt.Println("---------test insertion float----------")
	var x1 int64 = 10
	var x2 int64 = 100.00
	fmt.Printf("mytype 浮点数:%.2f\n", float64(x1)/float64(x2))
}

type (
	People struct {
		Length int
		Sex    int
		weight int
	}
	Title struct {
		Level int
	}
	Boss struct {
		People
		Title
	}
)

func (p *People) String() string {
	return fmt.Sprintf("length:%d sex:%d weight:%d", p.Length, p.Sex, p.weight)
}
func (t *Title) String() string {
	return fmt.Sprintf("level:%d", t)
}

func (p *Boss) String() string {
	return p.People.String() + " " + p.Title.String()
}

func TestInsert2() {
	boss := Boss{
		People: People{Length: 175, Sex: 1},
		Title:  Title{1},
	}
	boss.weight = 1

	fmt.Printf("mytype %s\n", boss.String())
}

func TestConv() {
	var f1 float64 = 10
	f2 := int64(f1)

	fmt.Printf("mytype f1 type:%v value:%f f2 type:%v value:%d\n", reflect.TypeOf(f1), f1, reflect.TypeOf(f2), f2)
}
