package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Imap_pop3 struct {
	UUID ImapPop3Instance `json:"imap-pop3,omitempty"`
}

type ImapPop3Instance struct {
	Logindisabled int    `json:"logindisabled,omitempty"`
	Starttls      string `json:"starttls,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	UserTag       string `json:"user-tag,omitempty"`
	Name          string `json:"name,omitempty"`
}

func PostTemplateImap_POP3(id string, inst Imap_pop3, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostTemplateImap_POP3")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/imap-pop3", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Imap_pop3
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostTemplateImap_POP3 REQ RES..........................", m)
			err := check_api_status("PostTemplateImap_POP3", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetTemplateImap_POP3(id string, name string, host string) (*Imap_pop3, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetTemplateImap_POP3")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/imap-pop3/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Imap_pop3
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetTemplateImap_POP3 REQ RES..........................", m)
			err := check_api_status("GetTemplateImap_POP3", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutTemplateImap_POP3(id string, name string, inst Imap_pop3, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutTemplateImap_POP3")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/imap-pop3/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Imap_pop3
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutTemplateImap_POP3 REQ RES..........................", m)
			err := check_api_status("PutTemplateImap_POP3", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteTemplateImap_POP3(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteTemplateImap_POP3")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/imap-pop3/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Imap_pop3
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}
	return nil
}
