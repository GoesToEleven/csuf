#!/bin/bash
gcc -std=c99 -c sum.c -o /tmp/sum.o
gccgo -g -c main.go -o /tmp/main.o
gccgo /tmp/sum.o /tmp/main.o -o /tmp/sum
/tmp/sum
