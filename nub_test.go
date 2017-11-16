package nub

import (
	"fmt"
	"testing"
)

func TestGetTLS(t *testing.T) {
	tls := GetTLS()
	if tls != 0 {
		t.Fatal()
	}
	SetTLS(100)
	tls = GetTLS()
	if tls != 100 {
		t.Fatal()
	}
}

func TestEcho(t *testing.T) {
	EchoCount = 0
	Ohce()
	if EchoCount != 10 {
		t.Fatal()
	}
}

func BenchmarkEcho(b *testing.B) {
	for _, gn := range []bool{false, true} {
		b.Run(fmt.Sprintf("%v", gn), func(b *testing.B) {
			if gn {
				orig := cecho
				defer func() { cecho = orig }()
				cecho = func() {
					Ohce()
				}
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				EchoCount = 0
				Ohce()
				if EchoCount != 10 {
					b.Fatal()
				}
			}
		})
	}
}
