package HoneyBeeClient

import (
	"bufio"
	"fmt"
	"net"
)

type dbConn struct {
	net.Conn
	Bucket   string
	Username string
}

//Connect creates a TCP database connection with the specified server
func Connect(ip string, port int) dbConn {
	addr := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err.Error())
	}
	return dbConn{conn, "", ""}
}

func (conn *dbConn) Set(key string, value string, options string) {
	request := fmt.Sprintf("SET %s %s %s\n", key, value, options)
	conn.Write([]byte(request))
}

func (conn *dbConn) Get(key string, options string) string {
	request := fmt.Sprintf("GET %s %s\n", key, options)
	conn.Write([]byte(request))
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		//just for now
		panic(err.Error())
	}
	return message
}

func (conn *dbConn) DeleteKey(key string, options string) {
	request := fmt.Sprintf("DELETE KEY %s %s\n", key, options)
	conn.Write([]byte(request))
}

func (conn *dbConn) DeleteBucket(bucket string, options string) {
	request := fmt.Sprintf("DELETE BUCKET %s %s\n", bucket, options)
	conn.Write([]byte(request))
}

func (conn *dbConn) Authenticate(user string, password string) {
	request := fmt.Sprintf("AUTH %s %s\n", user, password)
	conn.Write([]byte(request))
}

func (conn *dbConn) UseBucket(bucket string) {
	request := fmt.Sprintf("USE %s\n", bucket)
	fmt.Printf(request)
	conn.Write([]byte(request))
}

func (conn *dbConn) Disconnect() {
	conn.Close()
}
