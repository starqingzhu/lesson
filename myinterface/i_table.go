package myinterface

import (
	"fmt"
	"unsafe"
)

type nameOff int32
type typeOff int32
type textOff int32
type tflag uint8

type _type struct {
	size       uintptr
	ptrdata    uintptr // size of memory prefix holding all pointers
	hash       uint32
	tflag      tflag
	align      uint8
	fieldAlign uint8
	kind       uint8
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal func(unsafe.Pointer, unsafe.Pointer) bool
	// gcdata stores the GC type data for the garbage collector.
	// If the KindGCProg bit is set in kind, gcdata is a GC program.
	// Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
	gcdata    *byte
	str       nameOff
	ptrToThis typeOff
}
type name struct {
	bytes *byte
}
type imethod struct {
	name nameOff
	ityp typeOff
}

type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod
}

type itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32 // copy of _type.hash. Used for type switches.
	_     [4]byte
	fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}

type eface struct {
	_type *_type
	data  unsafe.Pointer
}

type iface struct {
	tab  *itab
	data unsafe.Pointer
}

func dumpEface(i interface{}) {
	ptrToEface := (*eface)(unsafe.Pointer(&i))
	fmt.Printf("interface Pointer eface:%+v\n", *ptrToEface)
	if ptrToEface._type != nil {
		fmt.Printf("\t _type: %+v\n", *(ptrToEface._type))
	}
	if ptrToEface.data != nil {
		fmt.Printf("\t data: %d\n", ptrToEface.data)
	}
}

func dumpItabOfIface(ptrToIface unsafe.Pointer) {
	p := (*iface)(ptrToIface)
	fmt.Printf("interface Pointer iface:%+v\n", *p)
	if p.tab != nil {
		fmt.Printf("\t tab: %+v\n", *(p.tab))
		fmt.Printf("\t _type: %+v\n", *(p.tab._type))
		fmt.Printf("\t data: %d\n", p.data)
	}
}

type TI int

func (t TI) Error() string {
	return "bad error"
}

func PrintInterfaceInterDetail() {
	var eif interface{} = TI(5)
	dumpEface(eif)

	var err error = TI(5)
	dumpItabOfIface(unsafe.Pointer(&err))
}
