package main

import (
	"fmt"
)

// 標準入力から受け取った文字列を出力するコードになります。
func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}

func input(r io.Reader) <-chan string {
	// TODO: チャネルを作る
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// TODO: チャネルに読み込んだ文字列を送る
		}
		// TODO: チャネルを閉じる
	}()
	// TODO: チャネルを返す
}
