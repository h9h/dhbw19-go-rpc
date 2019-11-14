package main

import (
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	var err error
	var reply int
  
  args := Args{7,8}
  
	// Create a TCP connection to localhost on port 1234
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	client.Call("Arith.Multiply", args, &reply)
  fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)
}
