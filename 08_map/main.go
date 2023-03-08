package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func main() {
	a := make(map[string]int, 8)
	fmt.Println(a == nil)
	a["歌手"] = 100
	a["年假"] = 5
	fmt.Printf("a:%v\n", a)

	b := map[int]bool{
		1: true,
		2: true,
	}
	fmt.Println(b)
	go1()
	go2()
	go3()
	go4()
}

func go1() {
	// 判断某个键存不存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["沙河娜扎"] = 100
	scoreMap["张小凡"] = 10

	// 判断张小凡在不在
	v, ok := scoreMap["张小凡"]
	if ok {
		fmt.Println("张小凡在的", v)
	}
}

func go2() {
	// 遍历map
	var scoreMap = make(map[string]int, 8)
	scoreMap["沙河娜扎"] = 100
	scoreMap["张小凡"] = 10

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	delete(scoreMap, "沙河娜扎")
	fmt.Println(scoreMap)

	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println("scoreMap:\n", scoreMap)

	//按照key的升序排序
	keys := make([]string, 0, 100)
	for v := range scoreMap {
		keys = append(keys, v)
	}
	sort.Strings(keys)
	fmt.Println("scoreMap:\n", scoreMap)
}

func go3() {
	//切片初始化
	var mapSlice = make([]map[string]int, 8)
	// 内部map初始化
	mapSlice[0] = make(map[string]int, 8)
	mapSlice[0]["沙河小王子"] = 100
	fmt.Println(mapSlice)

	//值为切片的map
	// map初始化
	sliceMap := make(map[string][]int, 8)
	// 切片初始化
	sliceMap["中国"] = make([]int, 8)
	sliceMap["中国"][0] = 100
	sliceMap["中国"][1] = 55
	sliceMap["中国"][3] = 71
	fmt.Println(sliceMap)

}

func go4() {
	// 映射
	// 统计“how do you do” 中每个单词出现的次数
	s := "how do you do"
	mapString := make(map[string]int, 10)
	slice := strings.Split(s, " ")
	for _, value := range slice {
		v, ok := mapString[value]
		if ok {
			mapString[value] = v + 1
		} else {
			mapString[value] = 1
		}
	}
	fmt.Println(mapString)

}
