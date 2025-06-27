package main
import "fmt"
func main(){
	var a=10
	b:=20
	c:=b
	b=30//to change the value dont use ":" in second time
	var d int = 40
	b="you put an int value in b at first..so the variable becomes an int type and you cant put other types in that variable"
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println(b)
}