package main

import (
	"fmt"
)

func init(){
	println("child init()")
}

func anything() {
	defer println("defer child() first")
	println("anything()")
	defer println("defer child() last")
}

func testRecover(src interface{}){
	defer func() {
		if x := recover(); x != nil {
			switch x.(type) {
				case int:
					fmt.Println("panic: int")
				case string:
					fmt.Println("panic: string")
				default:
					fmt.Println("panic: unknown")
			}
		}
	} ()
	panic (src)
}