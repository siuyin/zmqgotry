package main

import (
	"fmt"
	"log"
	"os"
	"time"

	zmq "github.com/pebbe/zmq2"
)

const tickInterval time.Duration = 6000 * time.Millisecond

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

	tk := time.Tick(tickInterval)
	for {
		select {
		case t := <-tk:
			ts := t.String()
			_, err := push.Send(ts, 0)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(ts)
		}
	}
}
