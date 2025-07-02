package main
import "fmt"
func welcome(){
	fmt.Println("Welcome to simple Application")
}
func getusername()string{
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	return name
}
func taketwonumbers()(int, int){
	var num1 int
	var num2 int
	fmt.Println("enter number 1: ")
	fmt.Scanln(&num1)
	fmt.Println("enter number 2: ")
	fmt.Scanln(&num2)
	return num1, num2

}
func add(num1 int,num2 int)int{
		sum:= num1+num2
		return sum
}
func printresult(name string,sum int){
	fmt.Println("Hello ",name)
	fmt.Println("The result of num1 and num2 is: ", sum)

}
func goodbyemessage(){
	fmt.Println("Thank you for using our application. Stay with Us")
	fmt.Println("GOOD BYE!!!")
}
func main(){
	welcome()
	name:=getusername()
	num1, num2:=taketwonumbers()
	sum:= add(num1,num2)
	printresult(name,sum)
	goodbyemessage()
	
}