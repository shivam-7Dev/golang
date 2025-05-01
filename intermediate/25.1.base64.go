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
	urlSafeExample()
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

func urlSafeExample() {
	fmt.Println("---URL-Safe Encoding Example---")

	// Original data with characters that would cause URL issues
	input := "Hello+World/with?special=chars"

	// Standard encoding (may cause issues in URLs)
	stdEncoded := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println("Standard encoded:", stdEncoded)

	// URL-safe encoding (safe for use in URLs)
	urlEncoded := base64.URLEncoding.EncodeToString([]byte(input))
	fmt.Println("URL-safe encoded:", urlEncoded)

	// Decode the URL-safe string
	decoded, _ := base64.URLEncoding.DecodeString(urlEncoded)
	fmt.Println("Decoded:", string(decoded))
}
