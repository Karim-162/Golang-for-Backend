package main

import "fmt"

func main(){
	var a,b int16= -32000,32000
	fmt.Println(a,b)
	var hamster,dna rune='🐹','🧬' 
	//rune hocche unicode charcter int32 er alias
	fmt.Printf("%c %c",hamster,dna)
	fmt.Printf("\n")
	fmt.Println(string(hamster))
	fmt.Println(string(dna))
	fmt.Println(string(hamster) + string(dna))

	

}