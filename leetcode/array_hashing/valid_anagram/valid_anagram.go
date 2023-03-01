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
