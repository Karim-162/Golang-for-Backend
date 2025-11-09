package main
import "fmt"

type User struct{
	name string
	age int
	salary float64
	//favfood []string
}

func printuser(usr *User){
	fmt.Println(usr)
	fmt.Println(usr.name)
	fmt.Println(usr.age)
	fmt.Println(usr.salary)
}

func main(){
	user:= User{
		name: "shafayet",
		age: 25,
		salary: 250.50,
	}
	p:=&user

	printuser(p)

}