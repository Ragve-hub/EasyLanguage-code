// client
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/natefinch/npipe.v2"
)

func main() {

	// variables

	//qty := 1
	var Position bool = false
	//	Ref1 := 10.2
	//	SRef1 := 1000

	// logging

	file, err := os.OpenFile("trades.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	// Pipe
	ln, err := npipe.Listen(`\\.\pipe\some_pipe`)
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// handle connection like any other net.Conn
		go func(conn net.Conn) {

			r := bufio.NewReader(conn)
			msg, err := r.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			s := strings.Split(msg, ";")
			//fmt.Println(s)

			//Parsing Tick
			//**************************************

			dates, err := time.Parse("20060102", s[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			times, err := time.Parse("150405", s[1])
			if err != nil {
				fmt.Println(err)
				return
			}

			price, err := strconv.ParseFloat(s[3], 32)
			if err != nil {
				panic(err)
			}

			volume, err := strconv.ParseFloat(s[2], 32)
			if err != nil {
				panic(err)
			}
			//Parsing Bid/Ask
			//**************************************

			ask_v, err := strconv.ParseFloat(s[5], 32)
			if err != nil {
				panic(err)
			}

			bid_v, err := strconv.ParseFloat(s[7], 32)
			if err != nil {
				panic(err)
			}

			ask, err := strconv.ParseFloat(s[4], 32)
			if err != nil {
				panic(err)
			}

			bid, err := strconv.ParseFloat(s[6], 32)
			if err != nil {
				panic(err)
			}
			// Print*******************************************

			fmt.Println("**bid-ask**")
			fmt.Println("ask:", ask, ask_v)
			fmt.Println("bid:", bid, bid_v)

			fmt.Println("**tick**")
			fmt.Println("date:", dates)
			fmt.Println("time:", times)
			fmt.Println("price:", price)
			fmt.Println("volume:", volume)
			fmt.Println("position:", Position)

			//Strategy*****************************************
			if Position == false && times.Second() == 25 {

				fmt.Println(dates, ";"+"Buy"+";", price, ";")
				Position = true

				log.Println("Buy ", float32(price))
			}

			if Position == true && times.Second() == 45 {

				fmt.Println(dates, ";"+"Cover"+";", price, ";")
				Position = false

				log.Println("Cover ", price)
			}

		}(conn)
	}

	//

}
