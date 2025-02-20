package m3

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var count int

// アルファベット多文字 or 数字から始まる単語が何種類あるか
func Solution(str string) int {

	slice := strings.FieldsFunc(str, func(r rune) bool {
		return unicode.IsSpace(r) || r == '.'
	})
	fmt.Printf("slice: %v\n", slice)
	uniqueSlice := DeleteDuplicate(slice) // 重複の削除
	fmt.Printf("uniqueSlice: %v\n", uniqueSlice)
	for _, v := range uniqueSlice {
		// Favorite food is ramen. Age 25 years old. hobby coding
		initial := v[0:1]
		ok := CheckRegex(initial)
		if ok {
			count++
		}
	}
	return count
}

// [Favorite food is ramen. Age is 25 years old. Favorite hobby is coding]
func DeleteDuplicate(strings []string) []string {
	// m[""]false が初期化した際の型
	m := make(map[string]bool)

	// 一応map[string]struct{}{}の方がいいかもしれない 値にアクセスしなくてもいいので
	// 詳細は # other_solution.goを参照

	var unique []string
	for _, v := range strings {
		// 空の要素が含まれる場合があったらスルーする 最初の文字が空白だったり、ピリオドで終わった文章など
		// どうしても空の要素が含まれる場合がある
		if len(v) == 0 {
			continue
		}
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
