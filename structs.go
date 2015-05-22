package main

/*
	#include <stdlib.h>
	struct Person {
	    char *name;
	    int age;
	    int height;
	    int weight;
	};
*/
import "C"
import "fmt"

type p C.struct_Person

func main() {

	person := p{C.CString("Giorgis"), 30, 6, 175}
	fmt.Println(person)
	fmt.Println(C.GoString(person.name))
	fmt.Println(person.age)
	fmt.Println(person.height)
	fmt.Println(person.weight)
}
