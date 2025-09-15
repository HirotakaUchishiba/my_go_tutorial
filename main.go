package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("1: URLを短縮")
		fmt.Println("2: URLを元に戻す")
		fmt.Println("3: 終了")
		fmt.Print("番号を入力してください: ")

		var choice string
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Println("URLを短縮します")
		case "2":
			fmt.Println("URLを元に戻します")
		case "3":
			{
				fmt.Println("終了します")
				os.Exit(0)
			}

		default:
			fmt.Println("無効な選択です")
		}
		fmt.Println()
	}

}
