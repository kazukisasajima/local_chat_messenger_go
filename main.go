package main

import (
	"fmt"
	"os"

	"local_chat_messenger_go/server"
	"local_chat_messenger_go/client"
)

func main() {
	// コマンドライン引数のチェック
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [server|client]")
		return
	}

	// サーバーまたはクライアントを選択して実行
	switch os.Args[1] {
	case "server":
		server.StartServer()
	case "client":
		client.StartClient()
	default:
		fmt.Println("Unknown command. Use 'server' or 'client'.")
	}
}
