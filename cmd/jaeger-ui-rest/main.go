package main

import (
	"flag"
	"log"
	"net"
	"os"

	"github.com/pkg/errors"
	jaegerui "github.com/tonistiigi/jaeger-ui-rest"
)

func main() {
	if err := run(); err != nil {
		log.Printf("Error: %+v", err)
		os.Exit(1)
	}
}

func run() error {
	var opt struct {
		addr string
	}

	flag.StringVar(&opt.addr, "addr", "127.0.0.1:0", "address to listen on")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		return errors.Errorf("unknown arguments: %v", args)
	}

	s := jaegerui.NewServer(jaegerui.Config{})

	ln, err := net.Listen("tcp", opt.addr)
	if err != nil {
		return errors.Wrapf(err, "failed to listen on %s", opt.addr)
	}

	log.Printf("Listening on %s", ln.Addr().String())

	return s.Serve(ln)
}
