package chord

import (
	"chord-dht/util"
	"errors"
	"fmt"
	"log"
	"math/big"
	"math/rand"
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
	Fingers        [161]*NodeEntry
	FingersStarts  [161]*big.Int
	FingersEnds    [161]*big.Int
	storageService StorageService
}

type NodeEntry struct {
	Id   *big.Int
	Addr string
}

func NewNode(addr string, storageService StorageService) *Node {
	fingers := [161]*NodeEntry{}
	fingerStarts := [161]*big.Int{}
	fingerEnds := [161]*big.Int{}
	id := util.Sha1_hash(addr)
	selfNodeEntry := &NodeEntry{Id: id, Addr: addr}

	for i := range fingers {
		fingers[i] = selfNodeEntry
	}

	m := int64(160)
	one := big.NewInt(1)
	two := big.NewInt(2)
	modVal := new(big.Int).Exp(two, big.NewInt(m), nil)
	fmt.Println("modVal: ", modVal)
	fingerEnds[0] = id
	for i := int64(1); i <= m; i++ {
		val := new(big.Int).Add(
			id,
			new(big.Int).Exp(
				two,
				big.NewInt(int64(i)-1),
				nil,
			),
		)
		fingerEnds[i] = new(big.Int).Mod(val, modVal)
		fingerStarts[i] = new(big.Int).Add(fingerEnds[i-1], one)

		// val2 := new(big.Int).Exp(
		// 	two,
		// 	new(big.Int).Sub(big.NewInt(i), one),
		// 	nil,
		// )
		// val3 := new(big.Int).Add(id, val2)
		// fmt.Println(new(big.Int).Mod(val3, modVal))
	}

	// fmt.Println(fingerEnds[:10])
	// fmt.Println(fingerStarts[:10])

	return &Node{
		Id:             id,
		Addr:           addr,
		Pred:           selfNodeEntry,
		Succ:           make([]*NodeEntry, 3),
		Fingers:        fingers,
		FingersEnds:    fingerEnds,
		FingersStarts:  fingerStarts,
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

func (n *Node) FixFinger() {
	rand.Seed(time.Now().UnixNano())

	for {
		idx := rand.Intn(161-1) + 1
		found, err := n.find(n.FingersStarts[idx])
		if err != nil {
			log.Println("fix finger: no successor found")
		} else {
			n.Fingers[idx] = found
		}
		time.Sleep(time.Millisecond * 80)
	}
}

func (n *Node) JoinTo(succ *NodeEntry) error {
	n.Succ[0] = succ
	n.Fingers[1] = succ

	err := RpcNotify(n.Succ[0].Addr, &NodeEntry{Id: n.Id, Addr: n.Addr})
	if err != nil {
		return fmt.Errorf("jointo notify error: %v", err)
	}

	fingers, err := RpcGetFingers(succ.Addr)
	if err != nil {
		return err
	}

	err = n.InitFingers(fingers)
	if err != nil {
		return err
	}

	go n.FixFinger()

	return nil
}

func (n *Node) InitFingers(fingers []*NodeEntry) error {
	m := 160
	fmt.Println(n.Fingers)
	for i := 1; i <= m-1; i++ {
		if util.BetweenBeginInclusive(n.Id, n.FingersStarts[i+1], n.Fingers[i].Id) {
			n.Fingers[i+1] = n.Fingers[i]
		} else {
			var found bool
			var next = n.Succ[0]
			var err error
			fmt.Printf("different ")
			for j := 0; j < 32; j++ {

				found, next, err = RpcFindSuccessor(next.Addr, n.FingersStarts[i+1])

				if err != nil {
					log.Println("init err:", err)
					return err
				}

				if found {
					break
				}
			}

			if !found {
				return fmt.Errorf("init finger error: sucessor not found")
			}

			n.Fingers[i+1] = next
		}
	}

	return nil
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

func (n *Node) Find(target *big.Int) (*NodeEntry, error) {
	found, err := n.find(target)
	if err != nil {
		return nil, err
	}

	return found, nil
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

	if util.Between(n.Id, target, n.Succ[0].Id) {
		return true, n.Succ[0]
	}

	if closest := n.ClosestProcedingFinger(target); closest != nil {
		return false, closest
	}

	// TODO: Check fingertable

	return false, n.Succ[0]
}

func (n *Node) ClosestProcedingFinger(target *big.Int) *NodeEntry {
	m := 160

	for i := m; i >= 1; i-- {
		if util.BetweenNoninclusive(n.Id, n.Fingers[i].Id, target) {
			return n.Fingers[i]
		}
	}

	return nil
}

func (n *Node) GetPredecessor() (*NodeEntry, error) {
	if n.Pred != nil {
		return n.Pred, nil
	} else {
		return nil, fmt.Errorf("%v has no predecessor", n.Addr)
	}
}

func (n *Node) GetFingers() ([][]byte, []string) {
	ids := make([][]byte, 161)
	addresses := make([]string, 161)

	for _, f := range n.Fingers {
		ids = append(ids, f.Id.Bytes())
		addresses = append(addresses, f.Addr)
	}

	return ids, addresses
}

// func (n *Node) UpdateFingers() {
// 	m := 160
// 	for i := 1; i <= m; i++ {

// 	}
// }

func (n *Node) GetKeyLocation(key string) (*NodeEntry, error) {

	succ, err := n.find(util.Sha1_hash(key))
	if err != nil {
		return nil, err
	}

	return succ, nil
}
