package udpServer

import (
	"fmt"
	"net"
)

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your mesage "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}
func init_udp_server(addr net.UDPAddr) {

	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	p := make([]byte, 2048)
	for {
		_, remoteaddr, err := ser.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
	}
}

func main() {

	addr := net.UDPAddr{
		Port: 1234,
		IP:   net.ParseIP("127.0.0.1")}
	init_udp_server(addr)
	//go sendResponse(ser, remoteaddr)
}
