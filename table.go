package GOSQL

import (
)

type Table struct {
	NAME     string
	DATABASE Database
	COLUMNS  []Column
}

func (table Table) find (where map[string]string) {

}
