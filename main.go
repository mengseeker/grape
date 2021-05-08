package main

import (
	"fmt"
	"grape/pkg/arrays"
	"strings"
)

func main() {
	clumn := "abc"
	sval := []string{"a", "b", "c"}
	ival := []int{1, 2, 3}
	fmt.Printf("%s in (\"%s\")\n", clumn, strings.Join(sval, "\", \""))
	fmt.Printf("%s in (%s)\n", clumn, arrays.JoinInt(ival, ","))

	var m map[string]interface{}

	for _, v := range m {
		fmt.Print(v)
	}
}
