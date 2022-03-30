package main

import "io"

type CustomString string

func (c *CustomString) Marshal(wr io.Writer) error {
	if _, err := io.WriteString(wr, string(*c)); err != nil {
		return err
	}
	return nil
}

func (c *CustomString) Unmarshal(rdr io.Reader) error {
	b := make([]byte, c.Size())
	if _, err := io.ReadFull(rdr, b); err != nil {
		return err
	}

	*c = CustomString(string(b))

	return nil
}

func (c *CustomString) Size() int {
	return 5
}

type TestBin struct {
	Name         string `bin:"len=4"`
	LenNameBytes int32
	NameBytes    []byte `bin:"len=LenNameBytes"`

	Age         int32
	Age2        int64
	Age3        int16
	Wealth      float64
	LenEmbedded int32
	Embedded    TestBin2
}

type TestBin2 struct {
	LenName int32
	Name    []byte `bin:"len=LenName"`
	Age     int32

	// Add custom type.
	Metadata CustomString
}

type TestStructWithoutStringOrBytes struct {
	Qty   int32
	Price int32
}
