package main

import "fmt"

func main() {
	err := parse([]byte("12"))
	fmt.Println(err)
}
