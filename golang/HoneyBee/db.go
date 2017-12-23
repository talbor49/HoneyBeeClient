package HoneyBee

import (
	"bufio"
	"fmt"
	"github.com/talbor49/HoneyBee/grammar"
	"log"
	"net"
	"strconv"
)

type dbConn struct {
	net.Conn
	Bucket   string
	Username string
}

func (conn *dbConn) sendDbRequest(request grammar.Request) grammar.Response {
	rawRequest := grammar.BuildRawRequest(request)
	fmt.Printf("Raw request sending: %s\n", rawRequest)
	fmt.Printf("Raw request type sending: %s\n", strconv.Itoa(int(rawRequest[0])))
	conn.Write(rawRequest)
	rawResponse, _ := bufio.NewReader(conn).ReadString('\n')
	response := grammar.GetResponseFromBuffer([]byte(rawResponse))
	// log.Printf("Raw Response got: %s", rawResponse)
	fmt.Println(grammar.ResponseToString(response))
	return response
}

//Connect creates a TCP database connection with the specified server
func Connect(ip string, port int) dbConn {
	addr := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Connected to %s:%d", ip, port)
	return dbConn{conn, "", ""}
}

func (conn *dbConn) Set(key string, value string) {
	log.Printf("SET %s %s", key, value)
	setRequest := grammar.Request{Type: grammar.SET_REQUEST, Status: grammar.REQUEST_STATUS}
	setRequest.RequestData = []string{key, value}
	conn.sendDbRequest(setRequest)
}

func (conn *dbConn) Get(key string) {
	log.Printf("GET %s", key)
	getRequest := grammar.Request{Type: grammar.GET_REQUEST, Status: grammar.REQUEST_STATUS}
	getRequest.RequestData = []string{key}
	conn.sendDbRequest(getRequest)
}

func (conn *dbConn) DeleteKey(key string) {
	log.Printf("DELETE KEY %s", key)
	deleteRequest := grammar.Request{Type: grammar.DELETE_REQUEST, Status: grammar.KEY}
	deleteRequest.RequestData = []string{key}
	conn.sendDbRequest(deleteRequest)
}

func (conn *dbConn) DeleteBucket(bucket string) {
	log.Printf("DELETE BUCKET %s", bucket)
	deleteRequest := grammar.Request{Type: grammar.DELETE_REQUEST, Status: grammar.BUCKET}
	deleteRequest.RequestData = []string{bucket}
	conn.sendDbRequest(deleteRequest)
}

func (conn *dbConn) Authenticate(user string, password string) {
	log.Printf("AUTH %s %s\n", user, password)
	authRequest := grammar.Request{Type: grammar.AUTH_REQUEST, Status: grammar.REQUEST_STATUS}
	authRequest.RequestData = []string{user, password}
	conn.sendDbRequest(authRequest)
}

func (conn *dbConn) UseBucket(bucket string) {
	useRequest := grammar.Request{Type: grammar.USE_REQUEST, Status: grammar.REQUEST_STATUS}
	useRequest.RequestData = []string{bucket}
	log.Printf("USE %s", bucket)
	conn.sendDbRequest(useRequest)
}

func (conn *dbConn) Quit() {
	quitRequest := grammar.Request{Type: grammar.QUIT_REQUEST, Status: grammar.REQUEST_STATUS}
	log.Printf("QUIT")
	conn.sendDbRequest(quitRequest)
	conn.Close()
}

func (conn *dbConn) CreateBucket(bucket string) {
	log.Printf("CREATE BUCKET %s", bucket)
	createRequest := grammar.Request{Type: grammar.CREATE_BUCKET_REQUEST, Status: grammar.BUCKET}
	createRequest.RequestData = []string{bucket}
	conn.sendDbRequest(createRequest)
}

func (conn *dbConn) CreateUser(username string, password string) {
	log.Printf("CREATE USER %s:%s", username, password)
	createRequest := grammar.Request{Type: grammar.CREATE_USER_REQUEST, Status: grammar.USER}
	createRequest.RequestData = []string{username, password}
	conn.sendDbRequest(createRequest)
}
