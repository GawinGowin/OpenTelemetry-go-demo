package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("defer main() first")
	anything()
	defer fmt.Println("defer main() last")
}

func init() {
	fmt.Println("init1()")
}

func init() {
	fmt.Println("init2()")
}