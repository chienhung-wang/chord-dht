package main

import (
	"bufio"
	"chord-dht/chord"
	pb "chord-dht/chord_pb"
	rpc "chord-dht/chord_rpc"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
)

func startServer(s *grpc.Server, lis net.Listener) error {
	err := s.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("Hello Chord")

	storageService := chord.NewStorageService()

	lis, err := net.Listen("tcp", "localhost:"+os.Args[1])
	if err != nil {
		log.Fatalln("Failed to listen to ", os.Args[1])
	}
	log.Println("Server listening at " + lis.Addr().String())
	s := grpc.NewServer()
	nodeServer := rpc.NewChordNodeServer(storageService)

	pb.RegisterChordNodeServer(s, nodeServer)

	log.Println("Server registered...")

	go func() {
		startServer(s, lis)
	}()

	log.Println("Start getting input...")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Command: ")
		scanner.Scan()
		input := scanner.Text()

		texts := strings.Split(input, " ")
		cmd := texts[0]

		switch cmd {
		case "GET":
			if len(texts) >= 3 {
				if key, val, err := rpc.Get(texts[2], texts[1]); err == nil {
					fmt.Printf("Key: %v, Val: %v\n -> GET\n", key, val)
				} else {
					fmt.Println("Error: ", err)
				}
			}
		case "PUT":
			if len(texts) >= 4 {
				if key, val, err := rpc.Put(texts[3], texts[1], texts[2]); err == nil {
					fmt.Printf("{Key: %v, Val: %v\n} -> PUT\n", key, val)
				} else {
					fmt.Println("Error: ", err)
				}
			}
		case "DELETE":
			if len(texts) >= 3 {
				if key, err := rpc.Delete(texts[2], texts[1]); err == nil {
					fmt.Printf("{Key: %v\n} -> DELET\n", key)
				} else {
					fmt.Println("Error: ", err)
				}
			}
		}

	}
}
