/*
Create a program(s) that uses each of the following packages - it is complete ACCEPTABLE
if your program is exactly the same as the program and code detailed in chapter 13:
strings
os
io/ioutil
path/filepath
errors
container/list
sort
hash/crc32
crypto/sha1
encoding/gob
net/http
net/rpc
math/rand
flag
sync
time

 */

package main

import (
	"fmt"
//	"reflect"
//	"runtime"
	"strings"
	"os"
	"io/ioutil"
	"path/filepath"
	"errors"
	"container/list"
	"math/rand"
	"sort"
	"hash/crc32"
	"crypto/sha1"
	"encoding/gob"
	"net"
	"net/http"
	"net/rpc"
	"io"
	"flag"
	"sync"
	"time"
)

func main() {

	//strings
	testStrings()

	//os,io/ioutil,path/filepath
	testOsPackage()

	//error
	err := testError()
	if err != nil {fmt.Println(err)}

	//contairs, lists, sort
	testContainers()

	//hashes
	testHashes()

	//testing net
	testNet()

	//test flag
	testFlag()

	//Test sync and time
	testSync()

	//end
	var input string
	fmt.Scanln(&input)
}

func testSync() {
	fmt.Println("\n\n")
	fmt.Println("****************************************************")
	fmt.Println("Sync, time")
	fmt.Println("****************************************************")

	m := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println(i, "start")
			time.Sleep(time.Second)
			fmt.Println(i, "end")
			m.Unlock()
		}(i)
	}
}

func testFlag() {
	fmt.Println("\n\n")
	fmt.Println("****************************************************")
	fmt.Println("Flag")
	fmt.Println("****************************************************")

	// Define flags
	maxp := flag.Int("max", 6, "the max value")
	// Parse
	flag.Parse()

	if *maxp == 6 {
		fmt.Println("No arugments passed using default: 6")
	}

	// Generate a number between 0 and max
	fmt.Println("Random number is:",rand.Intn(*maxp))
}

func testNet(){
	fmt.Println("\n\n")
	fmt.Println("****************************************************")
	fmt.Println("net stuff active")
	fmt.Println("****************************************************")
	fmt.Println("Press enter to exit...")


	go server()
	go client()

	go http.HandleFunc("/hello", hello)
	go http.ListenAndServe(":9000", nil)

	go server2()
	go client2()
}

type Server struct {}
func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func server2() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9998")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}
func client2() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9998")
	if err != nil {
		fmt.Println(err)
		return
	}
	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) =", result)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
	<head>
		<title>Hello World</title>
	</head>
	<body>
		Hello Go World!
	</body>
</html>`,
	)
}

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

func testHashes() {
	fmt.Println("\n\n")
	fmt.Println("****************************************************")
	fmt.Println("Hashes")
	fmt.Println("****************************************************")

	hashString := "lorem ipsum"
	fmt.Println("test string:", hashString)

	h := crc32.NewIEEE()
	h.Write([]byte(hashString))
	fmt.Println("crc32:", h.Sum32())

	ch := sha1.New()
	ch.Write([]byte(hashString))
	fmt.Println("sha1:", ch.Sum([]byte{}))
}

type Person struct {
Name string
Age int
}

type ByAge []Person
func (this ByAge) Len() int {
	return len(this)
}
func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age
}
func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type ByName []Person
func (this ByName) Len() int {
	return len(this)
}
func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}
func (this ByName) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func testContainers() {
	fmt.Println("\n\n")
	fmt.Println("****************************************************")
	fmt.Println("Containers, lists, sorts")
	fmt.Println("****************************************************")

	var x list.List
	for i := 0 ; i < 10 ; i++ {
		x.PushBack(rand.Int()%20)
	}
	fmt.Println("A list")
	for e := x.Front(); e != nil; e=e.Next() {
		fmt.Println(e.Value.(int))
	}

	//Sort
	kids := []Person{
		{"Kara", 2},
		{"Bethany", 1 },
		{"Zach", 3},
	}

	fmt.Println("People:", kids)
	sort.Sort(ByName(kids))
	fmt.Println("Sorted People by Name:", kids)
	sort.Sort(ByAge(kids))
	fmt.Println("Sorted People by Age", kids)

}

//Test errors
func testError() error {
	fmt.Println("\n\n")
	fmt.Println("****************************************************")
	fmt.Println("Errors Package")
	fmt.Println("****************************************************")
	err := errors.New("This is your new error")
	return err
}

//Test the os/ioutil/path/filepath
func testOsPackage() {
	fmt.Println("\n\n")
	fmt.Println("****************************************************")
	fmt.Println("OS,io/util,path/filepath Packages")
	fmt.Println("****************************************************")

	//write file
	file, err := os.Create("test.txt")
	if err != nil{ file.Close(); fmt.Println("file error"); return}
	file.WriteString("Lorem ipsum")
	fmt.Println("Wrote file test.txt with content 'Lorem Ipsum'")
	file.Close()


	//read file
	bs, err := ioutil.ReadFile("test.txt")
	if err != nil {fmt.Println("file error"); return}
	fmt.Println("Opened and Read test.txt:", string(bs))

	//directory path
	dir, err := os.Open(".")
	if err != nil { fmt.Println("directory error"); return}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil { fmt.Println("directory error"); return}
	fmt.Println("Directory Contents:")
	for _, fi := range fileInfos {
		fmt.Println("-",fi.Name())
	}

	//walk a directory
	fmt.Println("walking this directory")
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			fmt.Println(path)
			return nil
		})
}

//Test the string package
func testStrings() {

	fmt.Println("\n\n")
	fmt.Println("****************************************************")
	fmt.Println("Strings Package")
	fmt.Println("****************************************************")

	myString := "This is a string"
	fmt.Println("The test string is: ", myString)
	//	fmt.Println("Testing", runtime.FuncForPC(reflect.ValueOf(strings.Contains).Pointer()).Name())
	//	if strings.Contains(myString, "string") {
	//		fmt.Println("contains!")
	//	}

	fmt.Println("strings.Contains, does the test string contain 'string'? ", strings.Contains(myString, "string"))
	fmt.Println("strings.Count, how many 'i' are there?", strings.Count(myString, "i"))
	fmt.Println("strings.HasPrefix, does it start with 'Thi'?", strings.HasPrefix(myString, "Thi"))
	fmt.Println("strings.HasSuffix, does it end with 'ring'?", strings.HasSuffix(myString, "ring"))
	fmt.Println("strings.Index, where is does 'tr' start?", strings.Index(myString, "tr"))
	fmt.Println("strings.Join, joint with 'Yeehaw!':", strings.Join([]string{myString, "Yeehaw!"}, " "))
	fmt.Println("strings.Repeat, repeat a, 5 times:", strings.Repeat("a", 5))
	fmt.Println("strings.Replace, replace 'string' with 'bananana':", strings.Replace(myString, "string", "bananana",1))
	fmt.Println("strings.Split, split 'a-b-c-d-e-f' on '-':", strings.Split("a-b-c-d-e-f", "-"))
	fmt.Println("strings.ToLower:", strings.ToLower(myString))
	fmt.Println("strings.ToUpper:", strings.ToUpper(myString))
	fmt.Println("change the string to bytes:", []byte(myString))
	fmt.Println("change bytes to a string [a b c d e]:", string([]byte{'a', 'b', 'c', 'd', 'e'}))
}
