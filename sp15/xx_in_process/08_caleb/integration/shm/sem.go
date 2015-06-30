package main

/*
#cgo LDFLAGS: -pthread
#include <semaphore.h>
#include <stdlib.h>
#include <fcntl.h>

sem_t *sem_open_(const char *name, int oflag, mode_t mode, unsigned int value) {
  return sem_open(name, oflag, mode, value);
}
*/
import "C"
import "unsafe"

type (
	sem_t C.sem_t
)

func sem_open(name string, oflag, mode, value int) (*sem_t, error) {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	c_oflag := C.int(oflag)
	c_mode := C.mode_t(mode)
	c_value := C.uint(value)
	r, err := C.sem_open_(c_name, c_oflag, c_mode, c_value)
	return (*sem_t)(r), err
}

func sem_close(sem *sem_t) error {
	_, err := C.sem_close((*C.sem_t)(sem))
	return err
}

func sem_wait(sem *sem_t) error {
	_, err := C.sem_wait((*C.sem_t)(sem))
	return err
}

func sem_post(sem *sem_t) error {
	_, err := C.sem_post((*C.sem_t)(sem))
	return err
}

func sem_unlink(name string) error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))
	_, err := C.sem_unlink(c_name)
	return err
}
