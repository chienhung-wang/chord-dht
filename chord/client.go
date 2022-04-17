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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return c, conn, ctx, cancel, err
}

func RpcJoin(address string, ownAddr string, id *big.Int) (*NodeEntry, error) {
	client, conn, ctx, cancel, err := connectTo(address)
	if err != nil {
		log.Fatalln(err)
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

func RpcFindSuccessor(address string, id *big.Int) (bool, *NodeEntry, error) {
	client, conn, ctx, cancel, err := connectTo(address)
	defer conn.Close()
	defer cancel()

	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.FindSuccessor(ctx, &pb.NodeId{Id: id.Bytes()})

	if err != nil {
		return false, nil, err
	} else {
		succ := res.GetAddr()
		succ_id := big.NewInt(0).SetBytes(succ.GetHash())
		return res.GetFound(), &NodeEntry{Id: succ_id, Addr: succ.GetAddr()}, nil
	}
}

func RpcGetPredecessor(address string) (*NodeEntry, error) {
	client, conn, ctx, cancel, err := connectTo(address)
	defer conn.Close()
	defer cancel()
	if err != nil {
		log.Fatalln(err)
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

func (n *Node) Get(key string) (string, string, error) {
	loc, err := n.GetKeyLocation(key)
	if err != nil {
		return "", "", err
	}

	client, conn, ctx, cancel, err := connectTo(loc.Addr)
	defer conn.Close()
	defer cancel()

	if err != nil {
		log.Fatalln(err)
	}

	kv, err := client.MapGet(ctx, &pb.Key{Key: key})

	if err != nil {
		return "", "", err
	} else {
		return kv.GetKey(), kv.GetVal(), nil
	}
}

func (n *Node) Put(key string, val string) (string, string, error) {
	loc, err := n.GetKeyLocation(key)
	if err != nil {
		return "", "", err
	}

	client, conn, ctx, cancel, err := connectTo(loc.Addr)
	defer conn.Close()
	defer cancel()

	if err != nil {
		log.Fatalln(err)
	}

	kv, err := client.MapPut(ctx, &pb.KeyVal{Key: key, Val: val})

	if err != nil {
		return "", "", err
	} else {
		return kv.GetKey(), kv.GetVal(), nil
	}
}

func (n *Node) Delete(key string) (string, error) {
	loc, err := n.GetKeyLocation(key)
	if err != nil {
		return "", err
	}

	client, conn, ctx, cancel, err := connectTo(loc.Addr)
	defer conn.Close()
	defer cancel()

	if err != nil {
		log.Fatalln(err)
	}

	k, err := client.MapDelete(ctx, &pb.Key{Key: key})

	if err != nil {
		return "", err
	} else {
		return k.GetKey(), nil
	}
}
