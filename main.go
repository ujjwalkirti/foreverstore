package main

import (
	"fmt"
	"log"

	"github.com/ujjwalkirti/foreverstore/p2p"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	return nil
}
func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",
		Decoder:       p2p.DefaultDecoder{},
		ShakeHandFunc: p2p.NOPShakeHandFunc,
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v\n", msg)
		}
	}()
	err := tr.ListenAndAccept()
	if err != nil {
		log.Fatal(err)
	}
	select {}
}
