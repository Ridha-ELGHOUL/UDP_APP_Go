package net

import (
	"bufio"
	"fmt"
	"net"
)

func init_client_server() {
	p := make([]byte, 2048)
	conn, err := net.Dial("udp", "127.0.0.1:1234")

	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Printf(" UDP Client connected !!! \n")
	fmt.Fprintf(conn, "\n Hi UDP Server, How are you doing?")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()

}
func main() {
	init_client_server()
	fmt.Printf(" UDP Client TESTING ! \n")

}
