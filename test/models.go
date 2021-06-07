package main

type TestBin struct {
	Name      string `bin:"len=4"`
	NameBytes []byte `bin:"len=4"`
	Age       int32
	Age2      int64
	Age3      int16
	Embedded  TestBin2 `bin:"len=8"`
}

type TestBin2 struct {
	Name []byte `bin:"len=4"`
	Age  int32
}
