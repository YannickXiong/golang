/**
* @Author  : Yannick
* @File    : main.go
* @Date    : 2017-11-20
* @Desc    : boomer is an implementation of locust client by golang.
             This demo is to exercise basic usage of boomer.
*/

package main

import (
	"boomer"
	"log"
	"net/http"
	"time"
)

func now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// TestHTTP :
func TestHTTP() {
	url := "https://www.v2ex.com/api/nodes/show.json?name=python"
	startTime := now()
	resp, err := http.Get(url)
	endTime := now()
	log.Println(float64(endTime - startTime))

	if err != nil {
		boomer.Events.Publish("request_failure", "demo", "http", 0.0, err.Error())
	} else {
		boomer.Events.Publish("request_success", "demo", "http", float64(endTime-startTime), resp.ContentLength)
	}
}

func main() {
	task := &boomer.Task{
		Weight: 10,       // Weight 权重，和 Locust 的 task 权重类似，在有多个 task 的时候生效
		Fn:     TestHTTP, // Fn 类似于 Locust 的 task
	}

	boomer.Run(task)
}
