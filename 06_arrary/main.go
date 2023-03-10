package main

import "fmt"

func main() {
	// 遍历
	var array1 = [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array1)
	for i := 0; i < len(array1); i++ {
		fmt.Printf("找到%v个 \n", i)
	}
	for _, value := range array1 {
		fmt.Println(value)
	}
	// 二维数组
	cityArray := [...][2]string{
		{"北京", "西安"}, {"上海", "南京"}, {"西宁", "汕头"},
	}
	fmt.Println(cityArray)
	codeTest()

}

func codeTest() {
	var arr1 = [...]uint{1, 3, 5, 7, 8}
	var result1 uint
	for _, value := range arr1 {
		result1 += value
	}
	fmt.Println(result1)

	var arr2 = [5]uint{1, 3, 5, 7, 8}
	for index1, value1 := range arr2 {
		for index2, value2 := range arr2 {
			if value1+value2 == 8 {
				fmt.Println(index1, index2)
				break
			}
		}
	}
}
