package main

import "fmt"

func main() {
	var a float64 = 90.0 //Static Type
	a = 50.5
	var b string = "welcome";
	 y := 14 //Dynamic Type
	 
	fmt.Print(b)
	fmt.Println(a)
	fmt.Println(y)
	fmt.Printf("is of type %T\n" ,a)
	fmt.Printf("is of type %T" ,a)
}