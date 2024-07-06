package easy

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 72. Возрастает ли список?
// https://coderun.yandex.ru/problem/list-growing

// AscendingListScan решено
func AscendingListScan() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	arr := make([]int, 0)
	split := strings.Split(scanner.Text(), " ")

	for _, s := range split {
		i, _ := strconv.Atoi(s)

		arr = append(arr, i)
	}

	fmt.Println(IsListAscending(arr))
}

func IsListAscending(arr []int) string {
	if len(arr) < 2 {
		return "YES"
	}

	for i := 1; i < len(arr); i++ {
		if arr[i-1] >= arr[i] {
			return "NO"
		}
	}

	return "YES"
}
