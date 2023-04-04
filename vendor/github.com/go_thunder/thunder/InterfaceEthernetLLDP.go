package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"strconv"
	"util"
)

type EthernetLLDP struct {
	UUID EthernetLLDPInstance `json:"lldp,omitempty"`
}

type LLDPTxDot1Cfg struct {
	LinkAggregation int `json:"link-aggregation,omitempty"`
	Vlan            int `json:"vlan,omitempty"`
	TxDot1Tlvs      int `json:"tx-dot1-tlvs,omitempty"`
}
type LLDPNotificationCfg struct {
	Notification int `json:"notification,omitempty"`
	NotifEnable  int `json:"notif-enable,omitempty"`
}
type LLDPEnableCfg struct {
	Rx       int `json:"rx,omitempty"`
	Tx       int `json:"tx,omitempty"`
	RtEnable int `json:"rt-enable,omitempty"`
}
type LLDPTxTlvsCfg struct {
	SystemCapabilities int `json:"system-capabilities,omitempty"`
	SystemDescription  int `json:"system-description,omitempty"`
	ManagementAddress  int `json:"management-address,omitempty"`
	TxTlvs             int `json:"tx-tlvs,omitempty"`
	Exclude            int `json:"exclude,omitempty"`
	PortDescription    int `json:"port-description,omitempty"`
	SystemName         int `json:"system-name,omitempty"`
}
type EthernetLLDPInstance struct {
	LinkAggregation    LLDPTxDot1Cfg       `json:"tx-dot1-cfg,omitempty"`
	Notification       LLDPNotificationCfg `json:"notification-cfg,omitempty"`
	Rx                 LLDPEnableCfg       `json:"enable-cfg,omitempty"`
	SystemCapabilities LLDPTxTlvsCfg       `json:"tx-tlvs-cfg,omitempty"`
	UUID               string              `json:"uuid,omitempty"`
}

func PostInterfaceEthernetLLDP(id string, name int, inst EthernetLLDP, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostInterfaceEthernetLLDP")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ethernet/"+strconv.Itoa(name)+"/lldp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m EthernetLLDP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostInterfaceEthernetLLDP REQ RES..........................", m)
			err := check_api_status("PostInterfaceEthernetLLDP", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetInterfaceEthernetLLDP(id string, name string, host string) (*EthernetLLDP, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetInterfaceEthernetLLDP")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ethernet/"+name+"/lldp", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m EthernetLLDP
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetInterfaceEthernetLLDP REQ RES..........................", m)
			err := check_api_status("GetInterfaceEthernetLLDP", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
