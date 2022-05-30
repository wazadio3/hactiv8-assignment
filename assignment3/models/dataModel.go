package models

type Data struct {
	Status struct {
		Water string `json:"water"`
		Wind  string `json:"wind"`
	} `json:"status"`
}
