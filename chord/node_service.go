package chord

import (
	"chord-dht/util"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"
)

type NodeService interface {
	Join(id *big.Int, addr string) *NodeEntry
	FindSucessor(target *big.Int) *NodeEntry
	GetPredecessor() *NodeEntry
}

type Node struct {
	Id             *big.Int
	Addr           string
	Pred           *NodeEntry
	Succ           []*NodeEntry // Successor list
	Fingers        [160]*NodeEntry
	storageService StorageService
}

type NodeEntry struct {
	Id   *big.Int
	Addr string
}

func NewNode(addr string, storageService StorageService) *Node {
	fingers := [160]*NodeEntry{}
	id := util.Sha1_hash(addr)
	selfNodeEntry := &NodeEntry{Id: id, Addr: addr}

	for i := range fingers {
		fingers[i] = selfNodeEntry
	}

	return &Node{
		Id:             id,
		Addr:           addr,
		Pred:           selfNodeEntry,
		Succ:           make([]*NodeEntry, 3),
		Fingers:        fingers,
		storageService: storageService,
	}
}

func (n *Node) Join(id *big.Int, addr string) (*NodeEntry, error) {

	found, err := n.find(id)

	if err != nil {
		return nil, err
	}

	return found, nil
}

func (n *Node) Stabilize() error {
	for {
		if n.Succ[0] != nil {
			succ_pred, err := RpcGetPredecessor(n.Succ[0].Addr)
			if err != nil {
				fmt.Errorf("stabilization error: %v", err)
			}
			if succ_pred.Id.Cmp(n.Id) != 0 {
				if util.BetweenNoninclusive(n.Id, succ_pred.Id, n.Succ[0].Id) {
					n.Succ[0] = succ_pred
				}
			}

			err = RpcNotify(n.Succ[0].Addr, &NodeEntry{Id: n.Id, Addr: n.Addr})
			if err != nil {
				fmt.Errorf("stabilization error: %v", err)
			}

		}

		time.Sleep(time.Second)
	}

}

func (n *Node) Notify(caller *NodeEntry) {
	if n.Succ[0] == nil {
		n.Pred = caller
		n.Succ[0] = caller
		return
	}

	if util.BetweenNoninclusive(n.Pred.Id, caller.Id, n.Id) {
		n.Pred = caller
	}
}

func (n *Node) find(target *big.Int) (*NodeEntry, error) {

	found, next := n.FindSucessor(target)

	if found {
		return next, nil
	}

	for i := 0; i < 32; i++ {
		var err error

		found, next, err = RpcFindSuccessor(next.Addr, target)

		if err != nil {
			log.Println("find err:", err)
			return nil, err
		}
		if found {
			return next, nil
		}
	}

	return nil, errors.New("id not found")

}

func (n *Node) FindSucessor(target *big.Int) (bool, *NodeEntry) {
	if n.Succ[0] == nil {
		return true, &NodeEntry{Id: n.Id, Addr: n.Addr}
	}

	// log.Printf("compare 3 values --> \n%v\n%v\n%v\n", n.Id, target, n.Succ[0].Id)

	if util.Between(n.Id, target, n.Succ[0].Id) {
		return true, n.Succ[0]
	}

	// TODO: Check fingertable

	return false, n.Succ[0]
}

func (n *Node) GetPredecessor() (*NodeEntry, error) {
	if n.Pred != nil {
		return n.Pred, nil
	} else {
		return nil, fmt.Errorf("%v has no predecessor", n.Addr)
	}
}

func (n *Node) GetKeyLocation(key string) (*NodeEntry, error) {

	succ, err := n.find(util.Sha1_hash(key))
	if err != nil {
		return nil, err
	}

	return succ, nil
}

// func (kv *hashTable) Put(key string, val string) error {
// 	kv.table[key] = val
// 	return nil
// }

// func (kv *hashTable) Delete(key string) error {
// 	if _, ok := kv.table[key]; ok {
// 		delete(kv.table, key)
// 		return nil
// 	} else {
// 		return errors.New("key not found")
// 	}
// }
