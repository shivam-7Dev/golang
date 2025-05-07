package intermediate

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func IO() {

	fmt.Println("this is io package")

	fmt.Println("+++++Read from reader+++++++")
	readFromReader(strings.NewReader("hello world this is shivam sharma"))

	fmt.Println("+++++write to writer+++++++")
	//create a writeer
	var writer bytes.Buffer
	writeToWriter(&writer, "hello writer")
	fmt.Println("writer buffer is---", writer.String())

	fmt.Println("+++++buffer example+++++++")
	bufferExample()

	fmt.Println("+++++mulit reader example+++++++")
	multiReaderExample()

	fmt.Println("+++++Pipe example example+++++++")
	pipeExample()

	fmt.Println("+++++writing data to file+++++++")
	writeTOfile("writeTofile.txt", "data writter to file form the fucntion")
}

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		fmt.Println("error reading from buffer")
		return
	}
	fmt.Println("n bytes:", n)

	//convert the byte slice into string

	fmt.Println("bytes in string ", string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {

	_, err := w.Write([]byte(data))

	if err != nil {
		fmt.Println("error writing from buffer")
		return
	}

}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		fmt.Println("error closing from buffer")
		return
	}

}

func bufferExample() {
	var buf bytes.Buffer // value type=> memory on stack
	buf.WriteString("Hello buffer")
	fmt.Println(buf.String())
}

func multiReaderExample() {

	r1 := strings.NewReader("hello")

	r2 := strings.NewReader("world")

	mr := io.MultiReader(r1, r2)

	buf := new(bytes.Buffer) //pointer type => allocate memory on heap

	_, err := buf.ReadFrom(mr)
	if err != nil {
		fmt.Println("error reader  from mr")
		return
	}

	fmt.Println(buf.String())

}

func pipeExample() {
	fmt.Println("one")
	pr, pw := io.Pipe()

	//create a go routine
	//go routenes need to be executed immediately
	// they are anonlymous functions which care called immidiatly
	// they are like iffe
	//if we add go keyword before a function it becomes a go routine
	//go routines are functions which are anonmyous( no name) and they are ivoked immiedately
	// this anonyoum fucntion is extracted out of the main tread
	// this fuction is executed on different thread
	//if this function need 30 minutes to executed then next fucntion wont wait for 30 min
	// and the execution will fall on next line
	go func() {
		fmt.Println("two")
		//go routine to wirte the data

		pw.Write([]byte("hello pipe"))
		//writeer need to be closed when its done
		pw.Close()
	}()
	fmt.Println("three")
	buf := new(bytes.Buffer)
	fmt.Println("four")
	buf.ReadFrom(pr)
	fmt.Println("five")
	fmt.Println(buf.String())

}

func writeTOfile(pathtofile string, data string) {
	//open the file
	file, err := os.OpenFile(pathtofile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opein/createing file")
		return
	}

	defer closeResource(file)

	//write to the file

	_, err = file.Write([]byte("hello world "))
	if err != nil {
		fmt.Println("Error opein/createing file")
		return
	}

}
