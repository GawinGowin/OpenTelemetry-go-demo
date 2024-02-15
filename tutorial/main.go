package main

import (
	"fmt"
)

func main() {
	fmt.Println(A, B, C, D)
	if x, y := 1, 10; x < y {
		fmt.Printf("x(%d) is less than y(%d)\n", x, y)
	}
	var x int = 5
	var y int = 10
	if n := x * y; n%2 == 0 {
		fmt.Printf("n(%d) is even\n", n)
	}

	a := [3]string{"a", "b", "c"}

	fmt.Println(plus(x, y, 3.141592))
	fmt.Println(a)
	anything(1)

	i := 10
	for i > 0 {
		fmt.Print("a")
		i --;
	}
	fmt.Println("")
	switch y := x * x + 10;{
		case y > 10:
			fmt.Println("> 10")
		case y < 10:
			fmt.Println("< 10")
		default:
			fmt.Println("== 10")
	}
	L:
		for i:=1; i<=3; i++ {
			for j:=1; j<=3; j++ {
				if j > 1 {
					continue L
				}
				fmt.Printf("%d * %d = %d\n", i, j, i*j)
			}
			fmt.Println("Don't print this line")
		}
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	defer fmt.Println("defer3")
	// panic("runtime error")
	testRecover(1)
	println(func (x, y int)(int) {return x + y} (5, 10))

	// go sub()
	// for {
	// 	fmt.Println("main loop")
	// }
	fmt.Println("main()")
}

func init() {
	fmt.Println("init1()")
}

func init() {
	fmt.Println("init2()")
}

// func sub(){
// 	for {
// 		fmt.Println("sub loop")
// 	}
// }

func plus(x, y int, z float32) int {
	fmt.Println(z)
	defer fmt.Println("defer plus")
	return (x + y)
}

func anything(a interface{}){
	fmt.Println(a)
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

// Go言語では識別子の一文字目が大文字ならば外部パッケージからもアクセス可能
const (
	A = iota // Global
	B = iota // Global
	C
	D
	e //local
	f //local
)
