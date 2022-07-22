package main

import (
	"testing"
	"time"
)

func TestMessage_markID(t *testing.T) {
	msg := new(message)
	msg.markID()

	if msg.id == 0 {
		t.Errorf("message mark id err")
	}
}

func TestMessage_duplicate(t *testing.T) {
	msg := message{
		ts:      time.Now(),
		content: "mymsg",
		topic:   "mytpc",
		topicTp: TopicTypeSignle,
	}
	msg.markID()

	dup := msg.duplicate()

	if msg.id == dup.id {
		t.Errorf("message duplicate same id err")
	}
	if msg.ts.UnixMilli() != dup.ts.UnixMilli() {
		t.Errorf("message timestamp not equal err")
	}
	if msg.content != dup.content {
		t.Errorf("message content not equal err")
	}
	if msg.topic != dup.topic {
		t.Errorf("message topic not equal err")
	}
	if msg.topicTp != dup.topicTp {
		t.Errorf("message topic type not equal err")
	}
}
