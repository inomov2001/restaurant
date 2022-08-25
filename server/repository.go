package server

import (
	structs "github.com/gokurs/Projects/restaurant/repository/struct"
)

type Repository interface {
	//user1
	Food1() (structs.MenyuJson, error)                                    //get
	Food2() (structs.MenyuJson, error)                                    //get
	Food3() (structs.MenyuJson, error)                                    //get
	Salad() (structs.MenyuJson, error)                                    //get
	Drinks() (structs.MenyuJson, error)                                   //get
	OpenChek(table_id string) error                                       //post
	Shop(table_id, food_id, salad_id, drinck_id string) error //post
	Chek(table_id string) (map[string]int, error)                         //get
	//admin
	CountUsers() (map[int]int, error)        //get
	CountSum() (map[int]int, error)          //get
	ProductList() ([]structs.Product, error) //get
	UpdateProduct(id string) error           //put
	//user2
	Ingredient(id string) ([]structs.Product, error)                  //get
	Set(sum int) ([][]structs.MenyuJson, error)                       //get
	ShopCombo(food_id, salad_id, drinck_id string) error              //get
	DeleteBasket(table_id, food_id, salad_id, drinck_id string) error //delete
	TableList() ([]structs.Table, error)
}
