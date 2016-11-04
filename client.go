package client

import (
	"fmt"
	"net"
	"bufio"
)

var conn net.Conn

func Set(key string, value string) {
	request := fmt.Sprintf("SET %s %s", key, value)
	fmt.Fprintf(conn, request)
}

func Get(key string) string {
	request := fmt.Sprintf("GET %s", key)
	fmt.Fprintf(conn, request)
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		//just for now
		panic(err.Error())
	}
	return message
}

func DeleteKey(key string) {
	request := fmt.Sprintf("DELETE KEY %s", key)
	fmt.Fprintf(conn, request)
}

func DeleteBucket(bucket string) {
	request := fmt.Sprintf("DELETE BUCKET %s", bucket)
	fmt.Fprintf(conn, request)
}

func Authenticate(user string, password string, bucket string) {
	request := fmt.Sprintf("AUTH %s %s %s", user, password, bucket)
	fmt.Fprintf(conn, request)
}

func Connect(ip string, port int) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	_conn, err := net.Dial("tcp", addr)
	conn = _conn
	if err != nil {
		panic(err.Error())
	}
}


func Disconnect() {
	conn.Close()
}