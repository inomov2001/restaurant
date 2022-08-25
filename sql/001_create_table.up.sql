create table products (
                          id UUID not null primary key,
                          name varchar(64) not null UNIQUE,
                          created_at timestamp not null default current_timestamp
);

create table food (
    id UUID not null primary key,
    name varchar(80) not null UNIQUE,
    price int not null,
    category int not null,
    products_id UUID[] not null
);

create table salad (
  id UUID not null primary key,
  name varchar(80) not null UNIQUE,
  price int not null
);

create table drinks (
                        id UUID not null primary key,
                        name varchar(80) not null UNIQUE,
                        price int not null
);
create table basket(
                       chek_id UUID not null references chek(id),
                       product_id UUID not null references products(id)
---UPDATE Customers
---SET ContactName = 'Alfred Schmidt', City= 'Frankfurt'
---WHERE CustomerID = 1;
);

create table tables (
    id UUID not null primary key,
    number int not null,
    busy bool not null default false
);

create table chek (
    id UUID not null primary key,
    tables_id UUID not null references tables(id),
    payment bool not null default false,
    exit_time TIMESTAMP not null
);


--- func main() {
--- 	db := connectDB()
--- 	defer db.Close()

--- 	rows, err := db.Query(
--- 		"SELECT * FROM students",
--- 	)
--- 	if err != nil {
--- 		log.Panicf("failed to query: %v", err)
--- 	}

--- 	defer rows.Close()
--- 	students := make([]Student, 0)
--- 	for rows.Next() {
--- 		var s Student
--- 		if err = rows.Scan(&s.Id, &s.Name, &s.Score); err != nil {
--- 			log.Panicf("failed to scan: %v", err)
--- 		}
--- 		students = append(students, s)
--- 	}

--- 	fmt.Println(students)
--- }
