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
	mux:=http.NewServeMux()//NewServeMux() eta struct type er and eta ekta pointer mux variable a rakhtese
	mux.HandleFunc("/hello",helloHandler)// HandleFunc eta ekta receiver function jeta argument hishebe string ar ekta handler function nei
	mux.HandleFunc("/about",aboutHandler)// HandleFunc eta ekta receiver function jeta argument hishebe string ar ekta handler function nei
	//mux ta hocche Router
	//"/hello" "/about" ei string pattern guloke diye call korake bola hoi route
	//so main thing is amra router diye multiple route banate pari
	
	fmt.Println("Server Running On:3000")

	err:= http.ListenAndServe(":3000",mux)
	if err!=nil{
		fmt.Println("Error Starting the Server",err)
	}

}