package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var i int = 10;
	var s string = "Thailand";
	var a int;
	var b int; 
	var c = a+b;
	fmt.Println(i);
	fmt.Println(s);
	fmt.Print(reflect.TypeOf(c) ,c);
	
}