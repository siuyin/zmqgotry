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
	plr := zmq.NewPoller()
	plr.Add(pull, zmq.POLLIN)
	for {
		socks, err := plr.Poll(3 * time.Second)
		if err != nil {
			log.Fatal("poll:", err)
		}
		for _, sk := range socks {
			switch so := sk.Socket; so {
			case pull:
				s, err := so.Recv(0)
				if err != nil {
					log.Fatal("recv:", err)
				}
				fmt.Println(s)
			}
		}

		if len(socks) == 0 {
			fmt.Println("timeout")
		}

	}
}
