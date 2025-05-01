package intermediate

import (
	"encoding/base64"
	"fmt"
)

func Base64Intro() {

	fmt.Println("this is base 64 intro")
	binaryToBase64()
	decondingBase64tostring()
	encodeDecodeBytes()
}

func binaryToBase64() {
	fmt.Println("---binaryToBase64------")
	input := "Hello Shivam"
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println("Base64 Encoded:", encoded)
}

func decondingBase64tostring() {
	fmt.Println("---decondingBase64tostring------")
	input := "Hello Shivam"
	//encode the string
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	//decode the string
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Decoded:", string(decodedBytes))

}

func encodeDecodeBytes() {
	fmt.Println("---encodeDecodeBytes------")
	data := []byte{1, 2, 3, 4}
	encoded := base64.StdEncoding.EncodeToString(data)
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println("Original:", data)
	fmt.Println("Encoded :", encoded)
	fmt.Println("Decoded :", decoded)

}
