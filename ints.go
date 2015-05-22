package main

/*
#include <stdio.h>
#include <stdint.h>
int ic = 5;
unsigned int uic = 7;
int16_t is = 12345;
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var ig int = 10

	igc := int(C.ic)
	fmt.Println("value:", igc, "type:", reflect.TypeOf(igc))

	icg := C.int(ig)
	fmt.Println("value:", icg, "type:", reflect.TypeOf(icg))

	icp := (*C.int)(unsafe.Pointer(&ig))
	fmt.Println("value:", reflect.ValueOf(icp), "type:", reflect.TypeOf(icp))

	uigc := uint(C.uic)
	fmt.Println("value:", uigc, "type:", reflect.TypeOf(uigc))

	i64t := int16(C.is)
	fmt.Println("value:", i64t, "type:", reflect.TypeOf(i64t))

}
