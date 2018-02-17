package main

import (
	"github.com/talbor49/HoneyBeeClient/golang/HoneyBee"
	"log"
)


func main() {
	ip := "127.0.0.1"
	port := 8080
	conn := HoneyBee.Connect(ip, port)
	log.Println("Authenticating...")
	conn.CreateUser("talbor49", "1234")
	conn.Authenticate("talbor49", "1234")
	//log.Println("Creating...")
	//conn.CreateBuck
	conn.CreateBucket("myBucket")
	conn.UseBucket("myBucket")
	log.Println("Setting...")
	conn.Set("tal", "pro")
	log.Println("Getting...")
	conn.Get("tal")
	defer conn.Quit()
}
