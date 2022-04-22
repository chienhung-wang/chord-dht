package main

import (
	"chord-dht/chord"
	pb "chord-dht/chord_pb"
	rpc "chord-dht/server"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	"google.golang.org/grpc"
)

var wg sync.WaitGroup

func TestChordNetwork(t *testing.T) {
	const targetNumNode = 20
	const targetNumKey = 1000
	const targetNumGet = 1000
	const isNaive = true
	getFromRandomNode := true
	putToRandomNode := true
	port := 60445
	wg.Add(targetNumNode)

	// define chan
	var chans [targetNumNode]chan string
	for i := range chans {
		chans[i] = make(chan string)
	}
	var taskChan chan string
	var ackChan chan float64
	taskChan = make(chan string, 4096)
	ackChan = make(chan float64, 4096)

	// start goroutine for nodes
	for i := 0; i < targetNumNode; i++ {
		go chordNetWork(isNaive, strconv.Itoa(port+i), chans[i], taskChan, ackChan)
		if <-chans[i] == "fail" {
			continue
		}
		chans[i] <- "JOIN localhost:" + strconv.Itoa(port)
	}

	// wait for stabilize
	time.Sleep(targetNumNode * 2 * time.Second)
	start := time.Now()
	putSum := 0.0
	// add key-value pairs
	if putToRandomNode {
		for i := 0; i < targetNumKey; i++ {
			taskChan <- "PUT " + strconv.Itoa(i) + " " + strconv.Itoa(i)
		}
		for i := 0; i < targetNumKey; i++ {
			putSum += <-ackChan
		}
	} else {
		for i := 0; i < targetNumKey; i++ {
			chans[0] <- "PUT " + strconv.Itoa(i) + " " + strconv.Itoa(i)
		}
		for i := 0; i < targetNumKey; i++ {
			putSum += <-ackChan
		}
	}
	add := time.Now()
	getSum := 0.0
	// get
	if getFromRandomNode {
		for i := 0; i < targetNumGet; i++ {
			key := rand.Intn(targetNumKey)
			taskChan <- "GET " + strconv.Itoa(key)
		}
		for i := 0; i < targetNumGet; i++ {
			getSum += <-ackChan
		}
	} else {
		for i := 0; i < targetNumGet; i++ {
			key := rand.Intn(targetNumKey)
			chans[0] <- "GET " + strconv.Itoa(key)
		}
		for i := 0; i < targetNumGet; i++ {
			getSum += <-ackChan
		}
	}
	get := time.Now()

	// shut down
	for i := 0; i < targetNumNode; i++ {
		taskChan <- "END"
	}

	wg.Wait()

	fmt.Printf("Total time to finish test: %s \n", time.Since(start).String())
	fmt.Printf("Total time to put : %s \n", add.Sub(start))
	fmt.Printf("Total time to get : %s \n", get.Sub(add))
	fmt.Printf("Average time per put: %f \n", putSum/targetNumKey)
	fmt.Printf("Average time per get: %f \n", getSum/targetNumGet)

}

func chordNetWork(isNaive bool, port string, ch chan string, taskChan chan string, ackChan chan float64) {
	defer wg.Done()
	host_port := "localhost:" + port

	storageService := chord.NewStorageService()
	node := chord.NewNode(host_port, storageService)
	node.IsNaive = isNaive
	id := node.Id

	lis, err := net.Listen("tcp", host_port)
	if err != nil {
		ch <- "fail"
		log.Fatalln("Failed toc listen to port", host_port)
	}
	ch <- "success"
	log.Println("Server listening at " + lis.Addr().String())

	s := grpc.NewServer()
	rpcServer := rpc.NewChordNodeServer(storageService, node)

	pb.RegisterChordNodeServer(s, rpcServer)

	log.Println("Server registered...")

	go startServer(s, lis)

	go node.Stabilize()

	go node.UpdateBackupSuccessors()

	go node.FixFinger()

	log.Println("Node id ---> ", id)
	log.Println("Start getting input...")

	//input := <-ch

	for {
		input := ""
		select {
		case input = <-ch:
			fmt.Println("received", input)
		case input = <-taskChan:
			fmt.Println("received taskChan", input)
		}
		log.Println(" **** node " + port + " get input: " + input + "*****")
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
			ackChan <- 1
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
			log.Println("SuccList -> ")
			for _, suc := range node.SuccList {
				if suc == nil {
					break
				}
				fmt.Printf(" -> %v ", suc.Addr[10:])
			}
		case "GET":
			if len(texts) >= 2 {
				start := time.Now()
				if _, _, err := node.Get(texts[1]); err == nil {
					end := time.Now()
					//s := fmt.Sprintf("{Key: %v, Val: %v} -> GET\n", key, val)
					//println(s)
					ackChan <- end.Sub(start).Seconds()
				} else {
					ackChan <- -1
					fmt.Println("Error: ", err)
				}
			}
		case "PUT":
			start := time.Now()
			if len(texts) >= 3 {
				if _, _, err := node.Put(texts[1], texts[2]); err == nil {
					end := time.Now()
					//fmt.Printf("{Key: %v, Val: %v} -> PUT\n", key, val)
					ackChan <- end.Sub(start).Seconds()
				} else {
					//end := time.Now()
					ackChan <- -1
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
		case "END":
			return
		}
		//input = <-taskChan

	}
}
