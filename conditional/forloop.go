package conditional

import "fmt"

func ForDemo() {

	//Loop Initialization,loop conditional,loop incerement//
	//for i := 0; i < 10; i = i + 1 {
		//fmt.Println(i)

	//}

	Gryffindor := make([]int, 0, 40)

	for i := 1; i < 40; i++ { //(i = i + 1)

		Gryffindor = append(Gryffindor, i)
		}
		fmt.Println("Gryffindor" ,Gryffindor)

	for i := 1; i <= 40; i = i + 1 {

		fmt.Println(Gryffindor[i])
	}

}
