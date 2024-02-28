package main

import (
	"fmt"
)

type Stringify interface {
	ToString() string
}

type Person struct {
	Name string
	Age int
} 

func (p *Person) ToString() string {
	return fmt.Sprintf("%s(%d)", p.Name, p.Age)
}

func main() {
	v := Stringify(&Person{"Maria", 20})
	fmt.Println(v.ToString())
}
