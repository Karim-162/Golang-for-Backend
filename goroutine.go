package main

import(
	"fmt"
	"time"
)

var a = 10
const c = 11
func printhello(n int){
	fmt.Printf("hello goroutine number %d\n",n)
}

func main(){
	go printhello(1)
	go printhello(2)
	go printhello(3)
	go printhello(4)
	go printhello(5)

	fmt.Println(a," ",c)

	time.Sleep(5*time.Second)


}