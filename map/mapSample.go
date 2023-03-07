package main

import "fmt"

func main() {
	student := map[string]string{"name1": "sanika", "name2": "jerin"}

	for i := range student {

		fmt.Println(i, student[i])

	}
}
