package main

/*
	#include <stdlib.h>
	char* cstring = "C string example";
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {

	var gstring string = "Go string example"

	//Go to C String, Output: *C.char
	cs := C.CString(gstring)
	defer C.free(unsafe.Pointer(cs))
	fmt.Println(cs)

	//C to Go String, Output: string
	gs := C.GoString(C.cstring)
	fmt.Println(gs)

	//C string, length to Go string
	gs2 := C.GoStringN(C.cstring, (C.int)(len(gs)))
	fmt.Println(gs2)

}
