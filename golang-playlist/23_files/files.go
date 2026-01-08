package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("golang-playlist/23_files/example.txt")
	// //f is a pointer to os.File
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }
	// fmt.Println(f)

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	// log the error
	// 	panic(err)
	// }


	// fmt.Println("file name:", fileInfo.Name())
	// fmt.Println("file or folder:", fileInfo.IsDir())
	// fmt.Println("file size:", fileInfo.Size())
	// fmt.Println("file permission:", fileInfo.Mode())
	// fmt.Println("file modified at:", fileInfo.ModTime())





	// read file
	// f, err := os.Open("golang-playlist/23_files/example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close() //run when the main function ends

	// buf := make([]byte, 12)
	// d, err := f.Read(buf) // reads up to len(buf) bytes , d is number of bytes read
	// if err != nil {
	// 	panic(err) //for EOF error
	// }

	// for i := 0; i < len(buf); i++ {
	// 	fmt.Println("data", d, string(buf[i]),buf[i])
	// }
	// fmt.Println(buf)
	// fmt.Println(string(buf))

	// data, err := os.ReadFile("golang-playlist/23_files/example.txt") //data is []byte
	// //ReadFile reads the entire file content and loads it into memory
	// //so use it only for small files
	// //else use streaming fashion using os.Open and then read using bufio or file.Read
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("data:",data)
	// fmt.Printf("Type: %T\n", data)  // Output: Type: []byte

	// fmt.Println(string(data))





	// read folders

	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)//if n>0 reads n entries else all entries
	// if  err != nil {
	// 	panic(err)
	// }

	// // as there can be multiple files so we use for loop ( as we reading a Directory)
	// for _, fi := range fileInfo {//_ is index, fi is value
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }





	// create a file
	// f, err := os.Create("golang-playlist/23_files/example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("hi go")
	// f.WriteString("nice language")
	// bytes := []byte("Hello Golang")
	// f.Write(bytes)
	// fmt.Println("file created successfully")





	// read and write to another file (streaming fashion)

	sourceFile, err := os.Open("golang-playlist/23_files/example.txt")
	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("golang-playlist/23_files/example2.txt")
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)
	writer := bufio.NewWriter(destFile)

	for {//read byte by byte and write to new file
		b, err := reader.ReadByte()
		//b is a single byte read from the file
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}

			break //in case of EOF we break the loop
		}

		e := writer.WriteByte(b)
		if e != nil {
			panic(e)
		}

	}

	writer.Flush() //flush the writer buffer to the file
	//use writer.Flush() to ensure all buffered data is written to the underlying file

	fmt.Println("written to new file succesfully")




	

	// delete a file

	// sourceFile, err := os.Open("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// err := os.Remove("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(" file deleted successfully")

}
