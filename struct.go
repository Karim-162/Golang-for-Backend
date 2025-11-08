package main

import "fmt"


func processOperation(a int, b int, op func(x int, y int)){
	op(a,b)
}

func add(n1 int, n2 int){
	sum:=n1+n2
	fmt.Println(sum)
}

func main(){
	processOperation(4,7,add)
}