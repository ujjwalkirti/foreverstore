package main

import (
	"log"

	"github.com/ujjwalkirti/foreverstore/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		Decoder:       p2p.DefaultDecoder{},
		ShakeHandFunc: p2p.NOPShakeHandFunc,
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	err := tr.ListenAndAccept()
	if err != nil {
		log.Fatal(err)
	}
	select {}
}
