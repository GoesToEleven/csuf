package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func client() error {
	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		return err
	}
	defer conn.Close()

	br, bw := bufio.NewReader(conn), bufio.NewWriter(conn)

	msg := []byte("Hello World")
	log.Println("send", string(msg))
	bw.Write(msg)
	bw.WriteByte('\n')
	bw.Flush()

	msg, _, _ = br.ReadLine()
	log.Println("recv", string(msg))

	return nil
}

func server() error {
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		br, bw := bufio.NewReader(conn), bufio.NewWriter(conn)

		msg, _, _ := br.ReadLine()
		log.Println("recv", string(msg))
		log.Println("send", string(msg))
		bw.Write(msg)
		bw.WriteByte('\n')
		bw.Flush()

		// close the connection
		conn.Close()
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("expected `mode`")
	}

	switch os.Args[1] {
	case "client":
		client()
	case "server":
		server()
	}
}
