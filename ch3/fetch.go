package ch3

import "fmt"

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func Zero(ptr [32]byte) {
	for i := range ptr {
		fmt.Println(i)
		ptr[i] = 0
	}
}

func One(ptr []byte) {
	for i := range ptr {
		fmt.Println(i)
		ptr[i] = 1
	}
}
