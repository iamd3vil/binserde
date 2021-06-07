package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	m := TestBin{
		Name:      "Rick",
		NameBytes: []byte("Rick"),
		Age:       10,
		Age2:      123456789,
		Age3:      2,
		Embedded: TestBin2{
			Age:  28,
			Name: []byte("Rick"),
		},
	}
	b := m.Marshal()

	m2 := TestBin{}
	m2.Unmarshal(b)

	if !reflect.DeepEqual(m, m2) {
		t.Fatalf("Expected: %#+v, got: %#+v", m, m2)
	}
}

func BenchmarkMarshalUnmarshal(b *testing.B) {
	m := TestBin{
		Name:      "Rick",
		NameBytes: []byte("Rick"),
		Age:       10,
		Age2:      123456789,
		Age3:      2,
	}

	bs := m.Marshal()

	b.Run("BenchmarkMarshal", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i <= b.N; i++ {
			m.Marshal()
		}
	})

	b.Run("BenchmarkUnmarshal", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i <= b.N; i++ {
			m.Unmarshal(bs)
		}
	})

	b.Run("BenchmarkJsonMarshal", func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i <= b.N; i++ {
			json.Marshal(m)
		}
	})
}
