package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type ConnectionManager struct {
	// Map of connection id (client IP:port) and the connection object
	connections map[string]net.Conn
	mu          sync.Mutex
}

// Add connection
func (cm *ConnectionManager) Add(connId string, conn net.Conn) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	fmt.Printf("Connection added %s\n", connId)

	cm.connections[connId] = conn
	cm.printConnectionList()
}

// Remove connection
func (cm *ConnectionManager) Remove(connId string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	fmt.Printf("Connection removed %s\n", connId)

	delete(cm.connections, connId)
	cm.printConnectionList()
}

// Print the list of client connections
func (cm *ConnectionManager) printConnectionList() {
	fmt.Print("Connections: [ ")
	for connId, _ := range cm.connections {
		fmt.Printf("%s ", connId)
	}
	fmt.Println(" ]")
}

func (cm *ConnectionManager) SendMessage(sendertConnId string, message string) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	msgWithSender := fmt.Sprintf("%s >> %s", sendertConnId, message)

	for connId, connection := range cm.connections {
		if connId == sendertConnId {
			continue
		}
		_, err := connection.Write([]byte(msgWithSender))
		if err != nil {
			fmt.Printf("Failed to send message to %s err=%s\n", connId, err.Error())
			cm.Remove(connId)
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	fmt.Println("Listening on localhost:12345")

	connManager := ConnectionManager{
		connections: map[string]net.Conn{},
		mu:          sync.Mutex{},
	}

	for {
		connection, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		defer connection.Close()

		// Using goroutines
		// allowing multiple connection process concurently
		go ConnectionSession(connection, &connManager)
	}
}

func ConnectionSession(connection net.Conn, connManager *ConnectionManager) {
	connId := connection.RemoteAddr().String()
	connManager.Add(connId, connection)

	for {
		// Receiver
		buffer, err := bufio.NewReader(connection).ReadBytes('\n')
		if err != nil {
			fmt.Printf("%s << error: %s\n", connId, err.Error())
			connManager.Remove(connId)
			// return after connection has been removed
			return
		}
		fmt.Println(connId, "<<", string(buffer))
		connManager.SendMessage(connId, string(buffer))
	}
}
