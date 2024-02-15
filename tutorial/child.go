package main

import (
	"fmt"
)

func init(){
	println("child init()")
}

func anything() {
	println("anything()")
	defer println("defer anything()")
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