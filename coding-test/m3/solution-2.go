package m3

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

// コードレビューでパフォーマンス的なところで指摘をいただけた

// string.Splitで全てを読み込むとなると、テキストファイルのサイズが大きいと変数の容量も増加する
// 変数のサイズが大きくなるとメモリ確保やコピーなどに時間がかかる

// 今回の問題は全てのテキスト情報を保持しなくても、文字列をスペース等がくるまで読み取って処理を行い、
// 単語の重複チェックを行なっていくことで既に処理した文字列を破棄することが出来る

func OtherSolution() {
	scanner := bufio.NewScanner(os.Stdin)

	// Scan()内部で使用するために、どのように区切るかの設定を行う
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// 空白で区切る
		for i := 0; i < len(data); i++ {
			if data[i] == ' ' || data[i] == '.' {
				return i + 1, data[:i], nil
			}
		}
		return 0, data, bufio.ErrFinalToken
	})

	words := map[string]struct{}{}
	for scanner.Scan() {
		word := scanner.Text()
		fmt.Printf("word: %v\n", word)
		// 空の要素が含まれる場合があったらスルーする 最初の文字が空白だったり、ピリオドで終わった文章など
		// どうしても空の要素が含まれる場合がある
		if len(word) == 0 {
			continue
		}
		capital := rune(word[0]) // runeで比較したいので変換する
		if unicode.IsUpper(capital) || unicode.IsDigit(capital) {
			// 一致したらキーにその単語を、値は空で設定する
			words[word] = struct{}{}
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Printf("%d words\n\n", len(words))
	// mapをfor range で回す
	for word := range words {
		fmt.Println(word)
	}
}
