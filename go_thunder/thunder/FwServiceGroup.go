package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwServiceGroup struct {
	FwServiceGroupInstanceName FwServiceGroupInstance `json:"service-group,omitempty"`
}

type FwServiceGroupInstance struct {
	FwServiceGroupInstanceHealthCheck                    string                                 `json:"health-check,omitempty"`
	FwServiceGroupInstanceMemberListPort                 []FwServiceGroupInstanceMemberList     `json:"member-list,omitempty"`
	FwServiceGroupInstanceName                           string                                 `json:"name,omitempty"`
	FwServiceGroupInstancePacketCaptureTemplate          string                                 `json:"packet-capture-template,omitempty"`
	FwServiceGroupInstanceProtocol                       string                                 `json:"protocol,omitempty"`
	FwServiceGroupInstanceSamplingEnableCounters1        []FwServiceGroupInstanceSamplingEnable `json:"sampling-enable,omitempty"`
	FwServiceGroupInstanceTrafficReplicationMirrorIPRepl int                                    `json:"traffic-replication-mirror-ip-repl,omitempty"`
	FwServiceGroupInstanceUUID                           string                                 `json:"uuid,omitempty"`
	FwServiceGroupInstanceUserTag                        string                                 `json:"user-tag,omitempty"`
}

type FwServiceGroupInstanceMemberList struct {
	FwServiceGroupInstanceMemberListName                    string                                           `json:"name,omitempty"`
	FwServiceGroupInstanceMemberListPacketCaptureTemplate   string                                           `json:"packet-capture-template,omitempty"`
	FwServiceGroupInstanceMemberListPort                    int                                              `json:"port,omitempty"`
	FwServiceGroupInstanceMemberListSamplingEnableCounters1 []FwServiceGroupInstanceMemberListSamplingEnable `json:"sampling-enable,omitempty"`
	FwServiceGroupInstanceMemberListUUID                    string                                           `json:"uuid,omitempty"`
	FwServiceGroupInstanceMemberListUserTag                 string                                           `json:"user-tag,omitempty"`
}

type FwServiceGroupInstanceSamplingEnable struct {
	FwServiceGroupInstanceSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

type FwServiceGroupInstanceMemberListSamplingEnable struct {
	FwServiceGroupInstanceMemberListSamplingEnableCounters1 string `json:"counters1,omitempty"`
}

func PostFwServiceGroup(id string, inst FwServiceGroup, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwServiceGroup")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/service-group", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServiceGroup
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostFwServiceGroup", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwServiceGroup(id string, name1 string, host string) (*FwServiceGroup, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwServiceGroup")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/service-group/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServiceGroup
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetFwServiceGroup", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutFwServiceGroup(id string, name1 string, inst FwServiceGroup, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutFwServiceGroup")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/fw/service-group/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServiceGroup
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutFwServiceGroup", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteFwServiceGroup(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteFwServiceGroup")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/fw/service-group/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwServiceGroup
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteFwServiceGroup", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
