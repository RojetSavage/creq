package req

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func NewClient() *http.Client {
	fmt.Println("new client")
	var c http.Client = http.Client{
		Timeout: time.Second * 10,
	}
	return &c
}

func SetClientTimeout(c *http.Client, seconds string) {
	i, err := strconv.Atoi(seconds)

	if err != nil {
		return
	}

	c.Timeout = time.Second * time.Duration(i)
}
