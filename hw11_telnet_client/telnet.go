package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type telnetClient struct {
	address    string
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	connection net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (t *telnetClient) Close() error {
	return t.connection.Close()
}

func (t *telnetClient) Send() error {
	_, err := io.Copy(t.connection, t.in)
	return err
}

func (t *telnetClient) Receive() error {
	_, err := io.Copy(t.out, t.connection)
	return err
}

func (t *telnetClient) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return err
	}
	t.connection = conn
	return nil
}
