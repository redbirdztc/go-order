package main

import (
	"testing"
	"time"
)

func Test_msgToBinary(t *testing.T) {
	ts := time.Now()
	m := message{1, ts, "mycontent", "mytopic", TopicTypeSignle}
	r, _ := msgToBinary(&m)

	totalLength := bin2int64(r[:8])

	if int(totalLength) != len(r)-8 {
		t.Errorf("total length err, expect %v, got %v\n", len(r)-8, totalLength)
	}

	id := bin2int64(r[8:16])
	if id != 1 {
		t.Errorf("id err, expect %d, got %d\n", 1, id)
	}

	rsTs, _ := bin2time(r[16:24])
	if !rsTs.Equal(ts) {
		t.Errorf("timestamp err, expect %v, got %v\n", ts, rsTs)
	}

	cttLen := bin2int64(r[24:32])
	topicLen := bin2int64(r[32+cttLen : 40+cttLen])
	if totalLength != 40+cttLen+topicLen {
		t.Errorf("content length or topic length err. got content length: %v, topic length: %v", cttLen, topicLen)
	}

	ctt := bin2string(r[32 : 32+cttLen])
	if ctt != "mycontent" {
		t.Errorf("content err. expect %s, got %s", "mycontent", ctt)
	}

	topic := bin2string(r[40+cttLen : 40+cttLen+topicLen])
	if ctt != "mytopic" {
		t.Errorf("content err. expect %s, got %s", "mytopic", topic)
	}

	msgTp := TopicType(bin2int64(r[40+cttLen+topicLen:]))
	if msgTp != TopicTypeSignle {
		t.Errorf("content err. expect %d, got %d", TopicTypeSignle, msgTp)
	}
}
