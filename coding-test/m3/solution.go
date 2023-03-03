package m3

import (
	"fmt"
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

// [M3 2000 sho tsuboya M3]
func DeleteDuplicate(strings []string) []string {
	// m[""]false が初期化した際の型
	m := make(map[string]bool)

	// 一応map[string]struct{}{}の方がいいかもしれない 値にアクセスしなくてもいいので
	// 詳細は # other_solution.goを参照

	var unique []string
	for _, v := range strings {
		// m[v]がtrueでなければ = まだそのキーはないということ
		if _, ok := m[v]; !ok {
			m[v] = true
			unique = append(unique, v)
		}
	}
	return unique

}

var regex = regexp.MustCompile(`[A-Z0-9]`) // 準備

func CheckRegex(s string) bool {
	ok := regex.MatchString(s) // 判定
	return ok
}
