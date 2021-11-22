//all Go code is contained in packages
//you must include the main package if you want the code to be compiled
package main

//used to import packages into this package to make features available
//fmt comes from the Go standard library
import "fmt"

//the function body contains statements that are the actual instructions to perform program tasks
//every program needs main
//that is the starting point for all programs
func main() {
	fmt.Println("Hello, world!")

	//declaring and then initialising a variable
	var msg string
	msg = "Hello!"

	fmt.Println(msg)

	//variable decalred AND initialised
	newMsg := "Hi there!"

	fmt.Println(newMsg)

	//printing a variable
	fmt.Printf("Yeah!")
}
