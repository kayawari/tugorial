package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// 乱数の種を初期化
	rand.Seed(time.Now().UnixNano())

	// コマンドライン引数の数をカウント
	// 設定されていなければエラーにする
	c := len(os.Args) - 1
	if c < 1 {
		fmt.Fprint(os.Stderr, "[usage] %s choice1 choice2 ...", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf(os.Args[rand.Intn(c)+1])
}
