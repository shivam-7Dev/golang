package intermediate

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BuffioInto() {
	fmt.Println("this is buffer io package")
	// basicStringReader()
	// basicStingReader2()
	// basicbufiowirter()
	basicWriter()
}

func basicStringReader() {
	/*
		Read() reads data upto certain limit(few bytes) and we define those limts
		Read() also stores that read data in the given input.
		data := make([]byte, 2)
		n, err := reader.Read(data)
		it returns number of bytes read and error

		ReadString() function reads data updo the given delimter string
		like \n or \t and it return the read string and err if any
		line, _ := reader.ReadString('\n')
		here we get string as an output
	*/
	// var stringReader *strings.Reader

	stringReader := strings.NewReader("Hello! my\n name is \nshivam sharma")
	data := make([]byte, 2)
	reader := bufio.NewReader(stringReader)
	line, _ := reader.ReadString('\n')
	fmt.Println("line:", line)
	for {
		n, err := reader.Read(data)
		if err != nil {
			fmt.Println("error")
			return
		}
		fmt.Printf("nunber of bytes %d and charater %s\n", n, data[:n])
	}
}

func basicStingReader2() {
	stringReader := strings.NewReader("Hello! my \n name is \n shivam sharma")

	reader := bufio.NewReader(stringReader)

	//methods1:read line by line
	/*
		ReadString() function reads data updo the given delimter string
		like \n or \t and it return the read string and err if any
		line, _ := reader.ReadString('\n')
	*/

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line) // Use Print instead of Println to preserve formatting
		if err != nil {
			// EOF (End of File) is expected when we reach the end
			break
		}
	}

}

func basicbufiowirter() {
	file, err := os.Open("24.3.txt")
	if err != nil {
		fmt.Println("error while opeing file", err)
		return
	}
	reader := bufio.NewReader(file)

	newfile, err := os.Create("fist.txt")
	writer := bufio.NewWriter(newfile)
	if err != nil {
		fmt.Println("error while creatin file")
		return
	}

	for {

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error while reading file")
			return
		}
		writer.WriteString(line)
	}
}

func basicWriter() {

	//data to be written on the standard output
	data := []byte("hello,buffio\n")

	writer := bufio.NewWriter(os.Stdout)
	//writing byte slice first\
	n, err := writer.Write(data)
	fmt.Println("number of bytes written ", n)

	if err != nil {
		fmt.Println("error while wirting file", err)
		return
	}
	//fush the buffer so that all the data is flused to standard output
	writer.Flush()

	//writing a stinrg
	str := "hello world! my name is shivam sharma\n"

	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("erro while writting stintg", err)
	}
	fmt.Println("suscessfully written on the standard output", n)
	writer.Flush()
}
