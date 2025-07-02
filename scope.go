package main
import "fmt"

var a = 10
var b = 20

func add(x int,y int){
	z:=x+y
	fmt.Println(z)
}

func main(){
	p:=30
	q:=40

	add(a,p)
	add(p,q)
	add(p,b)
}