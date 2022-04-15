package chord_rpc

import (
	"chord-dht/chord"
	pb "chord-dht/chord_pb"
	"context"
	"errors"
)

type ChordNodeServer struct {
	storageService chord.StorageService
	*pb.UnimplementedChordNodeServer
}

func NewChordNodeServer(ss chord.StorageService) *ChordNodeServer {
	return &ChordNodeServer{
		storageService: ss,
	}
}

func (s *ChordNodeServer) MapGet(ctx context.Context, in *pb.Key) (*pb.KeyVal, error) {
	if val, err := s.storageService.Get(in.GetKey()); err == nil {
		return &pb.KeyVal{
			Key: in.GetKey(),
			Val: val,
		}, nil
	} else {
		return nil, errors.New("key not found")
	}
}
func (s *ChordNodeServer) MapPut(ctx context.Context, in *pb.KeyVal) (*pb.KeyVal, error) {
	s.storageService.Put(in.Key, in.Val)

	return &pb.KeyVal{
		Key: in.GetKey(),
		Val: in.GetVal(),
	}, nil
}
func (s *ChordNodeServer) MapDelete(ctx context.Context, in *pb.Key) (*pb.Key, error) {
	if err := s.storageService.Delete(in.GetKey()); err == nil {
		return &pb.Key{
			Key: in.GetKey(),
		}, nil
	} else {
		return nil, errors.New("key not found")
	}
}
