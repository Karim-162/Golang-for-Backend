package main

import "fmt"


type User struct{
	Name string //these Name and Age are member variable or property 
	Age int
}

func (usr User)printdetails(){//receiver func that only works for custom type making
	fmt.Println("Name: ",usr.Name)
	fmt.Println("Age: ",usr.Age)
}

func main(){
	var user1 User
	/*eije amra User type er Name and Age egula k nicchi shegulake Instance o bola hoi
	to amra line 16 er jonno bolte pari j amra User type er jonno ekta instance or
	object create korsi
	*/
	user1=User{
		Name:"habib",
		Age:30,
	}
	//printdetails(user1)
	user1.printdetails()

	
	user2:=User{
		Name:"Sazin",
		Age:25,
	}
	fmt.Println("Name: ",user2.Name)
	fmt.Println("Age: ",user2.Age)
}