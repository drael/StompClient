/*
    Simple stomp Client to Send data to ActiveMQ
    Default Stomp conf can be overwritten setting
    stomp environment variables like:
        STOMP_HOST
        STOMP_PORT
        STOMP_DEST

    Author: Rafael Santos <rafael@sourcecode.net.br>

*/

package StompClient

import (
    "github.com/gmallard/stompngo"
    "net"
    "log"
    "os"
)


type StompClient struct {
    Host            string
    Port            string
    Dest            string
    NetConnection   net.Conn
    Connection      *stompngo.Connection
}


// Set stomp info
func (g *StompClient) getVars() {
    g.Host = "localhost"
    g.Port = "61613"
    g.Dest = "/queue/events"

    // Check stomp environment variables
    stomp_host := os.Getenv("STOMP_HOST")
    if stomp_host != "" {
        g.Host = stomp_host
    }

    stomp_port := os.Getenv("STOMP_PORT")
    if stomp_port != "" {
        g.Port = stomp_port
    }

    stomp_dest := os.Getenv("STOMP_DEST")
    if stomp_dest != "" {
        g.Dest = stomp_dest
    }
}


// Open connection
func (nt *StompClient) netConnection() {
    n, err := net.Dial("tcp", net.JoinHostPort(nt.Host, nt.Port))
    if err != nil {
        log.Fatal(err)
    }

    nt.NetConnection = n
}


// Connect to stomp
func (conn *StompClient) Connect() (*stompngo.Connection) {
    conn.getVars()
    conn.netConnection()
    stompHeaders := stompngo.Headers{}

    c, err := stompngo.Connect(conn.NetConnection, stompHeaders)
    if err != nil {
        log.Print("Can't connect to Stomp.")
        log.Fatal(err)
    }

    conn.Connection = c

    return c
}


// Send data to ActiveMQ
func (s *StompClient) SendEvent(data, event string) {
    headers := stompngo.Headers{"destination", s.Dest, "eventtype", event}

    //sending message
    err := s.Connection.Send(headers, data)
    if err != nil {
        log.Print("Can't send the message.")
        log.Fatal(err)
    }
    log.Print("Message sent!")
}