package samplestring

import "fmt"

func StringDemo(){

	s1 := "This is Sample"
	fmt.Println("String is :", s1)
	s2 := "This Sample"
	fmt.Println("String is :", s2)
	s3 := s1 +"." + s2

	fmt.Println(s3)
	fmt.Println(len(s2))



}