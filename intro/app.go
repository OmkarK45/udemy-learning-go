package main // entry point to the go app

import "fmt"

/*
The main function is important.
Go will call and execute this when program starts.
It is reserved.

For a specific package, only one main function is allowed
across multiple files.
*/
func main() {
	fmt.Println("Hello World")
}
