package main
import "fmt"
import "net/http" //server banabo ami tai amar ekta built in package lagbe 

func helloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello World")
}
func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hi i am shafayet and I am a software engineer")
}

func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/hello",helloHandler)
	mux.HandleFunc("/about",aboutHandler)

	fmt.Println("Server Running On:3000")

	err:= http.ListenAndServe(":3000",mux)
	if err!=nil{
		fmt.Println("Error Starting the Server",err)
	}

}