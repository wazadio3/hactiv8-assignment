package main

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Product struct {
	Nama  string `json:"nama"`
	Brand string `json:"brand"`
	Stok  int    `json:"stok"`
	Price int    `json:"price"`
}

type Resp struct {
	Payload interface{} `json:"payload"`
}
