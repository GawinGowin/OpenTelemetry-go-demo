package main

import (
	"fmt"
)

// Go言語では識別子の一文字目が大文字ならば外部パッケージからもアクセス可能
const (
	A = iota // Global
	B = iota // Global
	C
	D
	e //local
	f //local
)

func main() {
	fmt.Println(A, B, C, D)
	a := [3]string{"a", "b", "c"}
	fmt.Println(a)
	i := 10
	for i > 0 {
		fmt.Print("a")
		i --;
	}
	fmt.Println("")
	x := 10
	switch y := x * x + 10;{
		case y > 10:
			fmt.Println("> 10")
		case y < 10:
			fmt.Println("< 10")
		default:
			fmt.Println("== 10")
	}
	println(func (x, y int)(int) {return x + y} (5, 10))
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
	// testRecover(1)

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

func plus(x, y int, z float32) int {
	fmt.Println(z)
	defer fmt.Println("defer plus")
	return (x + y)
}
