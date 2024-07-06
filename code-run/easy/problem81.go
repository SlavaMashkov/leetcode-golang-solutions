package easy

import (
	"fmt"
	"log"
)

// 81. Треугольник
// https://coderun.yandex.ru/problem/triangle

func TriangleMain() {
	var arr = make([]int, 3)
	for i := 0; i < len(arr); i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(IsTriangle(arr))
}

func IsTriangle(arr []int) string {
	if arr[0]+arr[1] <= arr[2] {
		return "NO"
	}

	if arr[0]+arr[2] <= arr[1] {
		return "NO"
	}

	if arr[1]+arr[2] <= arr[0] {
		return "NO"
	}

	return "YES"
}
