package codingtest

import (
	"bufio"
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
}
