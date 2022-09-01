package main

import (
	"encoding/csv"
	"fmt"
	"gopkg.in/natefinch/npipe.v2"
	"io"
	"log"
	"os"
)

func main() {
	// open file
	f, err := os.Open("si.txt")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		// fmt.Println(rec)

		//pipe

		conn, err := npipe.Dial(`\\.\pipe\some_pipe`)
		if err != nil {
			// handle error
		}
		if _, err := fmt.Fprintln(conn, rec); err != nil {
			// handle error
		}

	}
}
