package main

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"

	"google.golang.org/grpc"
	pb "hrudaya.com/go-tokenmgmt-grpc/proto"
)

const (
	port = ":50051"
)

type TknServer struct {
	pb.UnimplementedTknServer
	token_list *pb.TokenList
	mutex      sync.RWMutex
}

func (tkn_server *TknServer) CreateNewToken(serverCntx context.Context, in *pb.NewToken) (*pb.Token, error) {
	tkn_server.mutex.Lock()
	created_token := &pb.Token{Name: in.GetName(), Domain: in.GetDomain(), Id: in.GetId(), State: in.GetState()}
	tkn_server.token_list.Tokens = append(tkn_server.token_list.Tokens, created_token)
	tkn_server.mutex.Unlock()
	log.Printf("Created Token: %v", in.GetId())
	return created_token, nil
}

func (tkn_server *TknServer) GetToekns(serverCntx context.Context, in *pb.Token) (*pb.Token, error) {
	for i := 0; i < len(tkn_server.token_list.Tokens); i++ {
		if tkn_server.token_list.Tokens[i].Id == in.GetId() {
			if tkn_server.token_list.Tokens[i].Domain != "undefined" {
				dom := tkn_server.token_list.Tokens[i].Domain
				state := tkn_server.token_list.Tokens[i].State
				if strings.Split(strings.Split(state, ";")[1], ":")[1] != "undefined" {
					return tkn_server.token_list.Tokens[i], nil
				}
				mid := strings.Split(strings.Split(dom, ";")[1], ":")[1]
				high := strings.Split(strings.Split(dom, ";")[2], ":")[1]
				partVal := strings.Split(strings.Split(state, ";")[0], ":")[1]
				midInt, _ := strconv.Atoi(mid)
				highInt, _ := strconv.Atoi(high)
				_, hashVal := hasher(midInt, highInt, in.GetName())
				parInt, _ := strconv.Atoi(partVal)
				hashInt, _ := strconv.Atoi(hashVal)
				var finalSta string
				if parInt < hashInt {
					finalSta = "PartialValue:" + strconv.Itoa(parInt) + ";FinalValue:" + strconv.Itoa(parInt)
				} else {
					finalSta = "PartialValue:" + strconv.Itoa(parInt) + ";FinalValue:" + strconv.Itoa(hashInt)
				}
				tkn_server.token_list.Tokens[i].State = finalSta
				idList := ""
				for i := 0; i < len(tkn_server.token_list.Tokens); i++ {
					idList = idList + " " + tkn_server.token_list.Tokens[i].Id
				}
				fmt.Println("Current Token Id's after reading ", in.GetId()+" :", idList)
				return tkn_server.token_list.Tokens[i], nil
			}
			idList := ""
			for i := 0; i < len(tkn_server.token_list.Tokens); i++ {
				idList = idList + " " + tkn_server.token_list.Tokens[i].Id
			}
			fmt.Println("Current Token Id's after reading ", in.GetId()+" :", idList)
			return tkn_server.token_list.Tokens[i], nil
		}
	}
	return nil, fmt.Errorf("Token doesn't exit. Enter valid Token Id")
}

func (tkn_server *TknServer) DropToken(serverCntx context.Context, in *pb.TokenInfo) (*pb.EmptyToken, error) {
	tkn_server.mutex.Lock()
	for i := 0; i < len(tkn_server.token_list.Tokens); i++ {
		if tkn_server.token_list.Tokens[i].Id == in.GetId() {
			tkn_server.token_list.Tokens = append(tkn_server.token_list.Tokens[:i], tkn_server.token_list.Tokens[i+1:]...)
			break
		}
	}
	tkn_server.mutex.Unlock()
	mess := "Token " + in.GetId() + " has been dropped successfully...!!!!"
	fmt.Println(mess)
	idList := ""
	for i := 0; i < len(tkn_server.token_list.Tokens); i++ {
		idList = idList + " " + tkn_server.token_list.Tokens[i].Id
	}
	if idList == "" {
		idList = "Empty"
	}
	fmt.Println("Current Token Id's after dropping ", in.GetId()+" :", idList)
	return &pb.EmptyToken{Message: mess}, nil
}

func (tkn_server *TknServer) WriteToken(serverCntx context.Context, in *pb.NewToken) (*pb.Token, error) {
	tkn_server.mutex.Lock()
	var index int
	var flag bool
	for i := 0; i < len(tkn_server.token_list.Tokens); i++ {
		if tkn_server.token_list.Tokens[i].Id == in.GetId() {
			tkn_server.token_list.Tokens[i].Name = in.GetName()
			tkn_server.token_list.Tokens[i].State = in.GetState()
			tkn_server.token_list.Tokens[i].Domain = in.GetDomain()
			index = i
			flag = true
			break
		}
	}
	tkn_server.mutex.Unlock()
	if flag != true {
		log.Fatal("Token Not Found! Please enter correct ID.")
		return nil, nil
	}
	idList := ""
	for i := 0; i < len(tkn_server.token_list.Tokens); i++ {
		idList = idList + " " + tkn_server.token_list.Tokens[i].Id
	}
	fmt.Println("Current Token Id's after writing", in.GetId()+" :", idList)
	return tkn_server.token_list.Tokens[index], nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTknServer(s, &TknServer{token_list: &pb.TokenList{}})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start server at: %v", err)
	}
}

func hasher(start int, end int, name string) (int, string) {
	hasher := sha256.New()
	var min uint64
	var temp uint64
	var index int
	for i := start; i <= end; i++ {
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
