package main

/*
#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
*/
import "C"

func main() {
	f := new(C.FILE)
	p := C.CString("string content")
	C.fputs(p, (*C.FILE)(f))
}
