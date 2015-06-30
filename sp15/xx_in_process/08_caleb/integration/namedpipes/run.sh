#!/bin/bash
rm -rf /tmp/rot13.fifo

# create our fifo queue
mkfifo /tmp/rot13.fifo

# start reading from it
cat /tmp/rot13.fifo &

# start writing to it (- means stdin / stdout)
go run rot13/main.go - /tmp/rot13.fifo

# you can also do:
#
# go run rot13/main.go > /tmp/rot13.fifo
#
