package main

import (
	"log"

	"github.com/Shopify/sysv_mq"
)

func main() {
	log.SetFlags(0)

	mq, err := sysv_mq.NewMessageQueue(&sysv_mq.QueueConfig{
		Key:     1234,
		MaxSize: 8 * 1024,
		Mode:    sysv_mq.IPC_CREAT | 0600,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer mq.Close()
	defer mq.Destroy()
	for {
		msg, typ, err := mq.ReceiveBytes(0, 0)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("received", typ, string(msg))
	}

}
