// Code generated by binserde(Build:  (Commit: 2021-06-09 15:06:07 +0530 (1ec8c93), Build: 2022-03-08% 12:29:12 +0530)). DO NOT EDIT.
package main

import (
	"encoding/binary"
	"io"
)

func (s *TestBin) Marshal(wr io.Writer) error {
	buf := make([]byte, 8)
	var (
		byBuf []byte
		err   error
	)
	byBuf = make([]byte, len(s.Name))
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
	byBuf = make([]byte, len(s.NameBytes))
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

func (s *TestBin) Size() int {
	ln := 0
	ln += 4
	ln += 4
	ln += int(s.LenNameBytes)
	ln += 4
	ln += 8
	ln += 2
	ln += 4
	ln += int(s.LenEmbedded)
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
	byBuf = make([]byte, len(s.Name))
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
	return nil
}

func (s *TestBin2) Size() int {
	ln := 0
	ln += 4
	ln += int(s.LenName)
	ln += 4
	return ln
}
func (s *BCastHeader) Marshal(wr io.Writer) error {
	buf := make([]byte, 8)
	var (
		byBuf []byte
		err   error
	)
	byBuf = make([]byte, len(s.Reserved1))
	copy(byBuf, s.Reserved1)
	_, err = wr.Write(byBuf)
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint32(buf[:4], uint32(s.LogTime))
	_, err = wr.Write(buf[:4])
	if err != nil {
		return err
	}
	byBuf = make([]byte, len(s.AlphaChar))
	copy(byBuf, s.AlphaChar)
	_, err = wr.Write(byBuf)
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint16(buf[:2], uint16(s.TransCode))
	_, err = wr.Write(buf[:2])
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint16(buf[:2], uint16(s.ErrorCode))
	_, err = wr.Write(buf[:2])
	if err != nil {
		return err
	}
	binary.BigEndian.PutUint32(buf[:4], uint32(s.BCSeqNo))
	_, err = wr.Write(buf[:4])
	if err != nil {
		return err
	}
	byBuf = make([]byte, len(s.Reserved2))
	copy(byBuf, s.Reserved2)
	_, err = wr.Write(byBuf)
	if err != nil {
		return err
	}
	byBuf = make([]byte, len(s.Timestamp2))
	copy(byBuf, s.Timestamp2)
	_, err = wr.Write(byBuf)
	if err != nil {
		return err
	}
	byBuf = make([]byte, len(s.Filler2))
	copy(byBuf, s.Filler2)
	_, err = wr.Write(byBuf)
	if err != nil {
		return err
	}
	return nil
}

func (s *BCastHeader) Size() int {
	ln := 0
	ln += 4
	ln += 4
	ln += 2
	ln += 2
	ln += 2
	ln += 4
	ln += 4
	ln += 8
	ln += 8
	return ln
}
func (s *BCastPacket) Marshal(wr io.Writer) error {
	buf := make([]byte, 8)
	var (
		byBuf []byte
		err   error
	)
	if err := s.Header.Marshal(wr); err != nil {
		return err
	}
	binary.BigEndian.PutUint16(buf[:2], uint16(s.MessageLength))
	_, err = wr.Write(buf[:2])
	if err != nil {
		return err
	}
	byBuf = make([]byte, len(s.Packet))
	copy(byBuf, s.Packet)
	_, err = wr.Write(byBuf)
	if err != nil {
		return err
	}
	return nil
}

func (s *BCastPacket) Size() int {
	ln := 0
	ln += 38
	ln += 2
	ln += int(s.MessageLength)
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
	if _, err := io.ReadFull(rdr, buf[:4]); err != nil {
		return err
	}
	s.LenEmbedded = int32(binary.BigEndian.Uint32(buf[:4]))
	if err := s.Embedded.Unmarshal(rdr); err != nil {
		return err
	}
	return nil
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
	return nil
}
func (s *BCastHeader) Unmarshal(rdr io.Reader) error {
	buf := make([]byte, 8)
	bufReserved1 := make([]byte, 4)
	if _, err := io.ReadFull(rdr, bufReserved1); err != nil {
		return err
	}
	s.Reserved1 = string(bufReserved1)
	if _, err := io.ReadFull(rdr, buf[:4]); err != nil {
		return err
	}
	s.LogTime = int32(binary.BigEndian.Uint32(buf[:4]))
	bufAlphaChar := make([]byte, 2)
	if _, err := io.ReadFull(rdr, bufAlphaChar); err != nil {
		return err
	}
	s.AlphaChar = string(bufAlphaChar)
	if _, err := io.ReadFull(rdr, buf[:2]); err != nil {
		return err
	}
	s.TransCode = int16(binary.BigEndian.Uint16(buf[:2]))
	if _, err := io.ReadFull(rdr, buf[:2]); err != nil {
		return err
	}
	s.ErrorCode = int16(binary.BigEndian.Uint16(buf[:2]))
	if _, err := io.ReadFull(rdr, buf[:4]); err != nil {
		return err
	}
	s.BCSeqNo = int32(binary.BigEndian.Uint32(buf[:4]))
	bufReserved2 := make([]byte, 4)
	if _, err := io.ReadFull(rdr, bufReserved2); err != nil {
		return err
	}
	s.Reserved2 = string(bufReserved2)
	bufTimestamp2 := make([]byte, 8)
	if _, err := io.ReadFull(rdr, bufTimestamp2); err != nil {
		return err
	}
	s.Timestamp2 = string(bufTimestamp2)
	bufFiller2 := make([]byte, 8)
	if _, err := io.ReadFull(rdr, bufFiller2); err != nil {
		return err
	}
	s.Filler2 = string(bufFiller2)
	return nil
}
func (s *BCastPacket) Unmarshal(rdr io.Reader) error {
	buf := make([]byte, 8)
	if err := s.Header.Unmarshal(rdr); err != nil {
		return err
	}
	if _, err := io.ReadFull(rdr, buf[:2]); err != nil {
		return err
	}
	s.MessageLength = int16(binary.BigEndian.Uint16(buf[:2]))
	s.Packet = make([]byte, int(s.MessageLength))
	if _, err := io.ReadFull(rdr, s.Packet); err != nil {
		return err
	}
	return nil
}
