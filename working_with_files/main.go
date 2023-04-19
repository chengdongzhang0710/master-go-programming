package main

func main() {
	// newFile, err := os.Create("a.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// newFile.Close()
	// fmt.Printf("var newFile is of type %T\n", newFile)
	// file, err := os.OpenFile("a.txt", os.O_APPEND, 0644)
	// // file, err := os.Open("a.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// file.Close()
	// fmt.Printf("%#v, %#v\n", newFile, file)
	// fmt.Println("var newFile and var file is equal:", newFile == file)

	// byteSlice := []byte("I learn golang!")
	// fmt.Printf("%#v\n%s\n", byteSlice, byteSlice)
	// fmt.Println(byteSlice)

	// byteSlice := []byte("I learn Golang! 传")
	// bytesWritten, err := file.Write(byteSlice)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Bytes written: %d\n", bytesWritten)

	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// text := scanner.Text()
	// bytes := scanner.Bytes()
	// fmt.Println("Input text:", text)
	// fmt.Println("Input bytes:", bytes)

	// newFile, err := os.Create("info.txt")
	// newFile.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fileInfo, err := os.Stat("info.txt")
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		log.Fatal("The file is not exist!")
	// 	} else {
	// 		log.Fatal(err)
	// 	}
	// }
	// _ = fileInfo
	// err = os.Rename("info.txt", "data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = os.Remove("data.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// file, err := os.Open("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// data, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", data)

	// data, err := ioutil.ReadFile("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", data)

	// file, err := os.Open("info.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	// if err = scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// file, err := os.OpenFile("info.txt", os.O_WRONLY|os.O_CREATE, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// bs := []byte("The Go gopher is an iconic mascot!")
	// _, err = file.Write(bs)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bs := []byte("The Go gopher is an iconic mascot!")
	// err := ioutil.WriteFile("info.txt", bs, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
