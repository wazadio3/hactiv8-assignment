package main

import (
	"fmt"
	"net/http"
)

const PORT = ":3000"

func main() {
	http.HandleFunc("/register", Register)
	http.HandleFunc("/users", GetAllUser)
	http.HandleFunc("/add", AddNewProduct)
	http.HandleFunc("/products", GetAllProduct)
	// http.HandleFunc("/brand", GetProductByBrand)

	fmt.Println("Server Start")
	http.ListenAndServe(PORT, nil)
}
