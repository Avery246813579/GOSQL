package GOSQL

import (
	"fmt"
)

type Table struct {
	NAME 		string
	DATABASE	Database
	COLUMNS 	[]Column
}

func (h Table) GetTable(column Column) {
	fmt.Println("INITED" + column.NAME)
}
