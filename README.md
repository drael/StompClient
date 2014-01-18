StompClient
===========

Simple Stomp Client to send data to ActiveMQ.
Build on top of [stompngo](https://github.com/gmallard/stompngo)


Requirements:
	$ go get https://github.com/gmallard/stompngo


Usage:
```go
conn := StompClient.StompClient{}
conn.Connect()
conn.SendEvent(`{"foo": "bar"}`, "Hello")
```