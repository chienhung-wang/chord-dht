package chord_rpc

import (
	pb "chord-dht/chord_pb"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// type ChordNodeService struct {
// 	client *pb.ChordNodeClient
// }

// func NewChordNodeService() {

// }

func connectTo(address string) (pb.ChordNodeClient, *grpc.ClientConn, context.Context, context.CancelFunc, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("Connected to server...", address)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	c := pb.NewChordNodeClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	return c, conn, ctx, cancel, err
}

func Get(address string, key string) (string, string, error) {
	client, conn, ctx, cancel, err := connectTo(address)
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

func Put(address string, key string, val string) (string, string, error) {
	client, conn, ctx, cancel, err := connectTo(address)
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

func Delete(address string, key string) (string, error) {
	client, conn, ctx, cancel, err := connectTo(address)
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
