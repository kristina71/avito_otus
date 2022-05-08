package main

import (
	"fmt"

	"github.com/vigo/stringutils-demo"
)

func main() {
	const str = "Hello, OTUS!"
	fmt.Println(revertString(str))
}

func revertString(str string) string {
	return stringutils.Reverse(str)
}
