package net

import (
	"gx/ipfs/QmTu65MVbemtUxJEWgsTtzv9Zv9P8rvmqNA4eG9TrTRGYc/go-libp2p-peer"
	"gx/ipfs/Qmd52WKRSwrBK5gUaJKawryZQ5by6UbNB8KVW2Zy6JtbyW/go-libp2p-host"

	"github.com/filecoin-project/go-filecoin/net/pubsub"
)

// Network is a unified interface for dealing with libp2p
type Network struct {
	host host.Host
	*pubsub.Subscriber
	*pubsub.Publisher
	*Router
}

// New returns a new Network
func New(host host.Host, publisher *pubsub.Publisher, subscriber *pubsub.Subscriber, router *Router) *Network {
	return &Network{
		host:       host,
		Subscriber: subscriber,
		Publisher:  publisher,
		Router:     router,
	}
}

// GetPeerID gets the current peer id from libp2p-host
func (network *Network) GetPeerID() peer.ID {
	return network.host.ID()
}
