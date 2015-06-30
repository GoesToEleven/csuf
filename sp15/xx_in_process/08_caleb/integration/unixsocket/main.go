package main

import (
	"log"
	"net/rpc/jsonrpc"
)

type Pair struct {
	X, Y int
}

func main() {
	client, err := jsonrpc.Dial("unix", "/tmp/example.sock")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	for i := 0; i < 5; i++ {
		log.Println("send", "1+4")
		var reply int
		err = client.Call("add", Pair{1, 4}, &reply)
		if err != nil {
			panic(err)
		}
		log.Println("recv", reply)
	}
}
