package main

import (
	_ "fmt"
	_ "strconv"
	_ "time"
)

// func test() int {
// 	fmt.Println("测试赋值函数")
// 	return 1
// }

//	func re(a, b int) int {
//		defer func() {
//			err := recover()
//			if err != nil {
//				fmt.Println("err=", err)
//			}
//		}()
//		return a / b
//	}
func main() {

	// const (
	// 	n1, n2 = iota + 1, iota + 2 //  1   2
	// 	n3, n4                      // 2    3
	// 	n5, n6                      //  3    4
	// )
	// fmt.Println(n1, n2, n3, n4, n5, n6)
	// var num float32 = 1.98888888
	// fmt.Printf("%.2f\n", num)
	// fmt.Println(unsafe.Sizeof(num))
	// var bo = false
	// if bo {
	// 	fmt.Println("一")
	// } else {
	// 	fmt.Println("二")
	// }
	// var s1 = "hello world"
	// var s2 = "l"
	// num := strings.Index(s1, s2)
	// fmt.Println(num)
	// var s = 'a'
	// fmt.Printf("类型是%T,值为%v", s, s)
	// var str = "你好 hello"
	// runeStr := []rune(str)
	// byteStr := []byte(str)
	// fmt.Println(runeStr, byteStr)
	// var s rune = 'a'
	// fmt.Printf("type: %T,value: %v", s, s)
	// var a int64 = 10
	// str := strconv.FormatInt(a, 16)
	// // fmt.Println(str)
	// if a := 10; a > 5 {
	// 	fmt.Println("success")
	// }
	// fmt.Println()
	// for i := 1; i <= 9; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("i * j = %v\t", i*j)
	// 	}
	// 	fmt.Println("")
	// }
	//数组初始化
	// var arr1 = [...]string{"ch", "wo"}
	// fmt.Println(len(arr1))
	// var s1 = []int{1, 2, 3, 4, 5}
	// fmt.Println(s1)
	//选择排序
	// var arr1 = [7]int{9, 22, 5, 111, 55, 4, 2}
	// for i := 0; i < len(arr1); i++ {
	// 	for j := i + 1; j < len(arr1); j++ {
	// 		if arr1[i] > arr1[j] {
	// 			temp := arr1[i]
	// 			arr1[i] = arr1[j]
	// 			arr1[j] = temp
	// 		}
	// 	}
	// }
	// fmt.Println(arr1)

	// 冒泡排序
	// var arr2 = [5]int{8, 26, 38, 5, 9}
	// for j := 0; j < len(arr2)-1; j++ {
	// 	for i := 0; i < len(arr2)-1-j; i++ {
	// 		if arr2[i] > arr2[i+1] {
	// 			temp := arr2[i]
	// 			arr2[i] = arr2[i+1]
	// 			arr2[i+1] = temp
	// 		}
	// 	}
	// }
	// fmt.Println(arr2)
	// a := test()
	// fmt.Println(a)
	// re(5, 0)
	// fmt.Println("后面代码继续执行...")
	// fmt.Println("111")
	//输出当前时间
	// timeObj := time.Now()
	// fmt.Println(timeObj)

	//日期字符串转时间戳
	// var str = "2020-04-26 15:38:04"
	// var tmp = "2006-01-02 15:04:05"
	// timeObj, _ := time.ParseInLocation(tmp, str, time.Local)
	// fmt.Println(timeObj.Unix())
	// fmt.Println("1")
	// //时间戳转日期字符串
	// fmt.Println(time.Unix(timeObj.Unix(), 0).Format("2006-01-02 15:04:05"))
	// var a interface{}
	// a = "hello"
	// reflect.TypeOf(a)

}
