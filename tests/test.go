package main

import (
	"fmt"
	"GOSQL"
)

func main() {
	fmt.Println("Test")
	var Database = GOSQL.Database{HOST: "127.0.0.1", PORT: 3306, DATABASE: "Test123", PASSWORD: "", USERNAME: "root"}
	var Dogs = GOSQL.Table{NAME: "Dogs", COLUMNS: []GOSQL.Column{GOSQL.Column{NAME: "ID", TYPE: "INT", PRIMARY: true}, GOSQL.Column{NAME: "Name", TYPE: "VARCHAR", LENGTH: 100}}}
	Database.AttachTable(Dogs)
}
