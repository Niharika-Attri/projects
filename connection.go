package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type student struct {
	Username string    `db:"username"`// first letter capitalize: global property
	Name     string    `db:"name"`
	Email    string    `db:"email"`
	Dob      time.Time `db:"dob"`
	Password string    `db:"password"`
}

func main() {
	var db *sqlx.DB
	var err error

	// connect to postgresql database
	db, err = sqlx.Connect("postgres", "user=postgres dbname=Students sslmode=disable password=Na123@2005 host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	// test connection to database
	if err = db.Ping(); err != nil { // ping: verifies if connection is still alive, establishing connection if necessary, uses context.Background internally
		log.Fatal(err)
	} else {
		log.Println("successfully connected")
	}

	// SELECT
	place := student{} //initialise a student struct to store data

	rows, _ := db.Queryx("SELECT username, name, email, dob, password FROM std_data") // queries the database and returns an *sqlx.Rows

	for rows.Next() { //prepares the next row for scan method, returns true on success, false if an error occured or there is no next row
		err := rows.StructScan(&place) //scan current row into 'place' variable
		if err != nil {
			log.Fatalln("error scanning:", err)
		}
		log.Printf("%#v\n", place) //log the content of place struct for each row
	}

	// INSERT (hardcoded)
	// insertN := `insert into "std_data"(username, name, email, dob, password) values('annaH', 'Anna Heath', 'annaheath@gmail.com', '12-08-1999', 'annaheath')`
	// _, e := db.Exec(insertN)
	// if err != nil{
	// 	panic(e)
	// }

	//INSERT (dynamically)
	// insertDyn := `insert into "std_data" ("username", "name", "email", "dob", "password") values ($1, $2, $3, $4, $5)`
	// _, e := db.Exec(insertDyn, "ZoeArch", "Zoe Archer", "zoearcher12@gmail.com", "09-16-1989", "zoearcher12")
	// fmt.Println(e)
	// if err != nil {
	// 	panic(e)
	// }

	// UPDATE
	updateN := `update std_data set username = 'annaheath' where name = 'Anna Heath'`
	_, e = db.Exec(updateN)
	fmt.Println(e)
	if err != nil{
		panic(e)
	}
}
