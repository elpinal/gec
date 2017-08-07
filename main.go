package main

import "fmt"

func main() {
	expr, err := parse([]byte("12"))
	fmt.Println(expr, err)
}
