package main

import (
	"fmt"
	"log"
	"os"
	"time"

	zmq "github.com/pebbe/zmq2"
)

const timeoutInterval time.Duration = 1000 * time.Millisecond

func main() {
	if os.Getenv("PULL_CONNECT_ADDR") == "" {
		log.Fatal("Please set env: PULL_CONNECT_ADDR")
	}
	pull, err := zmq.NewSocket(zmq.PULL)
	if err != nil {
		log.Fatal("socket:", err)
	}
	defer pull.Close()
	err = pull.Connect(os.Getenv("PULL_CONNECT_ADDR"))
	if err != nil {
		log.Fatal("connect:", err)
	}

	fmt.Println("Starting puller ...")

	for {
		s, err := pull.Recv(0)
		if err != nil {
			log.Fatal("recv:", err)
		}
		fmt.Println(s)
	}
}
