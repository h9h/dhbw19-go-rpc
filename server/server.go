package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	log.Printf("Got call for Multiply with %d and %d \n", args.A, args.B)
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	log.Printf("Got call for Divide with %d and %d \n", args.A, args.B)
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)

	// Publish the server methods
	err := rpc.Register(arith)
	if err != nil {
		log.Fatal("Publiskinh server methods failed ", err)
	}
	// Register a HTTP handler
	rpc.HandleHTTP()

	// Listen to TPC connections on port 1234
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
	}

	log.Printf("Serving RPC server methods on port %d \n", 1234)

	// Start accept incoming HTTP connections
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
