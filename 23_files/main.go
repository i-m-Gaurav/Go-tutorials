package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	fileInfo, err := f.Stat()

	if err != nil {
		panic(err)
	}

	// we can get the file name.
	fmt.Println("File into: ", fileInfo.Name())
	fmt.Println("File into: ", fileInfo.IsDir())
	fmt.Println("File into: ", fileInfo.Size())
	fmt.Println("File into: ", fileInfo.Mode())
	fmt.Println("File into: ", fileInfo.ModTime())

	// first let's see how we read a file.

	buf := make([]byte, 20)

	d, err := f.Read(buf)

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buf); i++ {
		println("data", d, string(buf[i]))
	}

	println("data", d, buf)

	// this read file, load the whole content in the memory
	// so for the large file, this method is not a good way

	data, err := os.ReadFile("example.txt")

	if err != nil {
		panic(err)
	}

	println(string(data))

	// now lets read the read folders

	dir, err := os.Open(".")

	if err != nil {
		panic(err)
	}

	defer dir.Close()
	// here we will put the number , then it will show that files, in the dir
	// now lets put -1 to get all the files.
	fileDetails, err := dir.ReadDir(-1)

	for _, fi := range fileDetails {
		println(fi.Name())
	}

	//now lets write in the fi,e.

	newFile, err := os.Create("example2.txt")

	if err != nil {
		panic(err)
	}

	defer newFile.Close()

	newFile.WriteString("Hi golang")
	newFile.WriteString("You are good!")

	// it will append these string in the file.

	bytes := []byte("Hello golang how are you")

	newFile.Write(bytes)

	// read and write to another file in the
	// streaming fashion

	sourceFile, err := os.Open("example.txt")

	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("example3.txt")

	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()

		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}

			break
		}

		e := writer.WriteByte(b)

		if e != nil {
			panic(e)
		}

	}

	writer.Flush()

	println("Written to the dest file")


	// lets see how we delete a file

	error := os.Remove("ex.txt")

	if error != nil{
		panic(error)
	}

	println("file deleted Successfully")

}
