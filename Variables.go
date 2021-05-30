package main

import "fmt"

/**

# Variables
  ## declaring  :

	- var variableName type 

  ## declaring & initializing  :

	-  var variableName type = value 

  ## declaring & initializing (shorthand) :

	use it inside main function 
	
	- variableName := value

  ## Multiale Variable Declarations :
	- var var1 , var2 , var3 type
	- var var1 , var2 , var3 = val1 , val2 , val3 
	- var1 , var2 , var3 := val1 , val2 , val3

# Constant : use the keyword (const)
	- const constantName = value

# Grouping variables or constants 

	var(
		var1 = val1
		var2 = val2
	)

	const (
		const1 = val1
		const2 = val2
	)

**/

	// var myFavNum int
	// var myName string

	// var myFavNum int = 1
	// var myName string = "kira"

	// var myFavNum , myscore int
	// var myFavNum , myscore = 10 ,200

	//const pi = 3.14

	// var (
	// 	myFavNum int = 7
	// 	myName string = "kira"
	// )
	
	// const (
	// 	gravity = 10
	// 	pi = 3.14
	// )
func main() {

	myName := "kira"	
	myage := 18

	fmt.Println("Variables")
	fmt.Println(myName)
	fmt.Println(myage)
}