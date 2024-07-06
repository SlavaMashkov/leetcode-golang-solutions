package easy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 58. OpenCalculator
// https://coderun.yandex.ru/problem/open-calculator

func IsInputPossibleMain() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	set := make(map[int]bool)
	split := strings.Split(scanner.Text(), " ")

	for _, s := range split {
		i, _ := strconv.Atoi(s)

		set[i] = true
	}

	scanner.Scan()
	arr := make([]int, 0)
	split = strings.Split(scanner.Text(), "")

	for _, s := range split {
		i, _ := strconv.Atoi(s)

		arr = append(arr, i)
	}

	fmt.Println(IsInputPossible(set, arr))
}

func IsInputPossible(set map[int]bool, arr []int) int {
	counter := 0

	for _, v := range arr {
		if _, exists := set[v]; !exists {
			set[v] = true
			counter++
		}
	}

	return counter
}
