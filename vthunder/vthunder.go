package vthunder

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type vThunder struct {
	Host      string
	User      string
	Password  string
	Token     string
	Transport *http.Transport
}

type APIRequest struct {
	Method      string
	URL         string
	Body        string
	ContentType string
}

type RequestError struct {
	Code       int      `json:"code,omitempty"`
	Message    string   `json:"message,omitempty"`
	ErrorStack []string `json:"errorStack,omitempty"`
}

func (r *RequestError) Error() error {
	if r.Message != "" {
		return errors.New(r.Message)
	}

	return nil
}

func NewSession(host, user, passwd string) vThunder {
	var url string
	if !strings.HasPrefix(host, "http") {
		url = fmt.Sprintf("https://%s", host)
	} else {
		url = host
	}

	return vThunder{
		Host:     url,
		User:     user,
		Password: passwd,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}

func NewTokenSession(host string, user string, passwd string) (b vThunder, err error) {
	log.Println("[INFO] TOKEN")
	b.Host = host
	b.Token = getAuthHeader(host, user, passwd, "/axapi/v3/auth")
	return
}
