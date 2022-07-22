package main

import (
	"sync"
)

// TopicType Queue Type
type TopicType int

const (
	// TopicTypeSignle Signle to signle queue will send msg to random subscriber
	TopicTypeSignle TopicType = iota
	// TopicTypeFanout Fanout queue will send msg to every subscriber
	TopicTypeFanout
)

func topicMutex(topic string) *sync.RWMutex {
	mutex := &sync.RWMutex{}
	mutexI, _ := goorder._topicMutex.LoadOrStore(topic, mutex)
	return mutexI.(*sync.RWMutex)
}

// topic flag for producer sending and consumer receving
type topic struct {
	t      string
	tp     TopicType
	queues []*queue

	mu *sync.RWMutex
}

// newTopic create a `t` topic with a empty queue
func newTopic(t string, tp TopicType) *topic {
	mutex := topicMutex(t)
	mutex.Lock()
	defer mutex.Unlock()

	var qt queueType
	switch tp {
	case TopicTypeSignle:
		qt = queueTypeSingle
	case TopicTypeFanout:
		qt = queueTypeFanout
	}

	tpc := &topic{t: t, tp: tp, queues: make([]*queue, 0)}

	tpc.queues = append(tpc.queues, newQueue(qt, tpc))

	goorder.topicMap[t] = tpc

	return tpc
}

// topic length lte 128
func (t *topic) legal() bool {
	return len(t.t) <= 128
}

// for fanout topic type bind queue, will not check
func (t *topic) bindQueue(q *queue) {
	mu := topicMutex(t.t)
	mu.Lock()
	defer mu.Unlock()

	q.tpc = t
	t.queues = append(t.queues, q)
}

func (t *topic) distributeMsg(msg *message) {
	mu := topicMutex(t.t)
	mu.RLock()
	defer mu.RUnlock()

	qs := t.queues

	qs[0].addMsg(msg)
	for i := 1; i < len(qs); i++ {
		qs[i].addMsg(msg.duplicate())
	}
}
