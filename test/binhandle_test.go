package main

import (
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	m := TestBin{
		Name:      "Rick",
		NameBytes: []byte("Rick"),
		Age:       10,
		Age2:      123456789,
	}
	b := m.Marshal()
	fmt.Printf("Bytes: %v\n", b)
}
