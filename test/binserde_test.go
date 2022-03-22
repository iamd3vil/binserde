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
		Age:          10,
		Age2:         123456789,
		Age3:         2,
		Wealth:       450.3,
		LenEmbedded:  12,
		Embedded: TestBin2{
			LenName:  5,
			Name:     []byte("Sarat"),
			Age:      29,
			Metadata: CustomString("hello"),
		},
	}
	// Testing if the length is correct.
	m.NameBytes = make([]byte, m.LenNameBytes)
	copy(m.NameBytes, []byte("Ric"))
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
		Name:         "Rick",
		LenNameBytes: 4,
		Age:          10,
		Age2:         123456789,
		Age3:         2,
		Wealth:       450.3,
		LenEmbedded:  12,
		Embedded: TestBin2{
			LenName:  5,
			Name:     []byte("Sarat"),
			Age:      29,
			Metadata: CustomString("hello"),
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
