package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	fmt.Printf("<< connected to %s\n", connection.RemoteAddr().String())

	message := "Hello\n"
	fmt.Println(">>", message)
	_, err = connection.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	// Receiver
	buffer, err := bufio.NewReader(connection).ReadBytes('\n')
	if err != nil {
		fmt.Printf("<< error: %s\n", err.Error())
		return
	}
	fmt.Println("<<", string(buffer))
}
