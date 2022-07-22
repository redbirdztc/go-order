package main

import (
	"net"
	"strconv"
	"sync"
)

type subscriber struct {
	id   int
	ip   string
	port int
	conn net.Conn

	queues []*queue

	mu *sync.Mutex
}


type subIndex struct {
	hostHash  map[string]*subscriber // {ip:port} : {subscriber}
	queueHash map[int][]*subscriber  // {queue id} : {subscriber}

	mu *sync.Mutex
}

func newSubIndex() *subIndex {
	return &subIndex{
		hostHash:  make(map[string]*subscriber),
		queueHash: make(map[int][]*subscriber),
		mu:        &sync.Mutex{},
	}
}

// index subscriber that registered
func (i *subIndex) index(sub *subscriber, qid int) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	ipport := sub.ip + ":" + strconv.Itoa(sub.port)
	if _, ok := i.hostHash[ipport]; ok {
		return errSubscriberIndexed
	}
	i.hostHash[ipport] = sub

	for _, v := range i.queueHash[qid] {
		if v.id == sub.id {
			return errSubscriberIndexed
		}
	}
	i.queueHash[qid] = append(i.queueHash[qid], sub)
	return nil
}

// remove subscriber index when subscriber unregistered
func (i *subIndex) unindex(sub *subscriber) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	ipport := sub.ip + ":" + strconv.Itoa(sub.port)
	if _, ok := i.hostHash[ipport]; !ok {
		return errSubscriberNotIndexed
	}

	for _, v := range sub.queues {
		subs := i.queueHash[v.id]
		for i := range subs {
			if subs[i] == sub {
				subs[i], subs[len(subs)-1] = subs[len(subs)-1], subs[i]
			}
		}
		subs = subs[:len(subs)-1]
	}

	return nil
}
