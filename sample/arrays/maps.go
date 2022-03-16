package arrays

import "fmt"

func MapsDemo() {

	mapsofPopulation := make(map[string]int)

	mapsofPopulation["Dallas"] = 248074
	mapsofPopulation["Texas"] = 32689
	mapsofPopulation["Austin"] = 589568
	mapsofPopulation["Charlotte"] = 0

	fmt.Println(mapsofPopulation)
	fmt.Println("Texas Population", mapsofPopulation["Texas"])
	fmt.Println("Length", len(mapsofPopulation))
	mapsofPopulation["Dallas"] = 127823
	fmt.Println(mapsofPopulation)

	mapsofPopulation["dallas"] = 48659
	fmt.Println(mapsofPopulation)
	fmt.Println(mapsofPopulation["Boston"])

	fmt.Println(mapsofPopulation["Charlotte"])

	if _, ok := mapsofPopulation["Boston"]; !ok {
		fmt.Println("Boston Key is present:", ok)
	}

	if _, ok := mapsofPopulation["Charlotte"]; ok {
		fmt.Println("Charlotte Key is present:", ok)

	}

}