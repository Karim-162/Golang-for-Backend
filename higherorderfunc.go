package main

import "fmt"

/*
there are 3 conditions if any one of them are true then it can be called
a higher order function and it is also called first class function
1.If a function has received as a parameter
2.If the function has return
3.Both
[imp: a variable which can store any type and also function in it is called first class citizen.
EX: a:= func(){}]
*/
func processOperation(a int, b int, op func(x int, y int)){
	op(a,b)
}//Here we can see that a function has passed as parameter in another function so this is called higher order function. It is also called Callback function as we are sending a function inside another fuction

func add(n1 int, n2 int){
	sum:=n1+n2
	fmt.Println(sum)
}

func main(){
	processOperation(4,7,add)
}