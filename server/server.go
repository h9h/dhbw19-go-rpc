package main

import (
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
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
  arith := new(Arith)
	
  // Publish the receivers methods
	err := rpc.Register(arith)
	if err != nil {
		log.Fatal("Format of service Arith isn't correct. ", err)
	}
	// Register a HTTP handler
	rpc.HandleHTTP()
	// Listen to TPC connections on port 1234
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
	}

  log.Printf("Serving RPC server on port %d", 1234)

  // Start accept incoming HTTP connections
  err = http.Serve(listener, nil)
  if err != nil {
	  log.Fatal("Error serving: ", err)
  }
}
