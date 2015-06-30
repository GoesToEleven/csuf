package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"unsafe"

	"github.com/Shopify/sysv_mq"
	"github.com/edsrzf/mmap-go"
)

func send(mq *sysv_mq.MessageQueue, offset, size int) error {
	var data [16]byte
	binary.BigEndian.PutUint64(data[:], uint64(offset))
	binary.BigEndian.PutUint64(data[8:], uint64(size))
	return mq.SendBytes(data[:], 1, 0)
}

func recv(mq *sysv_mq.MessageQueue) (offset, size int, err error) {
	data, _, e := mq.ReceiveBytes(1, 0)
	if err != nil {
		err = e
		return
	}
	if len(data) < 16 {
		err = fmt.Errorf("expected offset and size")
		return
	}
	offset = int(binary.BigEndian.Uint64(data[:8]))
	size = int(binary.BigEndian.Uint64(data[8:]))
	return
}

func server(mq *sysv_mq.MessageQueue) {
	fd, err := shm_open("/shm-example", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	defer shm_unlink("/shm-example")

	err = fd.Truncate(50 * 1024 * 1024)
	if err != nil {
		panic(err)
	}

	data, err := mmap.Map(fd, mmap.RDWR|mmap.EXEC, 0)
	if err != nil {
		panic(err)
	}

	for {
		_, sz, err := recv(mq)
		if err != nil {
			panic(err)
		}
		arr := mkArr(data, sz)
		log.Println("received", len(arr), "doubles")
	}
}

func client(mq *sysv_mq.MessageQueue) {
	fd, err := shm_open("/shm-example", os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	data, err := mmap.Map(fd, mmap.RDWR|mmap.EXEC, 0)
	if err != nil {
		panic(err)
	}

	sz := 1024 * 1024 * 50 / 8
	arr := mkArr(data, sz)
	for i := 0; i < sz; i++ {
		arr[i] = rand.Float64()
	}
	log.Println("sending", sz, "doubles")
	send(mq, 0, sz)
}

func mkArr(data []byte, sz int) []float64 {
	// in general:
	//
	// (1) grab the address of the first element of the byte slice
	//     all the subsequent elements will be adjacent
	// (2) convert that address into an unsafe pointer
	// (3) convert the unsafe pointer into an array pointer. We tell
	//     the compiler that the array is as large as possible
	// (4) convert the array into a slice and give it a fixed length
	//     and capacity
	return (*[(1 << 31) - 1]float64)(unsafe.Pointer(&data[0]))[:sz:sz]
}

func main() {
	log.SetFlags(0)

	if len(os.Args) < 2 {
		log.Fatalln("expected mode")
	}

	mq, err := sysv_mq.NewMessageQueue(&sysv_mq.QueueConfig{
		Key:     1001,
		MaxSize: 8 * 1024,
		Mode:    sysv_mq.IPC_CREAT | 0600,
	})
	if err != nil {
		log.Fatalln(err)
	}
	defer mq.Close()

	switch os.Args[1] {
	case "server":
		server(mq)
	case "client":
		client(mq)
	}
}
