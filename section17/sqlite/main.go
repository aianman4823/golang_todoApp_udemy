package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	// Db, _ := sql.Open("sqlite3", "./example.sql")
	// defer Db.Close()

	// // テーブルの作成
	// cmd := `CREATE TABLE IF NOT EXISTS persons(
	// 	name string,
	// 	age int)`

	// _, err := Db.Exec(cmd)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// データの追加
	// Db, _ := sql.Open("sqlite3", "./example.sql")
	// defer Db.Close()

	// cmd2 := "INSERT INTO persons (name, age) VALUES (?, ?)"
	// _, err := Db.Exec(cmd2, "tarou", 20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// データの更新
	// Db, _ := sql.Open("sqlite3", "./example.sql")
	// defer Db.Close()

	// cmd2 := "update persons set age = ? where name = ?"
	// _, err := Db.Exec(cmd2, 25, "tarou")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Db, _ := sql.Open("sqlite3", "./example.sql")
	// defer Db.Close()

	// cmd := "INSERT INTO persons (name, age) VALUES (?, ?)"
	// _, err := Db.Exec(cmd, "hanako", 19)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// データの取得
	// Db, _ := sql.Open("sqlite3", "./example.sql")
	// defer Db.Close()

	// cmd := "select * from persons where age=?"
	// // QueryRow 1レコードを取得できる
	// row := Db.QueryRow(cmd, 25)
	// var p Person
	// err := row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("No row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// 複数のデータの取得
	// Db, _ := sql.Open("sqlite3", "./example.sql")
	// defer Db.Close()

	// cmd := "select * from persons"
	// // Query 複数レコードを取得できる
	// rows, _ := Db.Query(cmd)
	// defer rows.Close()

	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	pp = append(pp, p)
	// }

	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// データの削除
	Db, _ := sql.Open("sqlite3", "./example.sql")
	defer Db.Close()

	cmd := "delete from persons where name = ?"
	_, err := Db.Exec(cmd, "hanako")
	if err != nil {
		log.Fatalln(err)
	}
}
