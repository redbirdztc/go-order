package main

import "time"

// message object to send
type message struct {
	id int // 8B

	ts      time.Time // 8B
	content string    // dynamic max 2MB

	topic   string    // dynamic max 128B
	topicTp TopicType // 8B
}

func newMessage(ts time.Time, content string, topic string, topicTp TopicType) (m *message) {
	m = &message{
		ts:      ts,
		content: content,
		topic:   topic,
		topicTp: topicTp,
	}
	m.markID()
	return
}

// duplicate message with a new id
func (m *message) duplicate() *message {
	dup := &message{
		ts:      m.ts,
		content: m.content,
		topic:   m.topic,
		topicTp: m.topicTp,
	}
	dup.markID()
	return dup
}

func (m *message) markID() {
	m.id = goorder.idM.getID(idTypeMessage)
}
