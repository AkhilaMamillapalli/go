package conditional

import "fmt"

func SwitchDemo(){

	company := "Sage It"

	switch company{
	case "Sage It" :
		fmt.Println("Sage It is the Company Name")

	case "Google" :
		fmt.Println("Google is the Company Name")

	default :
		fmt.Println("Name of company is not Sage It or Google")
	}

	Abc := 100
	booleanval := Abc == 100
	switch booleanval{
	case Abc > 100 :
	 fmt.Println("Number is greater than 100")
	case Abc == 100 :
		fmt.Println("Number equal to 100")
	case Abc < 100 :
		fmt.Println("Number is less than 100")
	}

}