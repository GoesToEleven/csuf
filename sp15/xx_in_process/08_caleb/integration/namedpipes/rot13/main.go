package main

import (
	"io"
	"os"
	"os/signal"
)

func init() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			os.Exit(0)
		}
	}()
}

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

func main() {
	var in io.Reader = os.Stdin
	var out io.Writer = os.Stdout

	if len(os.Args) > 1 && os.Args[1] != "-" {
		// open the input file for reading
		f, err := os.Open(os.Args[1])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		in = f
	}

	if len(os.Args) > 2 && os.Args[2] != "-" {
		// open the output file for writing
		f, err := os.Create(os.Args[2])
		if err != nil {
			panic(err)
		}
		defer f.Close()
		out = f
	}

	// from here on is the same
	bs := []byte{0}
	for {
		_, err := in.Read(bs)
		if err != nil {
			break
		}
		bs[0] = rot13(bs[0])
		_, err = out.Write(bs)
		if err != nil {
			break
		}
	}
}
