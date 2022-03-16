package arrays

import "fmt"

func SampleDemo(){

	Sampleslice := make([]int,0)

	Sampleslice[0] = 10
	Sampleslice[1] = 20
	Sampleslice[2] = 30
	//Sampleslice[3] = 40
	//Sampleslice[4] = 50
	//Sampleslice[5] = 60
	//Sampleslice[6] = 70

	Sampleslice = append(Sampleslice, 80)
	Sampleslice = append(Sampleslice, 90)
	Sampleslice = append(Sampleslice, 100)

	fmt.Println("length :",len(Sampleslice))
	fmt.Println("capacity :",cap(Sampleslice))
	
	//Sampleslice2 := Sampleslice[0:2]

	//Sampleslice3 := Sampleslice[3:]

	//fmt.Println("Sa1 :", Sampleslice3)

	//Sampleslice[5] = 30

	//fmt.Println(Sampleslice3)

	//fmt.Println(Sampleslice)

	//fmt.Println("New Sample is :", Sampleslice2)

	fmt.Println("Slices are :", Sampleslice)
}



