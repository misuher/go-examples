package database

import (
	"database/sql"
	"fmt"
	"log"

	//import the mysql driver with a preceding _ to invoke its init() function
	_ "github.com/go-sql-driver/mysql"
)

func insert(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO test(id, username, password) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(1456, "name", "pass")
	if err != nil {
		log.Fatal(err)
	}
}

func update(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("UPDATE test SET username = ? where id = ?")
	_, err = stmt.Exec("newname", 1456)
	if err != nil {
		log.Fatal(err)
	}
}

func delete(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare("DELETE FROM test where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(1457)
	if err != nil {
		log.Fatal(err)
	}
}

func query(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT password FROM test WHERE username = ?", "username")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var userpass string
	for rows.Next() { //while there is more results
		err := rows.Scan(&userpass) //get the value
		if err != nil {
			panic(err)
		}
		fmt.Println(userpass)
	}
}

func main() {
	//Open db connection
	db, err := sql.Open("mysql", "user:password@tcp(ip:port)/slug?charset=utf8") //db type is *sql.DB
	if err != nil {
		panic(err)
	}
	defer db.Close()
	insert(db)
	update(db)
	delete(db)
	query(db)
}
