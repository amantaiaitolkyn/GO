package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello world")
	var name = "Aitolkyn"

	var age int = 12
	var point = 4.23444
	// var num  = 20
	fmt.Println(name +"\n"+ "Amantai",age)
	fmt.Printf("%T \n",point)
	fmt.Printf("%.3f \n",point)

	var ran = rand.Intn(32) + 1
	fmt.Println(ran)

}