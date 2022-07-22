package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	n := 1
	nbytes := i2bin(&n)

	for i := 0; i < len(nbytes); i++ {
		fmt.Printf("value: %v,addr: %p", nbytes[i], &nbytes[i])
	}
}

func i2bin(i *int) []byte {
	sh := &reflect.SliceHeader{Len: 8, Cap: 8, Data: uintptr(unsafe.Pointer(i))}
	return *(*[]byte)(unsafe.Pointer(sh))
}
