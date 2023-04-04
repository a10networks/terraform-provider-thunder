package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateDnsClassList struct {
	SlbTemplateDNSClassListInstanceName SlbTemplateDNSClassListInstance `json:"class-list,omitempty"`
}

type SlbTemplateDNSClassListInstance struct {
	SlbTemplateDNSClassListInstanceLidListLidnum []SlbTemplateDNSClassListInstanceLidList `json:"lid-list,omitempty"`
	SlbTemplateDNSClassListInstanceName          string                                   `json:"name,omitempty"`
	SlbTemplateDNSClassListInstanceUUID          string                                   `json:"uuid,omitempty"`
}

type SlbTemplateDNSClassListInstanceLidList struct {
	SlbTemplateDNSClassListInstanceLidListActionValue     string                                    `json:"action-value,omitempty"`
	SlbTemplateDNSClassListInstanceLidListConnRateLimit   int                                       `json:"conn-rate-limit,omitempty"`
	SlbTemplateDNSClassListInstanceLidListDNSCacheAction  SlbTemplateDNSClassListInstanceLidListDNS `json:"dns,omitempty"`
	SlbTemplateDNSClassListInstanceLidListLidnum          int                                       `json:"lidnum,omitempty"`
	SlbTemplateDNSClassListInstanceLidListLockout         int                                       `json:"lockout,omitempty"`
	SlbTemplateDNSClassListInstanceLidListLog             int                                       `json:"log,omitempty"`
	SlbTemplateDNSClassListInstanceLidListLogInterval     int                                       `json:"log-interval,omitempty"`
	SlbTemplateDNSClassListInstanceLidListOverLimitAction int                                       `json:"over-limit-action,omitempty"`
	SlbTemplateDNSClassListInstanceLidListPer             int                                       `json:"per,omitempty"`
	SlbTemplateDNSClassListInstanceLidListUUID            string                                    `json:"uuid,omitempty"`
	SlbTemplateDNSClassListInstanceLidListUserTag         string                                    `json:"user-tag,omitempty"`
}

type SlbTemplateDNSClassListInstanceLidListDNS struct {
	SlbTemplateDNSClassListInstanceLidListDNSCacheAction            string `json:"cache-action,omitempty"`
	SlbTemplateDNSClassListInstanceLidListDNSHonorServerResponseTTL int    `json:"honor-server-response-ttl,omitempty"`
	SlbTemplateDNSClassListInstanceLidListDNSTTL                    int    `json:"ttl,omitempty"`
	SlbTemplateDNSClassListInstanceLidListDNSWeight                 int    `json:"weight,omitempty"`
}

func PostSlbTemplateDnsClassList(id string, name1 string, inst SlbTemplateDnsClassList, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateDnsClassList")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/dns/"+name1+"/class-list", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateDnsClassList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateDnsClassList", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateDnsClassList(id string, name1 string, host string) (*SlbTemplateDnsClassList, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateDnsClassList")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/dns/"+name1+"/class-list", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateDnsClassList
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateDnsClassList", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
