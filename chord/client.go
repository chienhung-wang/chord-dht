package chord

import (
	pb "chord-dht/chord_pb"
	"context"
	"log"
	"math/big"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func connectTo(address string) (pb.ChordNodeClient, *grpc.ClientConn, context.Context, context.CancelFunc, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// log.Println("Connected to server...", address)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	c := pb.NewChordNodeClient(conn)

	time.Sleep(time.Microsecond * 300)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)

	return c, conn, ctx, cancel, err
}

func RpcJoin(address string, ownAddr string, id *big.Int) (*NodeEntry, error) {
	client, conn, ctx, cancel, err := connectTo(address)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	defer cancel()

	res, err := client.Join(ctx, &pb.NodeAddr{Hash: id.Bytes(), Addr: ownAddr})
	if err != nil {
		return nil, err
	}

	succId := big.NewInt(0).SetBytes(res.GetHash())

	return &NodeEntry{Id: succId, Addr: res.Addr}, nil
}

func RpcCheckAlive(address string) error {
	client, conn, ctx, cancel, err := connectTo(address)
	defer conn.Close()
	defer cancel()

	if err != nil {
		log.Println(err)
	}

	_, err = client.CheckAlive(ctx, &pb.Empty{Empty: true})
	if err != nil {
		return err
	}

	return nil
}

func RpcFindSuccessor(address string, id *big.Int, query int) (bool, *NodeEntry, int, error) {
	client, conn, ctx, cancel, err := connectTo(address)
	defer conn.Close()
	defer cancel()

	if err != nil {
		log.Println(err)
	}

	res, err := client.FindSuccessor(ctx, &pb.NodeId{Id: id.Bytes(), Query: int32(query + 1)})

	if err != nil {
		return false, nil, -1, err
	} else {
		succ := res.GetAddr()
		succ_id := big.NewInt(0).SetBytes(succ.GetHash())
		return res.GetFound(), &NodeEntry{Id: succ_id, Addr: succ.GetAddr()}, int(res.Query), nil
	}
}

func RpcGetPredecessor(address string) (*NodeEntry, error) {
	client, conn, ctx, cancel, err := connectTo(address)
	defer conn.Close()
	defer cancel()
	if err != nil {
		log.Println(err)
	}

	pred, err := client.GetPredecessor(ctx, &pb.Empty{Empty: true})
	if err != nil {
		log.Println("rpc get pred: ", err)
		return nil, err
	}

	id := big.NewInt(0).SetBytes(pred.GetHash())
	predEntry := &NodeEntry{Id: id, Addr: pred.GetAddr()}

	return predEntry, nil
}

func RpcGetFirstSuccessor(address string) (*NodeEntry, error) {
	client, conn, ctx, cancel, err := connectTo(address)
	defer conn.Close()
	defer cancel()
	if err != nil {
		log.Println(err)
	}

	firstSucc, err := client.GetFirstSuccessor(ctx, &pb.Empty{Empty: true})
	if err != nil {
		log.Println("rpc get first successor: ", err)
		return nil, err
	}

	id := big.NewInt(0).SetBytes(firstSucc.GetHash())

	return &NodeEntry{Id: id, Addr: firstSucc.GetAddr()}, nil
}

func RpcGetFingers(address string) ([]*NodeEntry, error) {
	client, conn, ctx, cancel, err := connectTo(address)
	defer conn.Close()
	defer cancel()
	if err != nil {
		log.Println(err)
	}

	table, err := client.GetFingers(ctx, &pb.Empty{Empty: true})
	if err != nil {
		return nil, err
	}

	fingers := make([]*NodeEntry, 161)
	addresses := table.GetAddrs()

	id_bytes := table.GetIds()
	for i := 0; i < 161; i++ {
		fingers[i] = &NodeEntry{Id: big.NewInt(0).SetBytes(id_bytes[i]), Addr: addresses[i]}
	}

	return fingers, nil
}

func RpcNotify(address string, src *NodeEntry) error {
	client, conn, ctx, cancel, err := connectTo(address)
	defer conn.Close()
	defer cancel()
	if err != nil {
		return err
	}

	client.Notify(ctx, &pb.NodeAddr{Hash: src.Id.Bytes(), Addr: src.Addr})

	return nil
}

func RpcKeyTransfer(address string, table map[string]string) error {
	client, conn, ctx, cancel, err := connectTo(address)
	defer conn.Close()
	defer cancel()
	if err != nil {
		return err
	}

	_, err = client.KeyTransfer(ctx, &pb.KeyValueMap{Data: table})
	if err != nil {
		return err
	}

	return nil
}

func (n *Node) Get(key string) (string, string, int, error) {
	loc, query, err := n.GetKeyLocation(key)
	if err != nil {
		return "", "", -1, err
	}

	client, conn, ctx, cancel, err := connectTo(loc.Addr)
	defer conn.Close()
	defer cancel()

	if err != nil {
		log.Println(err)
	}

	kv, err := client.MapGet(ctx, &pb.Key{Key: key})

	if err != nil {
		return "", "", -1, err
	} else {
		return kv.GetKey(), kv.GetVal(), query, nil
	}
}

func (n *Node) Put(key string, val string) (string, string, int, error) {
	loc, query, err := n.GetKeyLocation(key)
	if err != nil {
		return "", "", -1, err
	}

	client, conn, ctx, cancel, err := connectTo(loc.Addr)
	defer conn.Close()
	defer cancel()

	if err != nil {
		log.Println(err)
	}

	kv, err := client.MapPut(ctx, &pb.KeyVal{Key: key, Val: val})

	if err != nil {
		return "", "", -1, err
	} else {
		return kv.GetKey(), kv.GetVal(), query, nil
	}
}

func (n *Node) Delete(key string) (string, error) {
	loc, _, err := n.GetKeyLocation(key)
	if err != nil {
		return "", err
	}

	client, conn, ctx, cancel, err := connectTo(loc.Addr)
	defer conn.Close()
	defer cancel()

	if err != nil {
		log.Println(err)
	}

	k, err := client.MapDelete(ctx, &pb.Key{Key: key})

	if err != nil {
		return "", err
	} else {
		return k.GetKey(), nil
	}
}
