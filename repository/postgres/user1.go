package postgres

import (
	"database/sql"
	"errors"
	structs "github.com/gokurs/Projects/restaurant/repository/struct"
	"github.com/google/uuid"
)

func (u PostgresRepo) Food1() ([]structs.MenyuJson, error) {
	rows, err := u.db.Query(`SELECT id,name,price FROM food WHERE category=1`)
	if err != nil {
		//return structs.MenyuJson{}, err
		panic(err)
	}
	defer rows.Close()
	foodl := make([]structs.MenyuJson, 0)
	food1 := structs.MenyuJson{}
	for rows.Next() {
		if err = rows.Scan(&food1.Id, &food1.Name, &food1.Price); err != nil {
			//return structs.MenyuJson{}, err
			panic(err)
		}
		foodl = append(foodl, food1)
	}
	return foodl, nil
}

func (u PostgresRepo) Food2() ([]structs.MenyuJson, error) {
	rows, err := u.db.Query(`SELECT id,name,price FROM food WHERE category=2`)
	if err != nil {
		//return structs.MenyuJson{}, err
		panic(err)
	}
	defer rows.Close()
	foodl := make([]structs.MenyuJson, 0)
	food2 := structs.MenyuJson{}
	for rows.Next() {
		if err = rows.Scan(&food2.Id, &food2.Name, &food2.Price); err != nil {
			//return structs.MenyuJson{}, err
			panic(err)
		}
		foodl = append(foodl, food2)
	}
	return foodl, nil
}

func (u PostgresRepo) Food3() ([]structs.MenyuJson, error) {
	rows, err := u.db.Query(`SELECT id,name,price FROM food WHERE category=2`)
	if err != nil {
		//return structs.MenyuJson{}, err
		panic(err)
	}
	defer rows.Close()
	foodl := make([]structs.MenyuJson, 0)
	food3 := structs.MenyuJson{}
	for rows.Next() {
		if err = rows.Scan(&food3.Id, &food3.Name, &food3.Price); err != nil {
			//return structs.MenyuJson{}, err
			panic(err)
		}
		foodl = append(foodl, food3)
	}
	return foodl, nil
}

func (u PostgresRepo) Salad() ([]structs.MenyuJson, error) {
	rows, err := u.db.Query(`SELECT id,name,price FROM salad`)
	if err != nil {
		//return structs.MenyuJson{}, err
		panic(err)
	}
	defer rows.Close()
	saladl := make([]structs.MenyuJson, 0)
	salad := structs.MenyuJson{}
	for rows.Next() {
		if err = rows.Scan(&salad.Id, &salad.Name, &salad.Price); err != nil {
			//return structs.MenyuJson{}, err
			panic(err)
		}

	}
	return saladl, nil
}

func (u PostgresRepo) Drinks() ([]structs.MenyuJson, error) {
	rows, err := u.db.Query(`SELECT id,name,price FROM drinks`)
	if err != nil {
		//return structs.MenyuJson{}, err
		panic(err)
	}
	defer rows.Close()
	drinkl := make([]structs.MenyuJson, 0)
	drink := structs.MenyuJson{}
	for rows.Next() {
		if err = rows.Scan(&drink.Id, &drink.Name, &drink.Price); err != nil {
			//return structs.MenyuJson{}, err
			panic(err)
		}
	}
	return drinkl, nil
}

func (u PostgresRepo) Shop(table_id, product_id, food_id, salad_id, drinck_id string) error {
	if food_id != "" {
		row1 := u.db.QueryRow("SELECT id FROM food WHERE id=$1", product_id)
		if err := row1.Scan(&food_id); errors.Is(err, sql.ErrNoRows) {
			return errors.New("There is no this type of food")
		} else if err != nil {
			return err
		}
	} else if salad_id != "" {
		row2 := u.db.QueryRow("SELECT id FROM salad WHERE id=$1", product_id)
		if err := row2.Scan(&salad_id); errors.Is(err, sql.ErrNoRows) {
			return errors.New("There is no this type of salad")
		} else if err != nil {
			return err
		}
	} else if drinck_id != "" {
		row3 := u.db.QueryRow("SELECT id FROM drinks WHERE id=$1", product_id)
		if err := row3.Scan(&drinck_id); errors.Is(err, sql.ErrNoRows) {
			return errors.New("There is no this type of drinks")
		} else if err != nil {
			return err
		}
	}
	
	var chek_id string
	row, err := u.db.Query("SELECT id FROM chek WHERE tables_id=$1 AND payment=false", table_id)
	if err != nil {
		return err
	}

	row.Scan(&chek_id)
	_, err = u.db.Exec("INSERT INTO basket VALUES ($1,$2)", chek_id, product_id)
	if err != nil {
		return err
	}
	return nil
}

func (u PostgresRepo) OpenChek(table_id string) error {
	row, err := u.db.Query("SELECT busy FROM tables WHERE table_id=$1", table_id)
	if err != nil {
		return err
	}
	var b bool
	row.Scan(&b)
	if !b {
		_, err = u.db.Exec("INSERT INTO chek VALUES ($1,$2)", uuid.NewString(), table_id)
		if err != nil {
			return err
		}
		//tableni busy sini true qilish
		_, err = u.db.Exec("Update tables SET busy=true WHERE table_id=$1", table_id)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("Table is busy!")
}

func (u PostgresRepo) Chek(table_id string) (map[string]int, error) {
	mp := make(map[string]int)
	var id string
	row := u.db.QueryRow("SELECT id FROM chek WHERE tables_id=$1 AND payment=false", table_id)
	row.Scan(&id)
	row2, err := u.db.Query("SELECT product_id FROM basket WHERE chek_id=$1", id)
	if err != nil {
		return nil, err
	}
	var list []string
	for row2.Next() {
		var s string
		if err = row2.Scan(&s); err != nil {
			return nil, err
		}
		list = append(list, s)
	}
	for _, id := range list {
		var (
			nomi  string
			narxi uint32
		)
		row := u.db.QueryRow("SELECT name,price FROM food WHERE id=$1", id)
		if err = row.Scan(&nomi, &narxi); err != nil {
			row2 := u.db.QueryRow("SELECT name,price FROM salad WHERE id=$1", id)
			if err = row2.Scan(&nomi, &narxi); err != nil {
				row3 := u.db.QueryRow("SELECT name,price FROM drinks WHERE id=$1", id)
				if err = row3.Scan(&nomi, &narxi); err != nil {

				}
			}
		}

	}
	// tableni busysini false VA chekni paymentini true qilish kk
	_ = u.db.QueryRow("Update tables SET busy=false WHERE table_id=$1", table_id)
	_ = u.db.QueryRow("Update chek SET payment=true WHERE table_id=$1", table_id)
	return mp, nil
}
