package main

import "fmt"

var (
	country = "Thailand"   //global variable
	province = "trat"
	poscode = 23000
)
func main() {
	fmt.Println(country);
	y := false //local variable
	fmt.Println(y)

	println(country);
	println(province);
	println(poscode);
}