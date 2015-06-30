package main

//#include <fcntl.h>
import "C"
import (
	"log"
	"os"
)

func open() (mqd_t, error) {
	return mq_open("/example", C.O_CREAT|C.O_RDWR, 0644, mq_attr{
		mq_flags:   0,
		mq_maxmsg:  10,
		mq_msgsize: 1024,
		mq_curmsgs: 0,
	})
}

func recv() {
	q, err := open()
	if err != nil {
		panic(err)
	}
	defer mq_close(q)

	data, _, err := mq_receive(q)
	if err != nil {
		panic(err)
	}
	log.Println("received", string(data))
}

func send() {
	q, err := open()
	if err != nil {
		panic(err)
	}
	defer mq_close(q)

	data := []byte("Hello World")
	log.Println("sending", string(data))

	err = mq_send(q, data, 0)
	if err != nil {
		panic(err)
	}
}

func main() {
	log.SetFlags(0)
	if len(os.Args) < 2 {
		log.Fatalln("expected mode")
	}
	switch os.Args[1] {
	case "send":
		send()
	case "receive":
		recv()
	default:
		log.Fatalln("unknown mode")
	}
}
