package server

import (
	"math"
	"testing"
	"time"
)

func Test_time2Bin(t *testing.T) {
	ts := time.Now()

	b := time2bin(&ts)
	intExpt := ts.UnixMilli()

	var intGot int64 = int64(b[0])
	for i := 1; i < len(b); i++ {
		intGot <<= 8
		intGot |= int64(b[i])
	}

	if intExpt != intGot {
		t.Errorf("time2Bin err. excpect %v, got %v", intExpt, intGot)
	}
}

func Test_int2bin(t *testing.T) {
	var intExpt int64 = math.MinInt64

	b := int642bin(intExpt)

	var intGot int64 = int64(b[0])
	for i := 1; i < len(b); i++ {
		intGot <<= 8
		intGot |= int64(b[i])
	}

	if intExpt != intGot {
		t.Errorf("time2Bin err. excpect %v, got %v", intExpt, intGot)
	}

}

func Test_string2Bin(t *testing.T) {
	str := "abc"
	bytes := string2bin(str)
	if string(bytes) != str {
		t.Errorf("string2Bin err")
	}
}
