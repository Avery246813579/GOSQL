package GOSQL

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	USERNAME string
	PASSWORD string
	DATABASE string
	HOST     string
	PORT     int
}

func (d Database) Init(){
	db, err := sql.Open("mysql",
		d.USERNAME + ":" + d.PASSWORD + "@tcp(" + d.HOST + ":" + string(d.PORT) + ")/" + d.DATABASE)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func (d Database) AttachTable(table Table) {
	table.DATABASE = d


}
