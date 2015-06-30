/*

First off, in General Networking, there are 2 common types of handling connections.
This can be done either over TCP (Transmission Control Protocol) or UDP (User Datagram Protocol).
The most import difference between these two is that UDP will continuously send out streams/buffers
of bytes without checking to see if the network packets made it to the other side of the line.
This is useful in situations where security isn't much of an issue and where speed is important.
Most VoIP services (Skype, Hangouts), XMPP (chat) and even YouTube (I think) use UDP for their
streaming, as it has huge gains on performances and it doesn't matter all that much if a frame
made it to the other side of the line, as the person could simply repeat themselves.

TCP on the other hand, is "secure" by default. It performs several handshakes on a regular basis
with the endpoint so maintain connectivity and make sure that all packets are received on the other
side of the line.

Now, there are a LOT of protocols out there in the Wild Wild West called Internet.
List of TCP and UDP port numbers

As you can see, a lot of protocols support either TCP or UDP. HTTP on it's own is a TCP protocol
with port 80 (as you might know). Therefore, HTTPServer is pretty much just an extension of a
TCPServer, but with some add-ons such as REST. These add-ons are much welcome as HTTP processing
is a pretty common use case. Without HTTPServer, you would need to declare loads of functions on
your own.

http://stackoverflow.com/questions/23444308/tcp-server-vs-http-server-in-vert-x

*/

package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// listen on a port
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// accept a connection
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}

	c.Close()
}

func client() {
	// connect to the server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// send the message
	msg := "Hello World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}

/*
This example uses the encoding/gob package which makes it easy to encode Go values
so that other Go programs (or the same Go program in this case) can read them.
Additional encodings are available in packages underneath encoding (like encoding/json)
as well as in 3rd party packages.
*/
