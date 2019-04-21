package ircclient

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

type Client struct {
	Conn io.ReadWriter
}

type Identity struct {
	Username string
	Realname string
}

func New(c net.Conn) (Client, error) {
	if c == nil {
		return Client{}, errors.New("The connection must not be nil.")
	}
	return Client{Conn: c}, nil
}

func (c *Client) Connect(i Identity, mode int) error {
	c.Nick(i.Username)
	fmt.Fprintf(c.Conn, "USER %s %d * %s\r\n", i.Username, mode, i.Realname)
	return nil
}

func (c *Client) ConnectWithPassword(i Identity, password string) error {
	fmt.Fprintf(c.Conn, "PASS %s\r\n", password)
	return c.Connect(i, 0)
}

func (c *Client) Join(channels []string) error {
	fmt.Fprintf(c.Conn, "JOIN %s\r\n", strings.Join(channels, ","))
	return nil
}

/*
LeaveAllChannels will use the JOIN command with the special parameter "0" to
request to leave all channels the user is currently a member of.

See RFC 2812 §3.2.1:
https://tools.ietf.org/html/rfc2812#section-3.2.1
*/
func (c *Client) LeaveAllChannels() error {
	return c.Join([]string{"0"})
}

func (c *Client) Nick(username string) error {
	fmt.Fprintf(c.Conn, "NICK %s\r\n", username)
	return nil
}

func (c *Client) SendRawMessage(msg string) int {
	c.Conn.Write([]byte(msg))

	return 0
}
