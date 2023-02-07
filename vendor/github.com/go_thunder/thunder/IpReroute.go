package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type Reroute struct {
	UUID RerouteInstance `json:"reroute,omitempty"`
}

type SuppressProtocols struct {
	Static    int    `json:"static,omitempty"`
	Isis      int    `json:"isis,omitempty"`
	UUID      string `json:"uuid,omitempty"`
	Ospf      int    `json:"ospf,omitempty"`
	Rip       int    `json:"rip,omitempty"`
	Ibgp      int    `json:"ibgp,omitempty"`
	Ebgp      int    `json:"ebgp,omitempty"`
	Connected int    `json:"connected,omitempty"`
}
type RerouteInstance struct {
	Static SuppressProtocols `json:"suppress-protocols,omitempty"`
	UUID   string            `json:"uuid,omitempty"`
}

func PostIpReroute(id string, inst Reroute, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpReroute")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/reroute", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Reroute
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpReroute REQ RES..........................", m)
			err := check_api_status("PostIpReroute", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetIpReroute(id string, host string) (*Reroute, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpReroute")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/reroute/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Reroute
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIpReroute REQ RES..........................", m)
			err := check_api_status("GetIpReroute", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
