package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"time"
)

// Start a http server, then differentiate the connenter's
// charactor, publishing or subscribing.
//
// For publishing, topic, topic type and message content is needed.
// If the queue not exists, it will be created with goroutine.After
// distributing message to all queues or picking a random queue to
// send message (according to queue type), response ok to publisher.
//
// For subscribing, connection should keep alive and send described
// topics to register in specific queues. So that, when queue is
// not empty, it will try send message to subscriber. In case of
// subscribing fanout topic type, queue will be created if the
// connection hasn't registered.
//
// post /message
// {"topic":"mytopic","topicType":0,"message":"myMessage","timestamp":1658187351342534}
//
// post /subscribe
// {"topics":["mytopic","mytopic1"]}
//
func servehttp() {
	http.HandleFunc("/message", publishMsg)
	http.HandleFunc("/subscribe", publishMsg)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

type publishMsgReq struct {
	Topic   string    `json:"topic"`
	TopicTp TopicType `json:"topicType"`
	Msg     string    `json:"message"`
	Ts      int64     `json:"timestamp"`
}

func publishMsg(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}()

	r.Close = true

	body := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(body); err != nil && err != io.EOF {
		fmt.Println(err)
		debug.PrintStack()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req := &publishMsgReq{}
	if err := json.Unmarshal(body, req); err != nil {
		fmt.Println(err)
		debug.PrintStack()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	topic := goorder.getTopic(req.Msg)
	brokerMsg := newMessage(time.UnixMilli(req.Ts), req.Msg, req.Topic, req.TopicTp)
	topic.distributeMsg(brokerMsg)

	w.WriteHeader(http.StatusOK)
}

type subscribeReq struct {
	Topics []string `json:"topics"`
}

func subcribeTopics(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}()

	r.Close = true

	body := make([]byte, r.ContentLength)
	if _, err := r.Body.Read(body); err != nil && err != io.EOF {
		fmt.Println(err)
		debug.PrintStack()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req := &subscribeReq{}
	if err := json.Unmarshal(body, req); err != nil {
		fmt.Println(err)
		debug.PrintStack()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// it may cause dead lock
	for _, v := range req.Topics {
		tpc := goorder.getTopic(v)

		tpc.mu.Lock()

		queue

	}

	w.WriteHeader(http.StatusOK)
}
