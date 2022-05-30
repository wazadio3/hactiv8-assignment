package models

type Request struct {
	Order
	Items []Item `json:"items"`
}
