package main

/*
enum levels {
  low1,
  medium,
  high
};

typedef enum {
	LOW = 0,
    MEDIUM = 1,
    HIGH = 2
} security;

*/
import "C"
import "fmt"

func main() {

	cc := new(C.enum_levels)
	fmt.Println(*cc)

	fmt.Println(C.MEDIUM)

}
