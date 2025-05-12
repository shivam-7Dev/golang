package intermediate

import (
	"bufio"
	"os"
)

func IOINTRO() {
	println("====Readin files=====")
	readingOne()
	println("====writing to file=====")
}

func readingOne() {
	//read from 28-reading-files.txt
	/*
		1. create a source stream
		2. create a reader form that
		3. run a loop and read all the lines
	*/

	source, err := os.OpenFile("intermediate/28-flagsandcode.txt", os.O_RDONLY, 0400)
	if err != nil {
		println("error while opeing the source")
		return
	}
	//create a buffered reader from the source

	reader := bufio.NewReader(source)

	//infinte loop over the reader and print each line

	for {

		line, err := reader.ReadString('\n')
		if err != nil {
			println("error while reading the file")
			return
		}

		println(line)
	}

}
