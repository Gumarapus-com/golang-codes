package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		panic(err)
	}
	defer listen.Close()

	connection, err := listen.Accept()
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	fmt.Printf("<< connected to %s\n", connection.RemoteAddr().String())

	// Receiver
	buffer, err := bufio.NewReader(connection).ReadBytes('\n')
	if err != nil {
		fmt.Printf("<< error: %s\n", err.Error())
		return
	}
	fmt.Println("<<", string(buffer))

	message := "Thank you\n"
	fmt.Println(">>", message)
	connection.Write([]byte(message))
}
