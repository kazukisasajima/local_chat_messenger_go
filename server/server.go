package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/bxcodec/faker/v3"
)

func StartServer() {
	serverAddress := "/tmp/socket_file"

	// 以前の接続でソケットファイルが残っていた場合は削除
	if _, err := os.Stat(serverAddress); err == nil {
		if err := os.Remove(serverAddress); err != nil {
			fmt.Println("Error removing existing socket file:", err)
			return
		}
	}

	// UNIXソケットを作成（TCP,ストリームモード）
	listener, err := net.Listen("unix", serverAddress)
	if err != nil {
		fmt.Println("Error creating UNIX socket:", err)
		return
	}
	defer listener.Close()

	fmt.Printf("Starting up on %s\n", serverAddress)

	// 無限ループでクライアントからの接続を待機
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 接続を平行処理
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	logWithTimestamp("Connection established")

	// 無限ループでクライアントからデータを受信
	for {
		buffer := make([]byte, 16)
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				logWithTimestamp("Client closed the connection.")
			} else {
				fmt.Println("Error reading from connection:", err)
			}
			break
		}

		dataStr := string(buffer[:n])
		logWithTimestamp(fmt.Sprintf("Received: %s", dataStr))

		if n > 0 {
			response := generateResponse(dataStr)
			_, err := conn.Write([]byte(response))
			if err != nil {
				fmt.Println("Error sending response:", err)
				break
			}
		} else {
			logWithTimestamp("No data from client, closing connection")
			break
		}
	}
}

func generateResponse(clientMessage string) string {
	switch clientMessage {
	case "name":
		return faker.Name()
	case "email":
		return faker.Email()
	default:
		return "Unknown request"
	}
}

func logWithTimestamp(message string) {
	fmt.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
}
