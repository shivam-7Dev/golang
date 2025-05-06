package intermediate

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func BuffioInto() {
	fmt.Println("this is buffer io package")
	// basicStringReader()
	// basicStingReader2()
	// basicbufiowirter()
	// basicWriter()
	// souncedestination()

	err := copyFileBuffered("24-bufio-package.txt", "destination.txt")
	if err != nil {
		fmt.Println("Error copying file:", err)
	} else {
		fmt.Println("File copied successfully.")
	}
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

func souncedestination() {
	//open source file in readonly mode
	source, err := os.Open("24-bufio-package.txt")
	if err != nil {
		fmt.Println("error opeingn source file")
		return
	}
	defer source.Close()

	//open destination file in wirte only mode
	destination, err := os.OpenFile("24.4.copy.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("error opeingn destination file")
		return
	}
	defer destination.Close()
	//create bufio.reder
	reader := bufio.NewReader(source)
	//create bufio.writer
	writer := bufio.NewWriter(destination)
	// use for loop
	for {
		// read bytes or line for source

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading line")
			return
		}
		//write bytes or line to destination

		writer.WriteString(line)

	}

}

func copyFileBuffered(srcPath, dstPath string) error {
	// Step 1: Open source file
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open source: %w", err)
	}
	defer srcFile.Close()

	// Step 2: Create destination file
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return fmt.Errorf("failed to create destination: %w", err)
	}
	defer dstFile.Close()

	// Step 3 & 4: Create buffered reader and writer
	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)

	// Step 5: Create a buffer and loop
	buffer := make([]byte, 1024*1024) // 1 MB buffer
	for {
		n, err := reader.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("read error: %w", err)
		}
		if n == 0 {
			break // End of file
		}

		if _, err := writer.Write(buffer[:n]); err != nil {
			return fmt.Errorf("write error: %w", err)
		}
	}

	// Step 6: Flush writer
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("flush error: %w", err)
	}

	return nil
}
