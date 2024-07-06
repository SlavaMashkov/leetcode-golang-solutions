package easy

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 70. Ближайшее число
// https://coderun.yandex.ru/problem/nearest-number

func NearestNumberMain() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	_, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return
	}

	scanner.Scan()
	split := strings.Split(scanner.Text(), " ")

	arr := make([]int, 0)
	for _, s := range split {
		i, _ := strconv.Atoi(s)

		arr = append(arr, i)
	}

	scanner.Scan()
	x, err := strconv.Atoi(scanner.Text())

	fmt.Println(NearestNumber(arr, x))
}

func NearestNumber(arr []int, x int) int {
	minDiff := math.MaxInt
	var nearestNumber int

	for _, v := range arr {
		if abs(v-x) <= minDiff {
			nearestNumber = v
			minDiff = abs(v - x)
		}
	}

	return nearestNumber
}

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}
