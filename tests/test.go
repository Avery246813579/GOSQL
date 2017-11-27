package main

import (
	"fmt"
	"GOSQL"
)

func main() {
	fmt.Println("Test")
	var Database = GOSQL.Database{HOST: "127.0.0.1", PORT: 3306, DATABASE:"Test123", PASSWORD:"", USERNAME:"root"}
	Database.Init()
}
