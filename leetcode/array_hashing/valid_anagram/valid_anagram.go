package valid_anagram

import (
	"fmt"
)

// s, t それぞれ文字列が与えられたとき、tがsのアナグラムであればtrue,そうでなければfalseを返す
func ValidAnagram(s, t string) bool {

	if len(s) != len(t) {
		fmt.Println("長さがまず同一ではないです")
		return false
	}
}

// 文字列を1つずつ確認するには、文字列をruneのスライスに変換する
// runeは、UTF-8エンコードされたUnicode文字列を表す
