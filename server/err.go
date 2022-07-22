package main

//go:generate stringer -linecomment -type=brokerErr

// brokerErr
type brokerErr int

const (
	errSubscriberIndexed brokerErr = iota + 1 // subscriber has indexed in SuberscriberIndex
	errSubscriberNotIndexed  // subscriber hasn't indexed in SuberscriberIndex
)

func (err brokerErr) Error() string {
	return err.String()
}
