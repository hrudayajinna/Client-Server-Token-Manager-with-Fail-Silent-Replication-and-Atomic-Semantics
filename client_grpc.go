package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc"
	pb "hrudaya.com/go-tokenmgmt-grpc/proto"
)

const (
	address = "localhost:50051"
)

func main() {
	args := os.Args[1:]
	if args[0] == "-create" {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		clientConn := pb.NewTknClient(conn)

		clientCntx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := clientConn.CreateNewToken(clientCntx, &pb.NewToken{Name: "undefined", Domain: "undefined", State: "undefined", Id: args[2]})
		if err != nil {
			log.Fatalf("could not create token: %v", err)
		}
		fmt.Println("Token has been created successfully....!!!")
		fmt.Println("Token Details After Write operation:", r)
	}
	if args[0] == "-read" {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		clientConn := pb.NewTknClient(conn)

		clientCntx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := clientConn.GetToekns(clientCntx, &pb.Token{Id: args[2]})
		finVal := strings.Split(strings.Split(r.GetState(), ";")[1], ":")[1]
		fmt.Println("Final Value: ", finVal)
		if err != nil {
			fmt.Println("read err :", err)
		}
		fmt.Println("Token Details After Read operation: ", r)
	}
	if args[0] == "-write" {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		clientConn := pb.NewTknClient(conn)

		clientCntx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		low := args[6]
		mid := args[8]
		high := args[10]
		dom := "low:" + low + ";mid:" + mid + ";high:" + high
		lowInt, _ := strconv.Atoi(args[6])
		midInt, _ := strconv.Atoi(args[8])
		_, partVal := hasher(lowInt, midInt, args[4])
		state := "partialValue:" + partVal + ";finalValue:undefined"
		r, err := clientConn.WriteToken(clientCntx, &pb.NewToken{Name: args[4], Domain: dom, State: state, Id: args[2]})
		if err != nil {
			log.Fatalf("could not create token: %v", err)
		}
		finVal := strings.Split(strings.Split(r.GetState(), ";")[0], ":")[1]
		fmt.Println("Partial Value :", finVal)
		fmt.Println("Token Details : ", r)
	}
	if args[0] == "-drop" {
		conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		clientConn := pb.NewTknClient(conn)

		clientCntx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := clientConn.DropToken(clientCntx, &pb.TokenInfo{Id: args[1]})
		if err != nil {
			log.Fatalf("Unable to drop the token %v", err)
		}
		fmt.Println(r.Message)
	}

}

func hasher(start int, end int, name string) (int, string) {
	hasher := sha256.New()
	var min uint64
	var temp uint64
	var index int
	for i := start; i < end; i++ {
		hasher.Write([]byte(fmt.Sprintf("%s %d", name, i)))
		temp = binary.BigEndian.Uint64(hasher.Sum(nil))
		if i == start {
			min = temp
			index = 0
		}
		if i != start && min > temp {
			min = temp
			index = i
		}
	}

	return index, strconv.Itoa(int(min))
}
