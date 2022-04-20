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
	Succ           *NodeEntry
	SuccList       []*NodeEntry // Successor list
	Fingers        [161]*NodeEntry
	FingersStarts  [161]*big.Int
	FingersEnds    [161]*big.Int
	storageService StorageService
	Broken         bool
	IsNaive        bool
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

	}

	return &Node{
		Id:             id,
		Addr:           addr,
		Pred:           selfNodeEntry,
		SuccList:       make([]*NodeEntry, 3),
		Fingers:        fingers,
		FingersEnds:    fingerEnds,
		FingersStarts:  fingerStarts,
		storageService: storageService,
		Broken:         false,
		IsNaive:        false,
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
		// log.Println("")
		idx := rand.Intn(161-1) + 1
		found, err := n.find(n.FingersStarts[idx])
		if err != nil {
			log.Println("fix finger: no successor found", found)
		} else {
			n.Fingers[idx] = found
		}
		time.Sleep(time.Millisecond * 30)
	}
}

func (n *Node) JoinTo(succ *NodeEntry) error {
	n.SuccList[0] = succ
	n.Fingers[1] = succ

	err := RpcNotify(n.SuccList[0].Addr, &NodeEntry{Id: n.Id, Addr: n.Addr})
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

	return nil
}

func (n *Node) InitFingers(fingers []*NodeEntry) error {
	m := 160
	for i := 1; i <= m-1; i++ {
		if util.BetweenBeginInclusive(n.Id, n.FingersStarts[i+1], n.Fingers[i].Id) {
			n.Fingers[i+1] = n.Fingers[i]
		} else {
			var found bool
			var next = n.SuccList[0]
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

func (n *Node) UpdateBackupSuccessors() {
	for {
		time.Sleep(time.Second)

		if n.SuccList[0] == nil {
			continue
		}

		for i := 1; i < 3; i++ {
			firstSucc, err := RpcGetFirstSuccessor(n.SuccList[i-1].Addr)
			if err != nil || firstSucc.Id.Cmp(n.Id) == 0 {
				n.SuccList[i] = nil
				break
			}
			n.SuccList[i] = firstSucc
		}
	}
}

func (n *Node) GetAliveSuccessor() *NodeEntry {
	if n.SuccList[0] == nil {
		return nil
	}

	if err := RpcCheckAlive(n.SuccList[0].Addr); err == nil {
		return n.SuccList[0]
	}

	n.SuccList = n.SuccList[1:]
	n.SuccList = append(n.SuccList, nil)

	n.Broken = true
	defer func() { n.Broken = false }()

	for i := 1; i < 3; i++ {
		if n.SuccList[0] == nil {
			break
		}

		err := RpcCheckAlive(n.SuccList[0].Addr)
		if err == nil {
			for j := 1; j <= 160; j++ {
				if util.Between(n.Id, n.Fingers[j].Id, n.SuccList[0].Id) {
					n.Fingers[j] = n.SuccList[0]
				}
			}
			return n.SuccList[0]
		}
		n.SuccList = n.SuccList[1:]
		n.SuccList = append(n.SuccList, nil)
	}

	return nil
}

func (n *Node) Stabilize() {
	for {
		succ := n.SuccList[0]
		if succ != nil {
			succ_pred, err := RpcGetPredecessor(succ.Addr)
			if err != nil {
				n.Broken = true
				n.SuccList = n.SuccList[1:]
				if n.SuccList[0] != nil {
					for j := 1; j <= 160; j++ {
						if util.Between(n.Id, n.Fingers[j].Id, n.SuccList[0].Id) {
							n.Fingers[j] = n.SuccList[0]
						}
					}
				}
				n.Broken = false
				n.SuccList = append(n.SuccList, nil)
				continue
			}
			if succ_pred.Id.Cmp(n.Id) != 0 {
				if util.BetweenNoninclusive(n.Id, succ_pred.Id, succ.Id) {
					n.SuccList[0] = succ_pred
				}
			}

			err = RpcNotify(succ.Addr, &NodeEntry{Id: n.Id, Addr: n.Addr})
			if err != nil {
				fmt.Errorf("stabilization error: %v", err)
			}

		}

		time.Sleep(time.Millisecond * 300)
	}

}

func (n *Node) KeyTransfer(newPred *NodeEntry) map[string]string {
	table := n.storageService.GetLocalTable()
	res := make(map[string]string, 0)

	for k := range table {
		hash_val := util.Sha1_hash(k)
		if util.Between(n.Pred.Id, hash_val, newPred.Id) {
			res[k] = table[k]
			delete(table, k)
		}
	}

	return res

}

func (n *Node) VoluntarilyLeavingKeyTransfer() error {
	err := RpcKeyTransfer(n.GetAliveSuccessor().Addr, n.storageService.GetLocalTable())
	if err != nil {
		return fmt.Errorf("volauntarily leave error: %v", err)
	}
	return nil
}

func (n *Node) Notify(caller *NodeEntry) {
	if n.SuccList[0] == nil {
		RpcKeyTransfer(caller.Addr, n.KeyTransfer(caller))
		n.Pred = caller
		n.SuccList[0] = caller
		return
	}
	err := RpcCheckAlive(n.Pred.Addr)
	if util.BetweenNoninclusive(n.Pred.Id, caller.Id, n.Id) || err != nil {
		RpcKeyTransfer(caller.Addr, n.KeyTransfer(caller))
		fmt.Println("new pred -> ", caller.Addr)
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

	found, succ := n.FindSucessor(target)

	if found {
		return succ, nil
	}
	var next = succ
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
	if n.SuccList[0] == nil {
		return true, &NodeEntry{Id: n.Id, Addr: n.Addr}
	}

	succ := n.SuccList[0]

	if util.Between(n.Id, target, succ.Id) {
		return true, n.SuccList[0]
	}

	if !n.Broken && !n.IsNaive {
		if closest := n.ClosestProcedingFinger(target); closest != nil {
			return false, closest
		}
	}

	return false, succ
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

func (n *Node) GetKeyLocation(key string) (*NodeEntry, error) {

	succ, err := n.find(util.Sha1_hash(key))
	if err != nil {
		return nil, err
	}

	return succ, nil
}
