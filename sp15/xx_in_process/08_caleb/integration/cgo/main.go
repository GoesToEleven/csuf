package main

/*
long long sum(long long *xs, long long sz) {
  long long total = 0;
  long long i;
  for (i=0; i < sz; i++) {
    total += xs[i];
  }
  return total;
}
*/
import "C"
import (
	"fmt"
	"math/rand"
	"unsafe"
)

func main() {
	xs := make([]int64, 4096)
	for i := 0; i < len(xs); i++ {
		xs[i] = int64(rand.Int())
	}
	fmt.Println(C.sum((*C.longlong)(unsafe.Pointer(&xs[0])), C.longlong(len(xs))))
}
