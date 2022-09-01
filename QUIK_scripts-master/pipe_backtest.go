// client
package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"gopkg.in/natefinch/npipe.v2"
)

func main() {

	fmt.Println("Backtesting wait....")

	// Pipe
	ln, err := npipe.Listen(`\\.\pipe\some_pipe`)
	if err != nil {
		// handle error
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}

		// handle connection like any other net.Conn
		go func(conn net.Conn) {

			r := bufio.NewReader(conn)
			msg, err := r.ReadString('\n')
			if err != nil {
				// handle error
				return
			}

			s := strings.Split(msg, ";")
			fmt.Println(s)

		}(conn)
	}

	//

}
