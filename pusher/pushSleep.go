package main

import (
	"fmt"
	"log"
	"os"
	"time"

	zmq "github.com/pebbe/zmq2"
)

//const tickInterval time.Duration = 1000 * time.Millisecond
const tickInterval time.Duration = 1 * time.Second

func main() {
	push, err := zmq.NewSocket(zmq.PUSH)
	if err != nil {
		log.Fatal(err)
	}
	defer push.Close()

	err = push.Bind(os.Getenv("PUSH_BIND_ADDR"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting pusher ...")

	for {
		ts := time.Now().String()
		_, err := push.Send(ts, 0)
		if err != nil {
			log.Fatal("push:", err)
		}
		time.Sleep(tickInterval)
		fmt.Println(ts)
	}
}
