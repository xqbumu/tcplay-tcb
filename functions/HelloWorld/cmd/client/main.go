package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTPPath("tcp", "playground-1251008422.ap-shanghai.app.tcloudbase.com", "/scf/HelloWorld")
	log.Println("a")
	if err != nil {
		log.Fatal("dial error", err)
	}

	log.Println("b")
	var reply string
	err = client.Call("Function", `{"key1":"viking"}`, &reply)
	if err != nil {
		log.Fatal("call error", err)
	}

	log.Println(reply)
}
