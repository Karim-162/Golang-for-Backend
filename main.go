package main
import "fmt"
import "example.com/mathlib"

/*
PS F:\Phitron reset\Golang-for-Backend> go mod init example.com
go: creating new go.mod: module example.com
go: to add module requirements and sums:
        go mod tidy
*/

func main(){
	
	fmt.Println("showing custom package")
	mathlib.Add(4,7)

}