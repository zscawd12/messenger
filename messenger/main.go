package main

import (
	"fmt"

	"github.com/messenger/socket"
)

func main() {
	fmt.Println("socket start!!");

	socket.SocketConnect();
}
