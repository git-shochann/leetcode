package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var count int

// アルファベット多文字 or 数字から始まる単語が何種類あるか
func Solution(str string) int {

	slice := strings.Split(str, " ")
	fmt.Printf("slice: %v\n", slice)
	uniqueSlice := DeleteDuplicate(slice) // 重複の削除
	fmt.Printf("uniqueSlice: %v\n", uniqueSlice)
	for _, v := range uniqueSlice {
		// [M3 2000 sho tsuboya]
		initial := v[0:1]
		ok := CheckRegex(initial)
		if ok {
			count++
		}
	}
	return count
}

func DeleteDuplicate(strings []string) []string {
	var unique []string
	encounter := map[string]int{} // {"test1":1, "test2":2}
	for _, v := range strings {
		// v -> [M3 2000 sho tsuboya]
		if _, ok := encounter[v]; !ok {
			encounter[v] = 1
			unique = append(unique, v)
		} else {
			encounter[v]++
		}
	}
	return unique
}

var regex = regexp.MustCompile(`[A-Z0-9]`) // 準備

func CheckRegex(s string) bool {
	ok := regex.MatchString(s) // 判定
	return ok
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Printf("Solution(text): %v\n", Solution(text))
}
