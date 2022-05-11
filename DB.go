package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//var db *sql.DB
var schema = `
CREATE TABLE IF NOT EXISTS policy( 
id INTEGER  PRIMARY KEY AUTOINCREMENT,
isbn varchar(20) ,
title varchar(30)
);
`

func createDatabase() *sql.DB {
	db, _ := sql.Open("sqlite3", "./movies.db")
	statement, err := db.Prepare(schema)

	if err != nil {
		log.Println("failed to create Database", err)
	}
	statement.Exec()
	return db

}

func insertDatabase(m Movie, db *sql.DB) {

	statement, err := db.Prepare("insert into policy (isbn,title) values (?,?)")
	if err != nil {
		log.Println("Error in insertion :=>", err)

	}
	//fmt.Println("ID", m.ID)
	//fmt.Println("isbn", m.Isbn)
	//fmt.Println("title", m.Title)
	statement.Exec(m.Isbn, m.Title)
}

func Readdatabase(db *sql.DB) []Movie {
	fmt.Println("INSIDE READ :=>")
	movies := []Movie{}
	statement, err := db.Query("Select * from policy")
	if err != nil {
		log.Fatal(err)
	}
	var id int
	var isbn string
	var title string
	for statement.Next() {
		statement.Scan(&id, &isbn, &title)
		movie := Movie{ID: id, Isbn: isbn, Title: title}
		movies = append(movies, movie)
	}
	return movies
}
func DeleteDB(id string, db *sql.DB) {
	statement, err := db.Prepare("Delete from policy where id = ?")
	if err != nil {
		log.Println("Error in Deleting the database", err)
	}
	statement.Exec(id)
}

func UpdateDB(m Movie, id string, db *sql.DB) {
	statement, err := db.Prepare("Update policy set isbn=?,title=?  where id=?")
	if err != nil {
		log.Println("Error in Updation :", err)
	} else {
		fmt.Println("successfully Updated :=>")
	}
	statement.Exec(m.Isbn, m.Title, id)

}
