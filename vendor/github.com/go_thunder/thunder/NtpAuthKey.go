package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type NtpAuthKey struct {
	Key NtpAuthKeyInstance `json:"auth-key,omitempty"`
}

type NtpAuthKeyInstance struct {
	AlgType      string `json:"alg-type,omitempty"`
	AscKey       string `json:"asc-key,omitempty"`
	Encrypted    string `json:"encrypted,omitempty"`
	HexEncrypted string `json:"hex-encrypted,omitempty"`
	HexKey       string `json:"hex-key,omitempty"`
	Key          int    `json:"key,omitempty"`
	KeyType      string `json:"key-type,omitempty"`
	UUID         string `json:"uuid,omitempty"`
}

func PostNtpAuthKey(id string, inst NtpAuthKey, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostNtpAuthKey")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ntp/auth-key", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NtpAuthKey
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostNtpAuthKey", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetNtpAuthKey(id string, name1 string, host string) (*NtpAuthKey, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetNtpAuthKey")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ntp/auth-key/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NtpAuthKey
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetNtpAuthKey", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutNtpAuthKey(id string, name1 string, inst NtpAuthKey, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutNtpAuthKey")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/ntp/auth-key/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NtpAuthKey
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutNtpAuthKey", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteNtpAuthKey(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteNtpAuthKey")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/ntp/auth-key/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m NtpAuthKey
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)

		}
	}
	return nil
}
