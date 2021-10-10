package main

import (
	"fmt"
	"golang.conradwood.net/go-easyops/utils"
)

var (
	// initialized on startup with a random number for each instance and increased
	// for speed we don't insist on 100% uniqueness guarantees but best-effort instead
	requestCounter uint64
)

func start_requestid_creator() {
	requestCounter = uint64(utils.RandomInt(10000000))

}

func req_get_requestid() string {
	requestCounter++
	return fmt.Sprintf("%d", requestCounter)
}
