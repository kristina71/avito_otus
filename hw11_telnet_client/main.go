package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.SetFlags(0)

	timeout := flag.Duration("timeout", 10*time.Second, "server connect timeout")
	flag.Parse()
	if flag.NArg() != 2 {
		log.Fatal("Please define address and port")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	address := net.JoinHostPort(host, port)

	ctx, cancel := context.WithCancel(context.Background())

	go gracefulShutDown(cancel)

	client := NewTelnetClient(address, *timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	go send(client, cancel)
	go receive(client, cancel)

	<-ctx.Done()
}

func gracefulShutDown(cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	cancel()
}

func send(client TelnetClient, cancel context.CancelFunc) {
	if err := client.Send(); err != nil {
		log.Println(err)
	}
	cancel()
}

func receive(client TelnetClient, cancel context.CancelFunc) {
	if err := client.Receive(); client.Receive() != nil {
		log.Println(err)
	}
	cancel()
}
