# HoneyBee clients

The HoneyBee DB is a simple key-value database I developed in Golang.

I wrote a client / API for it in Python and in Go for it to be used more simply.

## Go client
Is located in the [golang](golang) folder.
You can import it like this:
```go
import "github.com/talbor49/HoneyBeeClient/golang/HoneyBee"
```

Then, after you have a database running
(on how to do so - https://github.com/talbor49/HoneyBee)
you can create a datbase connection like this:
```go
ip := "127.0.0.1"
port := 8080
conn := HoneyBee.Connect(ip, port)
```

Then, you can start using it!
```go
conn.Authenticate("<username>", "<password>")
conn.CreateBucket("myBucket")
conn.UseBucket("myBucket")
log.Println("Setting...")
conn.Set("tal", "pro")
log.Println("Getting...")
conn.Get("tal")
```


**A full example program can be found at [main.go](main.go)**