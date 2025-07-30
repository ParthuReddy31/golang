package main

import "fmt"

const LoginToken = "hgugusdf" //public

func main() {
	//variable
	var username string = "Parthu"
	fmt.Println(username)
	fmt.Printf("Variable type is: %T \n", username)
	fmt.Println("\n--------------")
	// Boolian
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type is: %T \n", isLoggedIn)
	fmt.Println("\n--------------")
	//unit
	var smallVal uint8 = 255 
	fmt.Println(smallVal)
	fmt.Printf("Variable type is: %T \n", smallVal)
	fmt.Println("\n--------------")
	//float
	var smallFloat float64 = 255.3245786536598765434567
	fmt.Println(smallFloat)
	fmt.Printf("Variable type is: %T \n", smallFloat)
	fmt.Println("\n--------------")

	//defaul values and some aliases
	var anotherVariable int 
	fmt.Println(anotherVariable)
	fmt.Printf("Variable type is: %T \n", anotherVariable)
	fmt.Println("\n--------------")

	// no mentioning variable  type

	var number = 23
	fmt.Println(number)
	fmt.Printf("Variable type is: %T \n", number)
	fmt.Println("\n--------------")

	number1 := "parthu"
	fmt.Println(number1)
	fmt.Printf("Variable type is: %T \n", number1)
	fmt.Println("\n--------------")

	fmt.Println(LoginToken)
	fmt.Printf("Variable type is: %T \n", LoginToken)
	fmt.Println("--------------")
}
