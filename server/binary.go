package server

import (
	"bytes"
	"encoding/binary"
	"time"
)

func time2bin(t *time.Time) []byte {
	nano := t.UnixNano()
	return int642bin(nano)
}

func int642bin(i int64) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 8))
	_ = binary.Write(buf, binary.BigEndian, i)
	return buf.Bytes()
}

func string2bin(s string) []byte {
	return []byte(s)
}

var location = time.Now().Local().Location()

func bin2time(b []byte) (time.Time, error) {
	nano := bin2int64(b)
	t := time.Unix(0, nano)
	return t, nil
}

func bin2int64(b []byte) (n int64) {
	buf := bytes.NewBuffer(b)
	if err := binary.Read(buf, binary.BigEndian, &n); err != nil {
		panic(err)
	}
	return
}

func bin2string(b []byte) string {
	return string(b)
}
