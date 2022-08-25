package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("start app")
	db, err := sql.Open("mysql", "root:@/goclass")
	if err != nil {
		panic(err)
	}
	fmt.Println(db)
	InsertToDB(db)
	UpdateDb(db)
	Delete(db)
}

func Delete(db *sql.DB) {
	statement:= "DELETE FROM users WHERE id=4"
	_, err := db.Exec(statement)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UpdateDb(db *sql.DB) {
	statement := "UPDATE users SET first_name='bahram', mobile='0930257894' WHERE id=36 "
	_, err := db.Exec(statement)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func InsertToDB(db *sql.DB) {
	sqlStatement := `
INSERT INTO users (first_name, last_name ,mobile ,registration_date , national_code)
VALUES (?, ?, ?, ? ,?)`
	_, err := db.Exec(sqlStatement, "arash", "rahimi", "09120302", time.Now(), "025478")
	if err != nil {
		fmt.Println(err)
		return
	}
}