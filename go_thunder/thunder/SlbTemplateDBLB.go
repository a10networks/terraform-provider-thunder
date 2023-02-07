package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type DBLB struct {
	UUID DblbInstance `json:"dblb,omitempty"`
}

type CalcSha1 struct {
	Sha1Value string `json:"sha1-value,omitempty"`
}
type DblbInstance struct {
	ServerVersion string   `json:"server-version,omitempty"`
	Name          string   `json:"name,omitempty"`
	ClassList     string   `json:"class-list,omitempty"`
	UserTag       string   `json:"user-tag,omitempty"`
	Sha1Value     CalcSha1 `json:"calc-sha1,omitempty"`
	UUID          string   `json:"uuid,omitempty"`
}

func PostTemplateDBLB(id string, inst DBLB, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostTemplateDBLB")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/dblb", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DBLB
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostTemplateDBLB REQ RES..........................", m)
			err := check_api_status("PostTemplateDBLB", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetTemplateDBLB(id string, name string, host string) (*DBLB, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetTemplateDBLB")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/dblb/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DBLB
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetTemplateDBLB REQ RES..........................", m)
			err := check_api_status("GetTemplateDBLB", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutTemplateDBLB(id string, name string, inst DBLB, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutTemplateDBLB")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/dblb/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DBLB
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutTemplateDBLB REQ RES..........................", m)
			err := check_api_status("PutTemplateDBLB", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteTemplateDBLB(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteTemplateDBLB")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/dblb/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m DBLB
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
