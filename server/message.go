package server

import "time"

// MessageType message will be send in some way
type MessageType int

const (
	// MessageTypeSingle One-2-one mode
	// For a message, it will be consumed for one time by one consumer whatever the number of subscribers
	MessageTypeSingle MessageType = iota
	// MessageTypeFanout One-2-many mode
	// For a message, it will be sent to all consumer that subscribe the message topic
	MessageTypeFanout
)

// Message object to send
type Message struct {
	id      int         // 8B
	ts      time.Time   // 8B
	content string      // dynamic max 2MB
	topic   Topic       // dynamic max 128B
	tp      MessageType // 8B
}
