package main

import (
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
	bs := []byte{0}
	for {
		_, err := os.Stdin.Read(bs)
		if err != nil {
			break
		}
		bs[0] = rot13(bs[0])
		_, err = os.Stdout.Write(bs)
		if err != nil {
			break
		}
	}
}
