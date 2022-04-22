package main

import (
	"bufio"
	"chord-dht/chord"
	pb "chord-dht/chord_pb"
	rpc "chord-dht/server"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
)

func startServer(s *grpc.Server, lis net.Listener) error {
	err := s.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

func getAddr() (port string, host_port string) {
	// host := os.Args[1]
	port = os.Args[2]
	host_port = ":" + port
	return
}

func main() {
	fmt.Println("Hello Chord")

	port, host_port := getAddr()

	storageService := chord.NewStorageService()
	node := chord.NewNode(host_port, storageService)
	id := node.Id

	lis, err := net.Listen("tcp", host_port)
	if err != nil {
		log.Fatalln("Failed to listen to port", port)
	}
	log.Println("Server listening at " + lis.Addr().String())

	s := grpc.NewServer()
	rpcServer := rpc.NewChordNodeServer(storageService, node)

	pb.RegisterChordNodeServer(s, rpcServer)

	log.Println("Server registered...")

	go startServer(s, lis)

	go node.Stabilize()

	go node.UpdateBackupSuccessors()

	go node.FixFinger()

	node.Extended = false
	if node.Extended {
		go node.FixExtendedFingers()
	}

	log.Println("Node id ---> ", id)

	log.Println("Start getting input...")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		texts := strings.Split(input, " ")
		cmd := texts[0]

		switch cmd {
		case "JOIN":
			if len(texts) >= 2 {
				addr := texts[1]
				if succ, err := chord.RpcJoin(addr, node.Addr, node.Id); err == nil {
					fmt.Printf("found successor's id: %v, addr: %v\n --> JOIN\n", succ.Id, succ.Addr)
					node.JoinTo(succ)
				} else {
					fmt.Println("Error: ", err)
				}
			}
		case "SUCC":
			if len(texts) >= 1 {
				if node.SuccList[0] == nil {
					fmt.Println("Successor is nil")
				} else {
					fmt.Println("Successor -> ", node.SuccList[0])
				}

			}
		case "SELF":
			fmt.Printf("Self:\nid: %v\nport: %v\n", node.Id, port)
		case "MAP":
			if len(texts) >= 1 {
				fmt.Println("Local Hash Table -> \n", storageService.GetLocalTable())
			}
		case "STAB":
			if len(texts) >= 1 {
				node.Stabilize()
			}
		case "PRED":
			if len(texts) >= 1 {
				fmt.Println("Predescessor -> ", node.Pred)
			}
		case "FINGERS":
			if len(texts) >= 1 {
				ring := make([]string, 0)
				for _, f := range node.Fingers {
					ring = append(ring, f.Addr[10:])
				}
				fmt.Println("Fingers -> ", ring)
			}
		case "SUCCLIST":
			println("SuccList -> ")
			for _, suc := range node.SuccList {
				if suc == nil {
					break
				}
				fmt.Printf(" -> %v ", suc.Addr)
			}
			println()
		case "GET":
			if len(texts) >= 2 {
				if key, val, query, err := node.Get(texts[1]); err == nil {
					fmt.Printf("{Key: %v, Val: %v} -> GET\n", key, val)
					fmt.Printf("%v hops ~ \n", query)
				} else {
					fmt.Println("Error: ", err)
				}
			}
		case "PUT":
			if len(texts) >= 3 {
				if key, val, query, err := node.Put(texts[1], texts[2]); err == nil {
					fmt.Printf("{Key: %v, Val: %v} -> PUT\n", key, val)
					fmt.Printf("%v hops ~ \n", query)
				} else {
					fmt.Println("Error: ", err)
				}
			}
		case "DELETE":
			if len(texts) >= 2 {
				if key, err := node.Delete(texts[1]); err == nil {
					fmt.Printf("{Key: %v} -> DELET\n", key)
				} else {
					fmt.Println("Error: ", err)
				}
			}
		case "TEST":
			var rounds int
			rounds, err = strconv.Atoi(texts[1])

			buff := make([]string, rounds)
			if err != nil {
				fmt.Println(err)
			}
			if len(texts) >= 2 {
				start := time.Now()
				puthops := 0
				for i := 0; i < rounds; i++ {
					rs := StringWithCharset(10, charset)
					buff[i] = rs
					_, _, query, err := node.Put(strconv.Itoa(i+1), strconv.Itoa(i+1))
					time.Sleep(time.Millisecond * 3)
					if err != nil {
						fmt.Println("PUT Err: ", err)
						continue
					}
					puthops += query
				}
				add := time.Now()
				gethops := 0
				for i := 0; i < rounds; i++ {
					_, _, query, err := node.Get(strconv.Itoa(i + 1))
					time.Sleep(time.Millisecond * 3)
					if err != nil {
						fmt.Println("GET Err: ", err)
						continue
					}
					gethops += query
				}

				get := time.Now()
				fmt.Printf("Total time to finish test: %s \n", time.Since(start).String())
				fmt.Printf("Total time to put : %s \n", add.Sub(start))
				fmt.Printf("Total time to get : %s \n", get.Sub(add))
				fmt.Printf("Total number of hops:\n")
				fmt.Printf("PUT: %v, GET: %v\n", puthops, gethops)
			}

		case "KILL":
			err := node.VoluntarilyLeavingKeyTransfer()
			if err != nil {
				fmt.Println(err)
			}
			os.Exit(0)
		}
	}
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
