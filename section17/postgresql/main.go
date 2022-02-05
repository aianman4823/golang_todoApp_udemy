package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Person struct {
	Name string
	Age  int
}

var Db *sql.DB
var err error

func main() {
	Db, err = sql.Open("postgres", "user=test_user dbname=testdb password=password sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	// C
	// cmd := `insert into persons (name, age) values ($1, $2)`
	// _, err = Db.Exec(cmd, "Nancy", 20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// R
	cmd := `select * from persons where age = $1`
	row := Db.QueryRow(cmd, 1000)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No Row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	cmd1 := `select * from persons`
	rows, _ := Db.Query(cmd1)

	defer rows.Close()

	var pp []Person
	for rows.Next() {
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}

	for _, v := range pp {
		fmt.Println(v.Name, v.Age)
	}

	// U
	cmd2 := `update persons set age = $1 where name = $2`
	_, err := Db.Exec(cmd2, 25, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}

	// D
	cmd3 := `delete from persons where name = $1`
	_, err = Db.Exec(cmd3, "Nancy")
	if err != nil {
		log.Fatalln(err)
	}
}
