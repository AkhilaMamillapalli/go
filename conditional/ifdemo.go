package conditional

import "fmt"

func IfDemo(){

//	batsmanScore := 101
//	batsmanScore = 99

//	if batsmanScore >= 100 {

//	fmt.Println("Batsman score is Century")

//	} else {

//	fmt.Println("Batsman score is not Century")

//	}
//	fmt.Println("Batsman has some score")

//	number := 89
//	booleanval := number == 89
//	fmt.Println(booleanval) 
//	if number%2 == 0 {
//		fmt.Println("Statement 1")
//	}else {
//		fmt.Println("Statement 2")
//	}
//	fmt.Println("Statement 3")
//	if true {
//		fmt.Println("Print True")

//	}else {
//		fmt.Println("Print False")
//	}
a := 100
if a >= 100{
	fmt.Println("A is century score")

}else if a <= 200{
fmt.Println("A is double century")
}else {
	fmt.Println("A is not a century")
}
b := 10
if b < 100 {
	if b > 50 {
		fmt.Println("Statement 1")
	} else {
		fmt.Println("Statement 2")

	}
}	
}

	


