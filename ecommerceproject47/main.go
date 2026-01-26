package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello World")
}
func abouthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Shafayet Karim,, an Aspiring Software Engineer")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imgUrl"`
}

var productList []Product

func getproducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, "plz give me GET request", 400)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(productList)

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", hellohandler)
	mux.HandleFunc("/about", abouthandler)
	mux.HandleFunc("/products", getproducts)

	fmt.Println("server running on :3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error Starting the Server", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "orange",
		Description: "orange is full of vitamin c",
		Price:       100,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/4/43/Ambersweet_oranges.jpg",
	}
	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "One Apple a day keeps doctor away",
		Price:       101,
		ImgUrl:      "https://hips.hearstapps.com/hmg-prod/images/apples-at-farmers-market-royalty-free-image-1627321463.jpg?crop=1xw:0.94466xh;center,top&resize=1200:*",
	}
	prd3 := Product{
		ID:          3,
		Title:       "banana",
		Description: "Banana is good for breakfast",
		Price:       102,
		ImgUrl:      "https://nutritionsource.hsph.harvard.edu/wp-content/uploads/2018/08/bananas-1354785_1920-1024x683.jpg",
	}
	prd4 := Product{
		ID:          4,
		Title:       "watermelon",
		Description: "watermelon is red",
		Price:       103,
		ImgUrl:      "https://centershealthcare.com/wp-content/uploads/2022/04/Watermelon.webp",
	}
	prd5 := Product{
		ID:          5,
		Title:       "grape",
		Description: "grapes are sour told by the FOX",
		Price:       104,
		ImgUrl:      "https://lemonsandletters.com/wp-content/uploads/2024/02/the-fox-and-the-grapes.jpg",
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
}
