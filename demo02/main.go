package main

func main() {

	// 读取文件的三种方式
	// 1. Read方式   os.Open()
	// file, err := os.Open("D:/gotest/test.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// defer file.Close()
	// var byteStr = make([]byte, 256)
	// n, err := file.Read(byteStr)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(n, string(byteStr))

	// 2. bufio.NewReader方式
	// file, err := os.Open("D:/gotest/test.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// defer file.Close()
	// reader := bufio.NewReader(file)
	// for {
	// 	line, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			if line != "" {
	// 				fmt.Print(line)
	// 			}
	// 			break
	// 		}
	// 		fmt.Println("error read file:", err)
	// 	}
	// 	fmt.Print(line)
	// }
	// 3. ioutil.ReadFile()方式
	// byteStr, err := os.ReadFile("D:/gotest/test.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(byteStr))

	// 写入文件的三种方式
	// 1. file.Write()
	// var str string
	// file, err := os.OpenFile("D:/gotest/test1.txt", os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	// for i := 0; i < 10; i++ {
	// 	str += "Hello,World\n"
	// }
	// data := []byte(str)
	// n, err := file.Write(data)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("写入数据成功--", n)
	// 2.bufio.NewWriter方式
	// var str string
	// file, err := os.OpenFile("D:/gotest/test1.txt", os.O_CREATE, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer file.Close()
	// writer := bufio.NewWriter(file)
	// for i := 0; i < 10; i++ {
	// 	str += "Hello,World\n"
	// }
	// // data := []byte(str)
	// n, err := writer.WriteString(str)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = writer.Flush()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("写入数据成功--", n)
	// 3. os.WriteFile，覆盖之前文件内容
	// var str = "hello,world"
	// data := []byte(str)
	// err := os.WriteFile("D:/gotest/test1.txt", data, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("写入数据成功")
}
