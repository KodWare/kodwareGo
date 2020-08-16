package main

import (
	"C"
	"fmt"
)

func main() {
	abc := C.the_answer()
	fmt.Println(abc)
}
