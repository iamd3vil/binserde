// Code generated by binserde(Build:  (Commit: 2022-03-30 16:58:22 +0530 (4ddc454), Build: 2022-04-19% 12:42:32 +0530)). DO NOT EDIT.
package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

func (s *TestBin) Marshal(wr io.Writer) error {
	buf := make([]byte, 8)
	var (
		byBuf []byte
		err   error
	)
	byBuf = make([]byte, 4)
	copy(byBuf, s.Name)
	_, err = wr.Write(byBuf)
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint32(buf[:4], uint32(s.LenNameBytes))
	_, err = wr.Write(buf[:4])
	if err != nil {
		return err
	}
	byBuf = make([]byte, int(s.LenNameBytes))
	copy(byBuf, s.NameBytes)
	_, err = wr.Write(byBuf)
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint32(buf[:4], uint32(s.Age))
	_, err = wr.Write(buf[:4])
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint64(buf[:8], uint64(s.Age2))
	_, err = wr.Write(buf[:8])
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint16(buf[:2], uint16(s.Age3))
	_, err = wr.Write(buf[:2])
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint64(buf[:8], math.Float64bits(s.Wealth))
	_, err = wr.Write(buf[:8])
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint32(buf[:4], uint32(s.LenEmbedded))
	_, err = wr.Write(buf[:4])
	if err != nil {
		return err
	}
	if err := s.Embedded.Marshal(wr); err != nil {
		return err
	}
	return nil
}

func (s *TestBin) MarshalToBytes() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, s.Size()))
	if err := s.Marshal(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *TestBin) Size() int {
	ln := 0
	ln += 4
	ln += 4
	ln += int(s.LenNameBytes)
	ln += 4
	ln += 8
	ln += 2
	ln += 8
	ln += 4
	ln += s.Embedded.Size()
	return ln
}
func (s *TestBin2) Marshal(wr io.Writer) error {
	buf := make([]byte, 8)
	var (
		byBuf []byte
		err   error
	)
	binary.BigEndian.PutUint32(buf[:4], uint32(s.LenName))
	_, err = wr.Write(buf[:4])
	if err != nil {
		return err
	}
	byBuf = make([]byte, int(s.LenName))
	copy(byBuf, s.Name)
	_, err = wr.Write(byBuf)
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint32(buf[:4], uint32(s.Age))
	_, err = wr.Write(buf[:4])
	if err != nil {
		return err
	}
	if err := s.Metadata.Marshal(wr); err != nil {
		return err
	}
	return nil
}

func (s *TestBin2) MarshalToBytes() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, s.Size()))
	if err := s.Marshal(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *TestBin2) Size() int {
	ln := 0
	ln += 4
	ln += int(s.LenName)
	ln += 4
	ln += s.Metadata.Size()
	return ln
}
func (s *TestStructWithoutStringOrBytes) Marshal(wr io.Writer) error {
	buf := make([]byte, 8)
	var (
		err error
	)
	binary.BigEndian.PutUint32(buf[:4], uint32(s.Qty))
	_, err = wr.Write(buf[:4])
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint32(buf[:4], uint32(s.Price))
	_, err = wr.Write(buf[:4])
	if err != nil {
		return err
	}
	return nil
}

func (s *TestStructWithoutStringOrBytes) MarshalToBytes() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0, s.Size()))
	if err := s.Marshal(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *TestStructWithoutStringOrBytes) Size() int {
	ln := 0
	ln += 4
	ln += 4
	return ln
}
func (s *TestBin) Unmarshal(rdr io.Reader) error {
	buf := make([]byte, 8)
	bufName := make([]byte, 4)
	if _, err := io.ReadFull(rdr, bufName); err != nil {
		return err
	}
	s.Name = string(bufName)
	if _, err := io.ReadFull(rdr, buf[:4]); err != nil {
		return err
	}
	s.LenNameBytes = int32(binary.BigEndian.Uint32(buf[:4]))
	s.NameBytes = make([]byte, int(s.LenNameBytes))
	if _, err := io.ReadFull(rdr, s.NameBytes); err != nil {
		return err
	}
	if _, err := io.ReadFull(rdr, buf[:4]); err != nil {
		return err
	}
	s.Age = int32(binary.BigEndian.Uint32(buf[:4]))
	if _, err := io.ReadFull(rdr, buf[:8]); err != nil {
		return err
	}
	s.Age2 = int64(binary.BigEndian.Uint64(buf[:8]))
	if _, err := io.ReadFull(rdr, buf[:2]); err != nil {
		return err
	}
	s.Age3 = int16(binary.BigEndian.Uint16(buf[:2]))
	if _, err := io.ReadFull(rdr, buf[:8]); err != nil {
		return err
	}
	s.Wealth = math.Float64frombits(binary.BigEndian.Uint64(buf[:8]))
	if _, err := io.ReadFull(rdr, buf[:4]); err != nil {
		return err
	}
	s.LenEmbedded = int32(binary.BigEndian.Uint32(buf[:4]))
	if err := s.Embedded.Unmarshal(rdr); err != nil {
		return err
	}
	return nil
}

func (s *TestBin) UnmarshalFromBytes(data []byte) error {
	buf := bytes.NewBuffer(data)
	return s.Unmarshal(buf)
}
func (s *TestBin2) Unmarshal(rdr io.Reader) error {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(rdr, buf[:4]); err != nil {
		return err
	}
	s.LenName = int32(binary.BigEndian.Uint32(buf[:4]))
	s.Name = make([]byte, int(s.LenName))
	if _, err := io.ReadFull(rdr, s.Name); err != nil {
		return err
	}
	if _, err := io.ReadFull(rdr, buf[:4]); err != nil {
		return err
	}
	s.Age = int32(binary.BigEndian.Uint32(buf[:4]))
	if err := s.Metadata.Unmarshal(rdr); err != nil {
		return err
	}
	return nil
}

func (s *TestBin2) UnmarshalFromBytes(data []byte) error {
	buf := bytes.NewBuffer(data)
	return s.Unmarshal(buf)
}
func (s *TestStructWithoutStringOrBytes) Unmarshal(rdr io.Reader) error {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(rdr, buf[:4]); err != nil {
		return err
	}
	s.Qty = int32(binary.BigEndian.Uint32(buf[:4]))
	if _, err := io.ReadFull(rdr, buf[:4]); err != nil {
		return err
	}
	s.Price = int32(binary.BigEndian.Uint32(buf[:4]))
	return nil
}

func (s *TestStructWithoutStringOrBytes) UnmarshalFromBytes(data []byte) error {
	buf := bytes.NewBuffer(data)
	return s.Unmarshal(buf)
}
