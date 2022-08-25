package structs

import "time"

type Food struct {
	Id         string
	Name       string
	Price      uint32
	Category   uint8
	ProductsId []string
}

type Salad struct {
	Id    string
	Name  string
	Price uint32
}

type Drinks struct {
	Id    string
	Name  string
	Price uint32
}

type MenyuJson struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
}

type Product struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Table struct {
	Id     string `json:"id"`
	Number uint8  `json:"number"`
	Buys   bool   `json:"buys"`
}

type ShopStruct struct {
	TableId   string `json:"table_id"`
	FoodId    string `json:"food_id"`
	SaladId   string `json:"salad_id"`
	DrinckId  string `json:"drinck_id"`
}
