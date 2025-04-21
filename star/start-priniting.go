package star

import "fmt"

func PrintStar() {
	fmt.Println("this is star priting  funtion-------------")
	fmt.Println()
	solidPatter()
	fmt.Println()
	rightAngleTrianlge()
	fmt.Println()
	invertedRigntAngleTrianlge()
	fmt.Println()
	rightAlignTrianlge()

}

func solidPatter() {

	for i := 0; i < 5; i++ {
		// Outer loop = rows
		for j := 0; j < 5; j++ {
			//inner loop = stars in each row.
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func rightAngleTrianlge() {

	for i := 0; i < 5; i++ {
		// Outer loop = rows
		for j := 0; j <= i; j++ {
			//iNumber of stars increases with each row.
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func invertedRigntAngleTrianlge() {
	for i := 0; i < 5; i++ {
		// Outer loop = rows

		//first loop for start
		//stars= 5-i
		for j := 0; j < 5-i; j++ {
			//iNumber of stars increases with each row.
			fmt.Print("*")
		}
		//second loop for space
		//space=i(row number)
		for j := 0; j < i; j++ {
			//iNumber of stars increases with each row.
			fmt.Print(" ")
		}
		fmt.Println("")
	}
}

func rightAlignTrianlge() {
	for i := 0; i < 5; i++ {

		//print space first

		for j := 0; j < 5-i; j++ {
			fmt.Print(" ")
		}

		// print start

		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}

		fmt.Println()

	}
}
