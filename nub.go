// Package gonetmap is a wrapper around the netmap library.
// +build freebsd linux darwin
package nub

/*
#include "nub.h"
*/
import "C"

var (
	EchoCount = 0
	cecho     = func() { (C.echo()) }
)

func SetTLS(tls int) {
	C.setTLS(C.int(tls))
}

func GetTLS() (tls int) {
	return int(C.getTLS())
}

//export Ohce
func Ohce() {
	EchoCount++
	if EchoCount == 10 {
		return
	}
	cecho()
}
