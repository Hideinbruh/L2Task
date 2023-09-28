package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Telnet struct {
	Timeout time.Duration
	Host    string
	Port    string
}

func (t *Telnet) createClient(conn net.Conn) {
	console := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(console)
	for {
		fmt.Print("от клиента: ")
		text, err := console.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		text = strings.TrimSpace(text)

		fmt.Fprintf(conn, text+"\n")
		if text == "exit" {
			fmt.Fprintf(os.Stdout, "%s\n", "соединение закрыто")
			return
		}
		message, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		message = strings.TrimSpace(message)
		fmt.Println("от сервера:", message)
	}
}

func (t *Telnet) parseArgs() bool {
	if len(os.Args) == 3 {
		t.Host = os.Args[1]
		t.Port = os.Args[2]
		return true
	}
	if len(os.Args) == 4 {
		arg := os.Args[1]
		substr := "--timeout="
		if strings.Contains(arg, substr) {
			timeDuration := strings.TrimPrefix(arg, substr)
			timeout, err := time.ParseDuration(timeDuration)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
			}
			t.Timeout = timeout
		} else {
			return false
		}
		t.Host = os.Args[2]
		t.Port = os.Args[3]
		return true
	}
	return false
}

func main() {
	t := &Telnet{}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT)

	ok := t.parseArgs()
	if !ok {
		fmt.Fprintln(os.Stderr, "Некорректный синтаксис")
	}
	d := net.Dialer{Timeout: t.Timeout}
	conn, err := d.Dial("tcp", t.Host+":"+t.Port)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer conn.Close()

	go t.createClient(conn)

	select {
	case <-quit:
		fmt.Fprintf(os.Stdout, "%s\n", "Завершение работы программы")
		os.Exit(0)

	}
}
