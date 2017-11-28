package GOSQL

import (
	"fmt"
)

func dictToString(element map[string]string) {
	//var toReturn string = ""

	for _, key := range element {
		fmt.Println(key)
	}
}
