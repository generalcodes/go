package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var msg, eachS string
var countS, constantS int

func main() {
	ch := make(chan string)

	readFileBlob, _ := ioutil.ReadFile("websitlist.txt")
	websites := strings.Split(string(readFileBlob), "\r\n")
	tStart := time.Now()

	for countS, eachS = range websites {
		go geteach(eachS, ch)
	}

	constantS = countS + 1

	tCurrent := time.Now()
	fmt.Printf("It takes %v to initialize all the goroutines\n", tCurrent.Sub(tStart))

	for contextCh := range ch {
		fmt.Println(contextCh)
		countS--
		if countS < 0 {
			close(ch)
		}
	}

	tCurrent = time.Now()
	fmt.Printf("It takes %v to visit all %d websites—Concurreny runs\n", tCurrent.Sub(tStart), constantS)

	fmt.Println("Press Enter to get the time in 'serial' runs...\nPress CTRL+C to end this program ...")
	fmt.Scanln()

	tStart = time.Now()
	for _, eachS = range websites {
		http.Get(eachS)
		fmt.Println(eachS)
	}
	tCurrent = time.Now()
	fmt.Printf("It takes %v to visit all %d websites—in normal serial runs\n", tCurrent.Sub(tStart), constantS)
}

func geteach(eachS string, ch chan string) {
	if _, err := http.Get(eachS); err != nil {
		msg = eachS + " is down"
	} else {
		msg = eachS + " is up"
	}
	ch <- msg
}
