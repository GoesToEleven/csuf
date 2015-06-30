package main

/*
#cgo LDFLAGS: -lrt
#include <stdlib.h>
#include <fcntl.h>
#include <mqueue.h>
#include <errno.h>

mqd_t _mq_open(const char *name, int oflag, mode_t mode, struct mq_attr *attr) {
  return mq_open(name, oflag, mode, attr);
}
*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

type (
	mqd_t   C.mqd_t
	mq_attr struct {
		mq_flags, mq_maxmsg, mq_msgsize, mq_curmsgs int
	}
)

func mq_open(name string, oflag int, mode int, attr mq_attr) (mqd_t, error) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	var cattr C.struct_mq_attr
	cattr.mq_flags = C.__syscall_slong_t(attr.mq_flags)
	cattr.mq_maxmsg = C.__syscall_slong_t(attr.mq_maxmsg)
	cattr.mq_msgsize = C.__syscall_slong_t(attr.mq_msgsize)
	cattr.mq_curmsgs = C.__syscall_slong_t(attr.mq_curmsgs)
	q, err := C._mq_open(cname, C.int(oflag), C.mode_t(mode), &cattr)
	if err != nil {
		log.Printf("%d\n", err)
		return mqd_t(q), err
	}
	return mqd_t(q), nil
}

func mq_getattr(q mqd_t) (mq_attr, error) {
	var attr mq_attr
	var cattr C.struct_mq_attr
	res, err := C.mq_getattr(C.mqd_t(q), &cattr)
	if res < 0 {
		return attr, err
	}
	attr.mq_flags = int(cattr.mq_flags)
	attr.mq_maxmsg = int(cattr.mq_maxmsg)
	attr.mq_msgsize = int(cattr.mq_msgsize)
	attr.mq_curmsgs = int(cattr.mq_curmsgs)
	return attr, nil
}

func mq_close(q mqd_t) {
	C.mq_close(C.mqd_t(q))
}

func mq_send(q mqd_t, data []byte, priority int) error {
	errno := C.mq_send(C.mqd_t(q), (*C.char)(unsafe.Pointer(&data[0])), C.size_t(len(data)), C.uint(priority))
	if errno < 0 {
		return fmt.Errorf("some error")
	}
	return nil
}

func mq_receive(q mqd_t) ([]byte, int, error) {
	attr, err := mq_getattr(q)
	if err != nil {
		return nil, 0, err
	}
	data := make([]byte, attr.mq_msgsize)
	var priority C.uint
	sz, err := C.mq_receive(C.mqd_t(q), (*C.char)(unsafe.Pointer(&data[0])), C.size_t(attr.mq_msgsize), &priority)
	if sz < 0 {
		return nil, 0, err
	}
	return data[:sz], int(priority), nil
}
