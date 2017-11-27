package GOSQL

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func mysqlTest() {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/Test123")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		ID int
		NAME string
		SEX string
	)
	rows, err := db.Query("select ID, NAME, SEX from John")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ID, &NAME, &SEX)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(ID, NAME, SEX)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
