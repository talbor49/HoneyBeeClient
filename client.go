package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	ip := flag.String("ip", "localhost", "ip the server is on")
	port := flag.Int("port", 4590, "port Honey Bee is open on")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)

	conn, err := net.Dial("tcp", addr)

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(conn, "SET foo bar")
	//fmt.Fprintf()
}
