package main

import "fmt"

func main(){
	//a:= 10 it's an expression
	/*this is an anonymous function without any fuction name: func(a int, b int){//this block of code indicate assign fuction in a variable
		c:=a+b
		fmt.Println(c)
	}*/
	add:= func(a int, b int){//this block of code indicate assign fuction in a variable
		c:= a+b
		fmt.Println(c)
	}
	add(4,7)
}