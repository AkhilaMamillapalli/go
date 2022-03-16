package structDemo

import "fmt"

type Movie struct {
	Name string
	YearOfRelease int
	Actors []string
}

func StructDemo(){
	var m1 Movie
	//m1 := new(Movie) it prints &along with content in it//
	m1.Name = "BatMan"
	m1.YearOfRelease = 2008
	m1.Actors = make([]string, 0)

	m1.Actors = append(m1.Actors, "ABC")
	m1.Actors = append(m1.Actors, "CBA")
	fmt.Println("Movie", m1)
	fmt.Println("Movie name:", m1.Name)
	fmt.Println("Movie Actor:", m1.Actors)
	fmt.Println("Movie YearOfRelease:", m1.YearOfRelease)

	slices := make([]int, 0)
	slices = append(slices,6890)
	slices = append(slices, 54345)
	fmt.Println("Slices:",slices)
	employees := make([]Employee, 0)
	

	var emp1 Employee
	emp1.Name = "Akhila"
	emp1.Salary = 298080
	emp1.Position = "CEO"
	
	var emp1Address EmployeeAddress
	emp1Address.HouseNumber = 123
	emp1Address.StateName = "Michigan"
	emp1Address.CityName = "Detroit"
	emp1Address.StreetName = "FarmingtonHills"
	emp1Address.ZipCode = 48335
	emp1.Address = emp1Address
	employees = append(employees, emp1)

	var emp2 Employee
	emp2.Name = "Sada"
	emp2.Salary = 8976987
	emp2.Position = "CEO"

	var emp2Address EmployeeAddress
	emp2Address.HouseNumber = 123
	emp2Address.StateName = "Chicago"
	emp2Address.CityName = "Detroit"
	emp2Address.StreetName = "FarmingtonHills"
	emp2Address.ZipCode = 48335
	emp2.Address = emp2Address
	employees = append(employees, emp2)
	fmt.Println("Employee:",employees)

	//for  i := 0 ; i <= 1 ; i++ {
		//fmt.Println(employees[i])//
	//}Syntax for for loop
    for _, val := range employees{
	fmt.Println(val)
	}//syntax for for range loop in structs
}

type Employee struct {
	Name string
	Salary int
	Position string
	Address EmployeeAddress
}

type EmployeeAddress struct {
	HouseNumber int
	StreetName string
	CityName string
	StateName string
	ZipCode int
	
}
