package m3

import (
	"bufio"
	"fmt"
	"os"
)

// コードレビューでパフォーマンス的なところで指摘をいただけた

// string.Splitで全てを読み込むとなると、テキストファイルのサイズが大きいと変数の容量も増加する
// 変数のサイズが大きくなるとメモリ確保やコピーなどに時間がかかる

// 今回の問題は全てのテキスト情報を保持しなくても、文字列をスペース等がくるまで読み取って処理を行い、
// 単語の重複チェックを行なっていくことで既に処理した文字列を破棄することが出来る

func OtherSolution() {
	scanner := bufio.NewScanner(os.Stdin)

	// どのように区切るかの設定を行う
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// 空白で区切る
		for i := 0; i < len(data); i++ {
			if data[i] == ' ' {
				return i + 1, data[:i], nil
			}
		}
		return 0, data, bufio.ErrFinalToken
	})

	// 1つずつ、見ていく -> 重複があったらそれはカウントしない
	words := map[string]struct{}{}
	for scanner.Scan() {
		word := scanner.Text()
		capital := word[0] // 単語の最初の文字を取得
		if 'A' <= capital && capital <= 'Z' || '0' <= capital && capital <= '9' {
			// 一致したらキーにその単語を、値は空にする
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
