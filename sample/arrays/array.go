package arrays

import "fmt"

func ArrayDemo(){

	grade1 := 10
	grade2 := 20
	grade3 := 30

	fmt.Println("grade1 :", grade1)
    fmt.Println("grade2 :", grade2)
	fmt.Println("grade3 :", grade3)

	grades := [3]int{grade1, grade2, grade3}
	fmt.Println("Grades in an array", grades)

	Fgrade := grades[2]
	fmt.Println("First grade in array", Fgrade)

	var arraysample [3]string
	fmt.Println("Array Sample is", arraysample)

	arraysample[0] = "Akhila"
	arraysample[1] = "Sada"
	arraysample[2] = "Name"

	arraysample2 := arraysample

	arraysample[2] = "Rose"

	fmt.Println("Array is :", arraysample)
	fmt.Println("Array1 is :", arraysample2)


	fmt.Println("Array Sample is :", arraysample)
	fmt.Println("Array Sample is :", &arraysample[1])
	
}
