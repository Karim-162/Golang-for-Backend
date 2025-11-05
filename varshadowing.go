package main
import "fmt"

var a=10

func main(){
	age:=47
	if age>18{
		a:=5
		fmt.Println(a)
	}
	fmt.Println(a)
}