package main

/*
#include <stdio.h>
#include <stdint.h>
struct columns {
    int column1;
    int column2;
    int column3;
};

int sum_columns(struct columns a) {
	return a.column1 + a.column2 + a.column3;
};

int sum_vals(int a, int b){
	return a + b;
}
*/
import "C"
import "fmt"

func main() {

	c := C.struct_columns{15, 30, 45}

	sum := C.sum_columns(c)
	fmt.Println(sum)

	var a int = 15
	var b int = 30

	s := C.sum_vals((C.int)(a), (C.int)(b))
	fmt.Println(s)
}
