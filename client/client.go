package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	var err error
	var reply int
	var replyDivision Quotient

	args := Args{25, 8}

	// Create a TCP connection to localhost on port 1234
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Method Multiply Error: ", err)
	}
	fmt.Printf("Multiplikation: %d*%d=%d \n", args.A, args.B, reply)

	err = client.Call("Arith.Divide", args, &replyDivision)
	if err != nil {
		log.Fatal("Method Divide Error: ", err)
	}
	fmt.Printf("Division mit Rest: %d/%d=%d Rest %d \n", args.A, args.B, replyDivision.Quo, replyDivision.Rem)
}
