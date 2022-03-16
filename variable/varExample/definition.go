package variable

import "fmt"

var Pkgvariable int = 100
var SmartPhone int = 150

func VarDeclaration(){

	var x int 
	x = 10
	fmt.Println(x)

	var y string = "Sage It"
	fmt.Println("y :", y)


	z := true
	fmt.Println("z :", z)

	
}

func Samplefunc(){
	var x int
	x = 40
	fmt.Println(x)

	fmt.Println("pkgvariable :",Pkgvariable)
}

