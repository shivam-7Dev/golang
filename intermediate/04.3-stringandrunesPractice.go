package intermediate

import "fmt"

func StringsPractice() {
	// one()
	stringIteration3()
}

func one() {
	var name string = "shivam"
	middleName := "kumar"
	lastName := `Sharma`

	fmt.Println(name)
	fmt.Println(middleName)
	fmt.Println(lastName)

	fullName := name + middleName + lastName

	fmt.Println(fullName)
}

func stringIteration3() {
	fmt.Println("\n--- String Iteration ---")
	s := "hello 世界"
	fmt.Println("Iterating by bytes:")
	for i := 0; i < len(s); i++ {
		fmt.Printf("Byte %d: %x\n", i, s[i])
	}

	fmt.Println("Iterating by runes:")
	for i, r := range s {
		fmt.Printf("Rune %d: %c (Unicode: %U, Raw Value: %d)\n", i, r, r, r)
	}
}
