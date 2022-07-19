# Feature

* GRPC
* Persistance
* Point to point & fanout
* Auto create topic (Once generated,it won't be changed)

# Design

How message will be handled

* Single to single

```mermaid
flowchart LR
pub(Publisher)-->|gen| msg(Message)
msg-->|sent to| hdl(PublishHandler)
hdl-->|by message topic| q(RandomQueue)
q-->|send msg| subscriber(RandomQueye's Subscriber)
```

* Fanout
  ```mermaid
  flowchart LR
  pub(Publisher)-->|gen| msg(Message)
  msg-->|send to| hdl(PublishHandler)
  hdl-->|by msg topic| q(Queue)
  hdl-->|by msg topic| q1(Queue)
  hdl-->|by msg topic| q2(Queue)
  q-->|send msg| qsub(QueueSubscriber)
  q1-->|send msg| qsub1(QueueSubscriber)
  q2-->|send msg| qsub2(QueueSubscriber)
  ```
