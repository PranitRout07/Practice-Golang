package models

type Products struct {
	Id int `json:"id"`
	ProductName string `json:"productname"`
	Price float64 `json:"price"`
}
