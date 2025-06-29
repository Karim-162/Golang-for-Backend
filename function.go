package main
import "fmt"
func add(a int, b int)int{
	sum:=a+b
	return(sum)
}
func main(){
	a:=10
	b:=20
	sum:=add(a,b)
	fmt.Println(sum)
}