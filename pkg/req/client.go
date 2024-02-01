package req

import (
	"args"
	"net/http"
	"strconv"
	"time"
)

func NewClient() *http.Client {
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

func ApplyFlagsToClient(c *http.Client, flags []args.UserFlag) error {
	var err error

	for _, flag := range flags {
		err = applyFlagToClient(c, flag)
	}

	return err
}

func applyFlagToClient(c *http.Client, f args.UserFlag) error {
	switch f.F {
	case "connect-timeout":
		SetClientTimeout(c, f.Parameter)
	}

	return nil
}
