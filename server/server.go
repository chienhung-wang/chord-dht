package server

import (
	"chord-dht/chord"
	pb "chord-dht/chord_pb"
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
)

type ChordNodeServer struct {
	storageService chord.StorageService
	nodeService    *chord.Node
	*pb.UnimplementedChordNodeServer
}

func NewChordNodeServer(ss chord.StorageService, n *chord.Node) *ChordNodeServer {
	return &ChordNodeServer{
		storageService: ss,
		nodeService:    n,
	}
}

func (s *ChordNodeServer) CheckAlive(ctx context.Context, in *pb.Empty) (*pb.Empty, error) {
	return &pb.Empty{Empty: true}, nil
}

func (s *ChordNodeServer) Join(ctx context.Context, in *pb.NodeAddr) (*pb.NodeAddr, error) {
	var id = big.NewInt(0).SetBytes(in.GetHash())
	if info, err := s.nodeService.Join(id, in.Addr); err == nil {
		return &pb.NodeAddr{Hash: info.Id.Bytes(), Addr: info.Addr}, nil
	} else {
		log.Println("join error", err)
		return nil, err
	}
}

func (s *ChordNodeServer) HelpFind(ctx context.Context, in *pb.NodeId) (*pb.NodeAddr, error) {
	id := new(big.Int).SetBytes(in.GetId())

	found, err := s.nodeService.Find(id)
	if err != nil {
		return new(pb.NodeAddr), err
	} else {
		return &pb.NodeAddr{Hash: found.Id.Bytes(), Addr: found.Addr}, nil
	}
}

func (s *ChordNodeServer) FindSuccessor(ctx context.Context, in *pb.NodeId) (*pb.FindFindSuccessorResp, error) {

	target := big.NewInt(0).SetBytes(in.GetId())
	found, next := s.nodeService.FindSucessor(target)

	return &pb.FindFindSuccessorResp{Found: found, Addr: &pb.NodeAddr{Hash: next.Id.Bytes(), Addr: next.Addr}}, nil
}

func (s *ChordNodeServer) GetFirstSuccessor(ctx context.Context, in *pb.Empty) (*pb.NodeAddr, error) {
	if s.nodeService.SuccList[0] == nil {
		return &pb.NodeAddr{Hash: *new([]byte), Addr: s.nodeService.Addr}, fmt.Errorf("no successor ")
	}
	succ := s.nodeService.GetAliveSuccessor()
	if succ == nil {
		return nil, fmt.Errorf("no successor ")
	}
	return &pb.NodeAddr{Hash: succ.Id.Bytes(), Addr: succ.Addr}, nil
}

func (s *ChordNodeServer) Notify(ctx context.Context, in *pb.NodeAddr) (*pb.Empty, error) {

	callerId := big.NewInt(0).SetBytes(in.GetHash())

	s.nodeService.Notify(&chord.NodeEntry{Id: callerId, Addr: in.Addr})

	return &pb.Empty{Empty: true}, nil
}

func (s *ChordNodeServer) GetPredecessor(ctx context.Context, empty *pb.Empty) (*pb.NodeAddr, error) {
	if pred, err := s.nodeService.GetPredecessor(); err == nil {
		return &pb.NodeAddr{Hash: pred.Id.Bytes(), Addr: pred.Addr}, nil
	} else {
		log.Println("getpred error: ", err)
		return nil, err
	}
}

func (s *ChordNodeServer) GetFingers(ctx context.Context, empty *pb.Empty) (*pb.GetFingersResp, error) {
	ids, addresses := s.nodeService.GetFingers()

	return &pb.GetFingersResp{Ids: ids, Addrs: addresses}, nil
}

func (s *ChordNodeServer) KeyTransfer(ctx context.Context, data *pb.KeyValueMap) (*pb.KeyTransferResp, error) {
	for k, v := range data.GetData() {
		s.storageService.Put(k, v)
	}

	return &pb.KeyTransferResp{Received: true}, nil
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
