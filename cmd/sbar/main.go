package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <command> [args...]\n", os.Args[0])
		os.Exit(1)
	}

	// TODO: ここで実際のコマンド実行とステータスバー表示を行う
	fmt.Println("sbar: not implemented yet")
}