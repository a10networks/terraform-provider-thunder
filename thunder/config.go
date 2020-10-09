package thunder

import (
	"fmt"
	"util"
)

type Config struct {
	Address        string
	Username       string
	Password       string
	LoginReference string
}

func (c *Config) Client() (Thunder, error) {
	logger := util.GetLoggerInstance()
	var vt Thunder

	if c.Address != "" && c.Username != "" && c.Password != "" {
		logger.Println("[INFO] Initializing THUNDER connection")
		var client Thunder
		var err error

		client, err = NewTokenSession(c.Address, c.Username, c.Password)

		if err != nil {
			logger.Printf("[ERROR] Error creating New Token Session %s ", err)
			return vt, err
		}

		logger.Println("[INFO] GOT TOKEN FROM API.......................... ")
		return client, nil
	}
	return vt, fmt.Errorf("THUNDER provider requires address, username and password")
}
