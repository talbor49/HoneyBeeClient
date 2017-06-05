package main

import (
	"log"
	"github.com/talbor49/HoneyBeeClient/client"
)

func main() {
	log.Println("Hello I am main! :)")
	ip := "127.0.0.1"
	port := 8080
	conn := client.Connect(ip, port)
	log.Println("Authenticating...")
	conn.Authenticate("talbor49", "1234")
	//log.Println("Creating...")
	//conn.CreateBuck
	conn.UseBucket("x")
	log.Println("Setting...")
	conn.Set("tal", "pro")
	log.Println("Getting...")
	conn.Get("tal")
	defer conn.Quit()
}
