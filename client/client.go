package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func StartClient() {
	serverAddress := "/tmp/socket_file"
	fmt.Printf("Connecting to %s\n", serverAddress)

	conn, err := net.Dial("unix", serverAddress)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			fmt.Println("Exiting client.")
			break
		}

		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("Error sending message:", err)
			continue
		}

		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		buffer := make([]byte, 32)

		n, err := conn.Read(buffer)
		if err != nil {
			if os.IsTimeout(err) {
				fmt.Println("Socket timeout, ending listening for server messages")
				break
			}
			fmt.Println("Error reading from server:", err)
			break
		}

		data := string(buffer[:n])
		if len(data) > 0 {
			fmt.Printf("Server response: %s\n", data)
		} else {
			break
		}
	}

	fmt.Println("Closing socket")
}
