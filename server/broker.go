package main

import (
	"sync"
)

var goorder = newBroker()

type broker struct {
	idM *idManager

	// map[topic name] topic
	topicMap    map[string]*topic
	_topicMutex *sync.Map

	subIndex
}

func newBroker() *broker {
	b := &broker{
		idM: newIDManager(),

		topicMap:    make(map[string]*topic),
		_topicMutex: new(sync.Map),
	}

	return b
}

func (b *broker) getTopic(t string) *topic {
	mutex := topicMutex(t)
	mutex.Lock()
	defer mutex.Unlock()

	return goorder.topicMap[t]
}
