StompClient
===========

Simple Stomp Client to send data to ActiveMQ.
Build on top of [stompngo](https://github.com/gmallard/stompngo)


Requirements:
* $ go get https://github.com/gmallard/stompngo


Usage
=====

```go
conn := StompClient.StompClient{}
conn.Connect()
conn.SendEvent(`{"foo": "bar"}`, "Hello")
```

Default Stomp conf can be overwritten setting environment variables like:
* STOMP_HOST
* STOMP_PORT
* STOMP_DEST