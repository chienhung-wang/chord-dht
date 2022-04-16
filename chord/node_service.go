package chord

import (
	"chord-dht/util"
	"math/big"
)

type NodeService interface {
}

type Node struct {
	Id      *big.Int
	Addr    string
	Fingers [161]*FingerEntry
}

type FingerEntry struct {
	Id   *big.Int
	Addr string
}

func NewNode(addr string) *Node {
	return &Node{
		Id:      util.Sha1_addr(addr),
		Addr:    addr,
		Fingers: [161]*FingerEntry{},
	}
}
