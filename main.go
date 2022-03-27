package main

import "fmt"

func main() {
	fmt.Println("Hello Pavel")
	fmt.Println("Second Line")

	var age int
	fmt.Println("My age is:", age)

	var uid = 12345
	fmt.Println("My uid", uid)
	fmt.Println("%t\n", uid)

	var (
		personName string = "Bob"
		personAge         = 42
		personUID  int
	)
	fmt.Printf("Name: %s\n Age %d\n UID: %d\n", personName, personAge, personUID)

}
