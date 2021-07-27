package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type RouterOspfArea struct {
	AreaIpv4 RouterOspfAreaInstance `json:"area,omitempty"`
}

type RouterOspfAreaInstance struct {
	AreaIpv4          string                          `json:"area-ipv4,omitempty"`
	AreaNum           int                             `json:"area-num,omitempty"`
	Authentication    RouterOspfAreaAuthCfg           `json:"auth-cfg,omitempty"`
	DefaultCost       int                             `json:"default-cost,omitempty"`
	FilterList        []RouterOspfAreaFilterLists     `json:"filter-lists,omitempty"`
	Nssa              RouterOspfAreaNssaCfg           `json:"nssa-cfg,omitempty"`
	AreaRangePrefix   []RouterOspfAreaRangeList       `json:"range-list,omitempty"`
	Shortcut          string                          `json:"shortcut,omitempty"`
	Stub              RouterOspfAreaStubCfg           `json:"stub-cfg,omitempty"`
	UUID              string                          `json:"uuid,omitempty"`
	VirtualLinkIPAddr []RouterOspfAreaVirtualLinkList `json:"virtual-link-list,omitempty"`
}

type RouterOspfAreaAuthCfg struct {
	Authentication int `json:"authentication,omitempty"`
	MessageDigest  int `json:"message-digest,omitempty"`
}

type RouterOspfAreaFilterLists struct {
	AclDirection   string `json:"acl-direction,omitempty"`
	AclName        string `json:"acl-name,omitempty"`
	FilterList     int    `json:"filter-list,omitempty"`
	PlistDirection string `json:"plist-direction,omitempty"`
	PlistName      string `json:"plist-name,omitempty"`
}

type RouterOspfAreaNssaCfg struct {
	DefaultInformationOriginate int    `json:"default-information-originate,omitempty"`
	Metric                      int    `json:"metric,omitempty"`
	MetricType                  int    `json:"metric-type,omitempty"`
	NoRedistribution            int    `json:"no-redistribution,omitempty"`
	NoSummary                   int    `json:"no-summary,omitempty"`
	Nssa                        int    `json:"nssa,omitempty"`
	TranslatorRole              string `json:"translator-role,omitempty"`
}

type RouterOspfAreaRangeList struct {
	AreaRangePrefix string `json:"area-range-prefix,omitempty"`
	Option          string `json:"option,omitempty"`
}

type RouterOspfAreaStubCfg struct {
	NoSummary int `json:"no-summary,omitempty"`
	Stub      int `json:"stub,omitempty"`
}

type RouterOspfAreaVirtualLinkList struct {
	AuthenticationKey         string `json:"authentication-key,omitempty"`
	Bfd                       int    `json:"bfd,omitempty"`
	DeadInterval              int    `json:"dead-interval,omitempty"`
	HelloInterval             int    `json:"hello-interval,omitempty"`
	Md5                       string `json:"md5,omitempty"`
	MessageDigestKey          int    `json:"message-digest-key,omitempty"`
	RetransmitInterval        int    `json:"retransmit-interval,omitempty"`
	TransmitDelay             int    `json:"transmit-delay,omitempty"`
	VirtualLinkAuthType       string `json:"virtual-link-auth-type,omitempty"`
	VirtualLinkAuthentication int    `json:"virtual-link-authentication,omitempty"`
	VirtualLinkIPAddr         string `json:"virtual-link-ip-addr,omitempty"`
}

func PostRouterOspfArea(id string, name1 string, inst RouterOspfArea, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostRouterOspfArea")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/router/ospf/"+name1+"/area", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterOspfArea
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostRouterOspfArea", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetRouterOspfArea(id string, name1 string, name2 string, name3 string, host string) (*RouterOspfArea, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetRouterOspfArea")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/router/ospf/"+name1+"/area/"+name2+"+"+name3, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterOspfArea
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetRouterOspfArea", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutRouterOspfArea(id string, name1 string, name2 string, name3 string, inst RouterOspfArea, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutRouterOspfArea")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/router/ospf/"+name1+"/area/"+name2+"+"+name3, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterOspfArea
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutRouterOspfArea", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteRouterOspfArea(id string, name1 string, name2 string, name3 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteRouterOspfArea")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/router/ospf/"+name1+"/area/"+name2+"+"+name3, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m RouterOspfArea
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteRouterOspfArea", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
