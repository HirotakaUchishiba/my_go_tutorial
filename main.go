package main

import (
	"fmt"
	"math/rand"
	"os"
)

type URL struct {
	ShortURL string
	LongURL  string
}

var urlMap = make(map[string]string)

func generateShortURL() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, 6)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

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
			// 1. ユーザーがLongURLを入力
			fmt.Print("短縮したいURLを入力してください: ")
			var longURL string
			fmt.Scanln(&longURL)
			fmt.Println("URLを短縮します")
			shortURL := generateShortURL()
			urlMap[shortURL] = longURL
			fmt.Println("短縮URL: ", shortURL)
		case "2":
			fmt.Println("短縮URLを入力してください:")
			var shortURL string
			fmt.Scanln(&shortURL)
			longURL := urlMap[shortURL]
			fmt.Println("元のURL: ", longURL)
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
