// client
package main

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"strings"

	"gopkg.in/natefinch/npipe.v2"
)

func main() {

	// variables

	//qty := 1
	var Position bool = false
	//	Ref1 := 10.2
	//	SRef1 := 1000

	//

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
			//fmt.Println(s)

			// tick
			datetime := time.Now().Format("20060802 15:04:05.999")
			// volume := s[2]
			price := s[3]

			//	fmt.Println("tick: ", datetime, price, volume)

			/*	bid-ask
				ask := s[4]
				ask_v := s[5]
				bid := s[6]
				bid_v := s[7]

				//bid-ask
				fmt.Println("bid-ask:", ask, ask_v, bid, bid_v)
			*/

			if Position == false && time.Now().Second() == 25 {

				fmt.Println(datetime + ";" + "Buy" + ";" + price + ";")
				Position = true
			}

			if Position == true && time.Now().Second() == 45 {

				fmt.Println(datetime + ";" + "Cover" + ";" + price + ";")
				Position = false
			}

		}(conn)
	}

	//

}
