package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var (
	enteringCliChan = make(chan client)
	leavingCliChan  = make(chan client)
	messagesChan    = make(chan string)
)

type client chan<- string

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clientSet := make(map[client]bool)
	for {
		select {
		case msg := <-messagesChan:
			for cli := range clientSet {
				cli <- msg
			}
		case cli := <-enteringCliChan:
			clientSet[cli] = true
		case cli := <-leavingCliChan:
			delete(clientSet, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	cliChan := make(chan string)
	go clientWriter(conn, cliChan)

	who := conn.RemoteAddr().String()
	cliChan <- "【通知】You are " + who
	messagesChan <- "【通知】" + who + " has arrived"
	enteringCliChan <- cliChan

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messagesChan <- "【消息】[" + who + "] " + input.Text()
	}

	leavingCliChan <- cliChan
	messagesChan <- "【通知】" + who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, cliChan <-chan string) {
	for msg := range cliChan {
		fmt.Fprintln(conn, msg)
	}
}
