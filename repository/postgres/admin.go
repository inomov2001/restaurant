package postgres

import (
	"time"

	structs "github.com/gokurs/Projects/restaurant/repository/struct"
)

func (u PostgresRepo) CountUsers() (map[int]int, error) {
	var SumUsers map[int]int
	rows, err := u.db.Query("SELECT tables_id FROM chek WHERE exit_time>$1", time.Now().Add(-time.Hour*24))
	if err != nil {
		return map[int]int{}, err
	}
	var countTables map[string]int //aynan bitta tabledagi odamlar soni uchun
	for rows.Next() {
		var table_id string
		if err := rows.Scan(&table_id); err != nil {
			return nil, err
		}

		val, ok := countTables[table_id]
		if !ok {
			countTables[table_id] = 1
		} else {
			countTables[table_id] += val //bitta tabledagi odamlar soni uchun
		}
	}
	s := 0
	for id, val := range countTables {
		row := u.db.QueryRow("SELECT number FROM tables WHERE id=$1", id)
		var number int /// table raqami
		row.Scan(&number)
		SumUsers[number] = val // bitta tabledagi odamlar soni
		s += val               // ummiy bir kunlik odamlar soni
	}
	SumUsers[0] = s // nolinchi indeksda hardoim ummiy odamlar soni turadi
	return SumUsers, nil
}

func (u PostgresRepo) CountSum() (map[int]int, error) {
	var CountSum map[int]int // ummiy bir kunlik budjet
	rows, err := u.db.Query("SELECT chek_id FROM basket WHERE exit_time>$1", time.Now().Add(-time.Hour*24))
	if err != nil {
		return nil, err
	}
	var countChek map[string]int //ayni bir chekdagi budjet
	for rows.Next() {
		var chek_id string
		if err := rows.Scan(&chek_id); err != nil {
			return nil, err
		}
		val, ok := countChek[chek_id]
		if !ok {
			countChek[chek_id] = 1
		} else {
			countChek[chek_id] += val
		}
	}
	s := 0
	for id, val := range countChek {
		row := u.db.QueryRow("SELECT product_id FROM basket WHERE id=$1", id)
		var product_id int
		row.Scan(&product_id)
		CountSum[product_id] = val
		s += val
	}
	CountSum[0] = s
	return CountSum, nil
}

func (u PostgresRepo) ProductList() ([]structs.Product, error) {
	rows, err := u.db.Query("SELECT *FROM products")
	if err != nil {
		return nil, err
	}
	products := structs.Product{}
	for rows.Next() {
		if err := rows.Scan(&products.Id, products.Name, products.CreatedAt.AddDate); err != nil {
			return nil, err
		}
	}
	return []structs.Product{}, err
}

func (u PostgresRepo) UpdateProduct(id string) error {
	if _, err := u.db.Exec("UPDATE products SET created_at WHERE id=$1", id); err != nil {
		return err
	}
	return nil
}
