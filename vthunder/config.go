package vthunder

import (
	"fmt"
	"util"
	//"log"
)

type Config struct {
	Address        string
	Username       string
	Password       string
	LoginReference string
}

func (c *Config) Client() (vThunder, error) {
	logger := util.GetLoggerInstance()
	var vt vThunder

	if c.Address != "" && c.Username != "" && c.Password != "" {
		logger.Println("[INFO] Initializing VTHUNDER connection")
		var client vThunder
		var err error

		client, err = NewTokenSession(c.Address, c.Username, c.Password)

		if err != nil {
			logger.Printf("[ERROR] Error creating New Token Session %s ", err)
			return vt, err
		}

		logger.Println("[INFO] GOT TOKEN FROM API.......................... " + client.Token)
		return client, nil
	}
	return vt, fmt.Errorf("VTHUNDER provider requires address, username and password")
}
