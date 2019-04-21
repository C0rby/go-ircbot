package ircclient

import (
	"fmt"
	"net"
	"testing"
)

func printServerSide(server net.Conn, expectedLines int, finished chan bool, t *testing.T) {
	for i := 1; i <= expectedLines; i++ {
		out := make([]byte, 1024)
		if _, err := server.Read(out); err != nil {
			t.Error("some error")
		}
		fmt.Print(string(out))
	}
	fmt.Println("")
	finished <- true
}

func TestNewClientWithNilConnection(t *testing.T) {
	var conn net.Conn = nil
	_, err := New(conn)
	if err == nil {
		t.Error("Client should not be creatable without a connection")
	}
}

func TestConnect(t *testing.T) {
	server, c := net.Pipe()
	defer c.Close()
	defer server.Close()

	finished := make(chan bool)

	go printServerSide(server, 2, finished, t)

	client, _ := New(c)
	err := client.Connect(Identity{Username: "Corbot"}, 0)
	if err != nil {
		t.Error("Some error")
	}

	<-finished
}

func TestConnectWithPassword(t *testing.T) {
	server, c := net.Pipe()
	defer c.Close()
	defer server.Close()

	finished := make(chan bool)

	go printServerSide(server, 3, finished, t)

	client, _ := New(c)
	err := client.ConnectWithPassword(Identity{Username: "Corbot"}, "password")
	if err != nil {
		t.Error("Some error")
	}

	<-finished
}
