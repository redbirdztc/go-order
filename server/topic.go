package server

// Topic flag for producer sending and consumer receving
type Topic string

// topic max length is 127
func (t *Topic) legal() bool {
	return len(*t) <= 128
}
