package GOSQL

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strconv"
)

type Database struct {
	USERNAME   string
	PASSWORD   string
	DATABASE   string
	HOST       string
	PORT       int
}

func (d Database) AttachTable(table Table) {
	table.DATABASE = d

	db, err := sql.Open("mysql",
		d.USERNAME+":"+d.PASSWORD+"@tcp("+d.HOST+":"+strconv.Itoa(d.PORT)+")/"+d.DATABASE)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var statement = "CREATE TABLE " + table.NAME + "("
	var elements = ""
	var post = ""
	for _, element := range table.COLUMNS {
		elements += ", " + element.NAME + " " + element.TYPE

		if element.LENGTH != 0 {
			elements += "(" + strconv.Itoa(element.LENGTH) + ")"
		}

		if element.PRIMARY{
			post += ", PRIMARY KEY(" + element.NAME + ")"
		}
	}

	fmt.Println(statement + elements[2:] + post + ")")
	rows, err := db.Query(statement + elements[2:] + post + ")")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

}
