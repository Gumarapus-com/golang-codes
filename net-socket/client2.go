package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		panic(err)
	}
	defer connection.Close()
	fmt.Println("connected")

	var wg sync.WaitGroup

	// Receiver
	wg.Go(func() {
		for {
			buffer, err := bufio.NewReader(connection).ReadBytes('\n')
			if err != nil {
				fmt.Printf("<< error: %s\n", err.Error())
				return
			}
			fmt.Println(string(buffer))
		}
	})

	// Sender
	wg.Go(func() {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(">> ")
		for scanner.Scan() {
			message := scanner.Text()
			if message == "" {
				fmt.Print(">> ")
				continue
			}
			_, err := connection.Write([]byte(fmt.Sprintln(message)))
			if err != nil {
				fmt.Printf("Err: %s\n", err.Error())
				break
			}
			fmt.Print(">> ")
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "read error:", err)
		}
	})

	wg.Wait()
}
