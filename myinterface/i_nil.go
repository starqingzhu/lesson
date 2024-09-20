package myinterface

func PrintNilInterface() {
	var i interface{} //空接口
	var err error     //非空接口

	println("interface:", i)
	println("interface:", err)

	println("interface i==nil:", i == nil)
	println("interface err==nil:", err == nil)
	println("interface err==i:", err == i)

}

func PrintEmptyInterface() {
	var eif1 interface{}
	var eif2 interface{}
	var n, m int64 = 17, 18

	eif1 = n
	eif2 = m

	println("empty interface eif1:", eif1)
	println("empty interface eif2:", eif2)

	eif2 = int64(17)
	println("empty interface eif1:", eif1)
	println("empty interface eif2:", eif2)
	println("empty interface eif1==eif2:", eif1 == eif2)

}

type (
	NonEmptyInterface interface {
		M1()
		M2()
	}
	TT struct {
		n int
		s string
	}
)

func (TT) M1() {}
func (TT) M2() {}

func PrintNotEmptyInterface() {
	var err1 error
	//var err2 error
	err1 = nil

	println("not empty interface err1:", err1)
	println("not empty interface err1==nil:", err1 == nil)
}
