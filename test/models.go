package main

type TestBin struct {
	Name         string `bin:"len=4"`
	LenNameBytes int32
	NameBytes    []byte `bin:"len=LenNameBytes"`

	Age         int32
	Age2        int64
	Age3        int16
	LenEmbedded int32
	Embedded    TestBin2 `bin:"len=LenEmbedded"`
}

type TestBin2 struct {
	LenName int32
	Name    []byte `bin:"len=LenName"`
	Age     int32
}

type BCastHeader struct {
	Reserved1  string `bin:"len=4"`
	LogTime    int32
	AlphaChar  string `bin:"len=2"`
	TransCode  int16
	ErrorCode  int16
	BCSeqNo    int32
	Reserved2  string `bin:"len=4"`
	Timestamp2 string `bin:"len=8"`
	Filler2    string `bin:"len=8"`
}

type BCastPacket struct {
	Header        BCastHeader `bin:"len=38"`
	MessageLength int16
	Packet        []byte `bin:"len=MessageLength"`
}
