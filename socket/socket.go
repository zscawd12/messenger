package socket

import (
	"net"
	"fmt"
)

type (

	socket struct {
		id int
		name string
		
	}

)

const (
	CONN_HOST = 
	CONN_PORT = 
	CONN_TYPE = "tcp"
)

func SocketConnect() {
	listener, err := net.Listen(CONN_TYPE, CONN_HOST+CONN_PORT)
	if err != nil {
		fmt.Printf("failed to Listen : %v\n", err)
		return
	}
	defer listener.Close()
	
	fmt.Println("Listening on " + CONN_HOST + CONN_PORT)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("failed to accepting : %v\n", err)
			return
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("failed to connection read : %v\n", err)
	}
	fmt.Printf("%v 개\n", reqLen)
	conn.Write([]byte("Message recived.\n"))
}
