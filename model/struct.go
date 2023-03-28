package model

type Product struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	CategoryID string `json:"category_id"`
}

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Balance int    `json:"balance"`
}

type Company struct{
	Name string `json:"Name"`
	Balance int `json:"Balance"`
}
type Products struct{
	Pr []Product
}