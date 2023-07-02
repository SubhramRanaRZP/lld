package main

import (
	"fmt"
	"time"
)

type reqResp struct {
	req    map[string]string
	result string
	resp   map[string]string
}

func prePostProcessor(key, val string, queue chan *reqResp) {
	printWithSleep("[prePost] : preparing request body")
	body := &reqResp{
		req: map[string]string{
			key: val,
		},
	}
	print("[prePost] : request body prepared")

	print("[prePost] : body pushed to queue")
	queue <- body
	print("[prePost] : waiting for the response")

	body = <-queue
	print("[prePost] : received the updated body")

	printWithSleep("[prePost] : update the response field")
	body.resp = map[string]string{
		"resp": "all done",
	}

	fmt.Println("final processed body => ", *body)
}

func coreProcessor(queue chan *reqResp) {
	print("[core] : waiting for the request body")

	body := <-queue
	print("[core] : body received")

	printWithSleep("[core] : preparing body's result")
	body.result = "SUCCESS"

	printWithSleep("[core] : pushing body back to queue")
	queue <- body
}

func main() {
	q := make(chan *reqResp)
	go coreProcessor(q)
	go prePostProcessor("name", "Subhram", q)
	time.Sleep(7 * time.Second)
}

func print(msg string) {
	fmt.Println(msg)
}

func printWithSleep(msg string) {
	print(msg)
	time.Sleep(time.Second)
}
