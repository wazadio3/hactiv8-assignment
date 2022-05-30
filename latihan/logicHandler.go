package main

import (
	"encoding/json"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		json.NewEncoder(w).Encode("error")
		return
	}

	var u User
	json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode("OKE")

	Users = append(Users, u)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		json.NewEncoder(w).Encode("error")
		return
	}

	var res Resp
	// for _, v := range Users {
	// 	fmt.Println("Username :", v.Username)
	// 	fmt.Println("Password :", v.Password)
	// }

	res.Payload = Users
	json.NewEncoder(w).Encode(res)
}

func AddNewProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		json.NewEncoder(w).Encode("error")
		return
	}

	var isAuthorized bool
	username, password, ok := r.BasicAuth()
	if !ok {
		json.NewEncoder(w).Encode("error")
		return
	}
	for _, v := range Users {
		if (v.Username == username) && (v.Password == password) {
			isAuthorized = true
			break
		}
	}

	if isAuthorized {
		var p Product
		json.NewDecoder(r.Body).Decode(&p)
		Products = append(Products, p)

		json.NewEncoder(w).Encode("OKE")
		return
	}

	json.NewEncoder(w).Encode("UnAuthorized")

}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		json.NewEncoder(w).Encode("error")
		return
	}

	var res Resp

	brand := r.URL.Query().Get("brand")
	if brand != "" {
		var response = []Product{}
		for _, v := range Products {
			if v.Brand == brand {
				response = append(response, Product{
					Nama:  v.Nama,
					Brand: v.Brand,
					Stok:  v.Stok,
					Price: v.Price,
				})
			}
		}

		res.Payload = response

		json.NewEncoder(w).Encode(res)
		return
	}

	res.Payload = Products
	json.NewEncoder(w).Encode(res)
}

// func GetProductByBrand(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		json.NewEncoder(w).Encode("error")
// 		return
// 	}

// 	var response = []Product{}
// 	brand := r.URL.Query().Get("brand")
// 	for _, v := range Products {
// 		if v.Brand == brand {
// 			response = append(response, Product{
// 				Nama:  v.Nama,
// 				Brand: v.Brand,
// 				Stok:  v.Stok,
// 				Price: v.Price,
// 			})
// 		}
// 	}

// 	json.NewEncoder(w).Encode(Products)
// }
