package main

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestMarshalUnmarshal(t *testing.T) {
	m := TestBin{
		Name:         "Rick",
		LenNameBytes: 4,
		NameBytes:    []byte("Rick"),
		Age:          10,
		Age2:         123456789,
		Age3:         2,
		LenEmbedded:  12,
		Embedded: TestBin2{
			Age:     28,
			LenName: int32(len([]byte("Rick"))),
			Name:    []byte("Rick"),
		},
	}
	t.Run("io-reader-writer", func(t *testing.T) {
		b := bytes.NewBuffer([]byte{})
		if err := m.Marshal(b); err != nil {
			t.Fatal(err)
		}

		m2 := TestBin{}
		m2.Unmarshal(b)

		if !reflect.DeepEqual(m, m2) {
			t.Fatalf("Expected: %#+v, got: %#+v", m, m2)
		}
	})

	t.Run("bytes", func(t *testing.T) {
		b, err := m.MarshalToBytes()
		if err != nil {
			t.Fatal(err)
		}

		m2 := TestBin{}
		m2.UnmarshalFromBytes(b)

		if !reflect.DeepEqual(m, m2) {
			t.Fatalf("Expected: %#+v, got: %#+v", m, m2)
		}
	})
}

func BenchmarkMarshalUnmarshal(b *testing.B) {
	m := TestBin{
		// Name:         "Rick",
		// LenNameBytes: 4,
		// NameBytes:    []byte("Rick"),
		// Age:          10,
		// Age2:         123456789,
		// Age3:         2,
		Name:         "Rick",
		LenNameBytes: 4,
		NameBytes:    []byte("Rick"),
		Age:          10,
		Age2:         123456789,
		Age3:         2,
		LenEmbedded:  12,
		Embedded: TestBin2{
			Age:     28,
			LenName: int32(len([]byte("Rick"))),
			Name:    []byte("Rick"),
		},
	}

	bs := bytes.NewBuffer([]byte{})

	b.Run("BenchmarkMarshal", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i <= b.N; i++ {
			bs.Reset()
			m.Marshal(bs)
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
