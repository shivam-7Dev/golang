package intermediate

import (
	"errors"
	"fmt"
)

/*
✅ YES — you must check err != nil after each function call where an error could occur.
Because in Go, error checking is explicit.
Go's design philosophy is "explicit is better than implicit"

	func mainFunction() error {
		err := doFirstThing()
		if err != nil {
			return err
		}

		err = doSecondThing()
		if err != nil {
			return err
		}

		err = doThirdThing()
		if err != nil {
			return err
		}

		// ... and so on
	}

😵‍💫 But isn't it repetitive?
Yes, it can feel repetitive.
That's why Go developers use certain patterns to make it less ugly.


	if err := doFirstThing(); err != nil {
		return err
	}

	if err := doSecondThing(); err != nil {
		return err
	}


*/

/*
Custom error
**/

type myError struct {
	message string
	code    int
}

func (m *myError) Error() string {
	//Sprintf formats according to a format specifier and returns the resulting string.
	return fmt.Sprintf("Error:%s", m.message, m.code)
}

func Errors() {

	isEven, error := isEven(2)
	if error != nil {
		fmt.Println("Error in Funciton Erros:", error)
		return
	}
	fmt.Println("number is even", isEven)

	// if err := customErrorTest(); err != nil {
	// 	fmt.Println("someting went worng", err)
	// 	// Type assert to access the custom error fields
	// 	if myErr, ok := err.(*myError); ok {
	// 		fmt.Println("Error code:", myErr.code)
	// 		fmt.Println("Error message:", myErr.message)
	// 	} else {
	// 		fmt.Println("Different type of error:", err)
	// 	}
	// }

	err2 := readData()
	if err2 != nil {
		fmt.Println("error:", err2)
		return
	}

}

func isEven(num int) (bool, error) {
	if num%2 == 0 {
		return true, nil
	}
	// return false, fmt.Errorf("number is odd")
	return false, errors.New("number is odd")
}

func customErrorTest() error {
	return &myError{
		message: "unauthorized",
		code:    401,
	}
}

func readData() error {
	err := readConfig()
	if err != nil {
		return fmt.Errorf("readData error: %w", err)
	}
	return nil
}

func readConfig() error {
	return errors.New("Config error")
}
