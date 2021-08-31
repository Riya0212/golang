package main

import "fmt"

func main() {
	var a int
	var b, c int8
	a = 120
	b = 24
	c = 10
	fmt.Printf("%v, %T", a, a)
	fmt.Println("b+c= ", b+c)
	fmt.Println("b+c= ", b-c)
	fmt.Println("b+c= ", b*c)
	fmt.Println("b+c= ", b/c)
	fmt.Println("b+c= ", b%c)
	fmt.Printf("a+b= %v,%T ", a+int(b), a+int(b))
}
