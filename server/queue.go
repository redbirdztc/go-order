package main

import (
	"sync"
)

type queueType int

const (
	queueTypeSingle queueType = iota
	queueTypeFanout
)

type queue struct {
	id   int
	msgs []*message
	qt   queueType

	tpc *topic

	mu *sync.RWMutex
}

func newQueue(qt queueType,tpc *topic) *queue {
	return &queue{
		id:   goorder.idM.getID(idTypeQueue),
		msgs: make([]*message, 0),
		qt:   qt,

		tpc: tpc,

		mu: &sync.RWMutex{},
	}
}

func (q *queue) addMsg(msg *message) {
	m := q.mu
	m.Lock()
	defer m.Unlock()

	q.msgs = append(q.msgs, msg)
}
