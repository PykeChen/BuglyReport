package mhttp

import (
	"fmt"
	"net/http"
	"time"
)

var DoClient = &http.Client{
	Timeout: time.Minute * 2,
}

func init() {
	fmt.Printf("Http client init %v", time.Now().Format("2006-01-02 15:04:05"))
}


