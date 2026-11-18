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

	for {
		connection, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		defer connection.Close()
		fmt.Printf("[%s] << connected\n", connection.RemoteAddr().String())

		go func() {
			connectionId := connection.RemoteAddr().String()
			err = ConnectionSession(connection)
			if err != nil {
				fmt.Printf("[%s] << error: %s\n", connectionId, err.Error())
			}
		}()
	}
}

func ConnectionSession(connection net.Conn) error {
	connectionId := connection.RemoteAddr().String()

	// Receiver
	buffer, err := bufio.NewReader(connection).ReadBytes('\n')
	if err != nil {
		fmt.Printf("[%s] << error: %s\n", connectionId, err.Error())
		return err
	}
	fmt.Printf("[%s] << %s", connectionId, string(buffer))

	message := "Thank you\n"

	fmt.Printf("[%s] >> %s\n", connectionId, message)
	connection.Write([]byte(message))
	return nil
}
