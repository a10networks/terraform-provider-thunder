package go_thunder

import (
	"bytes"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"strconv"
	"util"
)

type IP struct {
	UUID        string    `json:"uuid,omitempty"`
	Dhcp        int       `json:"dhcp,omitempty"`
	Ipv4Address []Address `json:"address-list,omitempty"`
}

type Address struct {
	Ipv4Address string `json:"ipv4-address,omitempty"`
	Ipv4Netmask string `json:"ipv4-netmask,omitempty"`
}

type EthernetInstance struct {
	FecForcedOn int `json:"fec-forced-on,omitempty"`
	TrapSource  int `json:"trap-source,omitempty"`
	Dhcp        IP  `json:"ip,omitempty"`

	UUID  string `json:"uuid,omitempty"`
	Speed string `json:"speed,omitempty"`

	MediaTypeCopper  int `json:"media-type-copper,omitempty"`
	Ifnum            int `json:"ifnum,omitempty"`
	L3VlanFwdDisable int `json:"l3-vlan-fwd-disable,omitempty"`
	RemoveVlanTag    int `json:"remove-vlan-tag,omitempty"`
	MonitorList      []struct {
		MonitorVlan int    `json:"monitor-vlan,omitempty"`
		Monitor     string `json:"monitor,omitempty"`
		MirrorIndex int    `json:"mirror-index,omitempty"`
	} `json:"monitor-list,omitempty"`
	CPUProcess    int `json:"cpu-process,omitempty"`
	AutoNegEnable int `json:"auto-neg-enable,omitempty"`
	//	Map           struct {
	//		Inside      int    `json:"inside,omitempty"`
	//		MapTInside  int    `json:"map-t-inside,omitempty"`
	//		UUID        string `json:"uuid,omitempty"`
	//		MapTOutside int    `json:"map-t-outside,omitempty"`
	//		Outside     int    `json:"outside,omitempty"`
	//	} `json:"map,omitempty"`
	TrafficDistributionMode string `json:"traffic-distribution-mode,omitempty"`
	TrunkGroupList          []struct {
		UUID           string `json:"uuid,omitempty"`
		TrunkNumber    int    `json:"trunk-number,omitempty"`
		UserTag        string `json:"user-tag,omitempty"`
		UdldTimeoutCfg struct {
			Slow int `json:"slow,omitempty"`
			Fast int `json:"fast,omitempty"`
		} `json:"udld-timeout-cfg,omitempty"`
		Mode         string `json:"mode,omitempty"`
		Timeout      string `json:"timeout,omitempty"`
		Type         string `json:"type,omitempty"`
		AdminKey     int    `json:"admin-key,omitempty"`
		PortPriority int    `json:"port-priority,omitempty"`
	} `json:"trunk-group-list,omitempty"`

	CPUProcessDir string `json:"cpu-process-dir,omitempty"`

	Name      string `json:"name,omitempty"`
	Duplexity string `json:"duplexity,omitempty"`

	UserTag string `json:"user-tag,omitempty"`
	Mtu     int    `json:"mtu,omitempty"`

	SamplingEnable []struct {
		Counters1 string `json:"counters1,omitempty"`
	} `json:"sampling-enable,omitempty"`
	LoadInterval int `json:"load-interval,omitempty"`

	Action       string `json:"action,omitempty"`
	FecForcedOff int    `json:"fec-forced-off,omitempty"`

	FlowControl int `json:"flow-control,omitempty"`
}

type Ethernet struct {
	UUID EthernetInstance `json:"ethernet,omitempty"`
}

type EthList struct {
	UUID []EthernetInstance `json:"ethernet-list,omitempty"`
}

func postEthernet(id string, ip map[int]IP, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	for k, v := range ip {

		logger.Println("[INFO] Iteration - " + strconv.Itoa(k))
		logger.Println("[INFO] Headers - " + headers["Accept"] + "," + headers["Content-Type"] + "," + headers["Authorization"])

		payloadBytes, err := json.Marshal(v)

		logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

		if err != nil {
			logger.Println("[INFO] Marshalling failed with error\n ", err)
		}
		resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/interface/ethernet/"+strconv.Itoa(k)+"/ip", bytes.NewReader(payloadBytes), headers)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			logger.Println("The HTTP request failed with error \n", err)
			return err
		} else {
			data, _ := ioutil.ReadAll(resp.Body)
			var m ServiceGroup
			erro := json.Unmarshal(data, &m)
			if erro != nil {
				fmt.Printf("Unmarshal error %s\n", err)
			} else {
				fmt.Println("response Body:", string(data))
				logger.Println("response Body:", string(data))
				err := check_api_status("postEthernet", data)
				if err != nil {
					return err
				}
			}
		}

	}
	return nil
}

func PutEthernet(id string, ethernet map[int]Ethernet, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	for k, v := range ethernet {

		logger.Println("[INFO] Iteration - " + strconv.Itoa(k))
		logger.Println("[INFO] Headers - " + headers["Accept"] + "," + headers["Content-Type"] + "," + headers["Authorization"])

		payloadBytes, err := json.Marshal(v)

		logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

		if err != nil {
			logger.Println("[INFO] Marshalling failed with error \n", err)
		}
		resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/interface/ethernet/"+strconv.Itoa(k), bytes.NewReader(payloadBytes), headers)

		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
			logger.Println("The HTTP request failed with error \n", err)
			return err
		} else {
			data, _ := ioutil.ReadAll(resp.Body)

			logger.Println("HTTP Response -", string(data))
			err := check_api_status("PutEthernet", data)
			if err != nil {
				return err
			}

		}

	}
	return nil
}

func GetEthernet(id string, name string, host string) (*Ethernet, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/interface/ethernet/"+name+"/ip", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Ethernet
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetEthernet", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
