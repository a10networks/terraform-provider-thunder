package go_vthunder

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"util"
)

type Vtep struct {
	UUID VtepInstance `json:"vtep,omitempty"`
}
type SamplingEnableOT struct {
	Counters1 string `json:"counters1,omitempty"`
}
type VniList struct {
	Segment int    `json:"segment,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}
type SourceIPAddress struct {
	IPAddress string    `json:"ip-address,omitempty"`
	UUID      string    `json:"uuid,omitempty"`
	Segment   []VniList `json:"vni-list,omitempty"`
}
type HostList struct {
	DestinationVtep string `json:"destination-vtep,omitempty"`
	IPAddr          string `json:"ip-addr,omitempty"`
	OverlayMacAddr  string `json:"overlay-mac-addr,omitempty"`
	Vni             int    `json:"vni,omitempty"`
	UUID            string `json:"uuid,omitempty"`
}
type DestinationIPAddressList struct {
	UUID      string    `json:"uuid,omitempty"`
	IPAddress string    `json:"ip-address,omitempty"`
	Segment   []VniList `json:"vni-list,omitempty"`
	UserTag   string    `json:"user-tag,omitempty"`
	Encap     string    `json:"encap,omitempty"`
}
type VtepInstance struct {
	UUID      string                     `json:"uuid,omitempty"`
	UserTag   string                     `json:"user-tag,omitempty"`
	Counters1 []SamplingEnableOT         `json:"sampling-enable,omitempty"`
	IPAddress SourceIPAddress            `json:"source-ip-address,omitempty"`
	Encap     string                     `json:"encap,omitempty"`
	IPAddr    []HostList                 `json:"host-list,omitempty"`
	ID        int                        `json:"id,omitempty"`
	Segment   []DestinationIPAddressList `json:"destination-ip-address-list,omitempty"`
}

func GetVtep(id string, name string, host string) (*Vtep, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/overlay-tunnel/vtep/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Vtep
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
			return &m, nil
		}
	}
}

func PostVtep(id string, sg Vtep, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(sg)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/overlay-tunnel/vtep/", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Vtep
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}

func PutVtep(id string, name string, sg Vtep, host string) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(sg)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/overlay-tunnel/vtep/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Vtep
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
		}
	}

}

func DeleteVtep(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/overlay-tunnel/vtep/"+name, nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Vtep
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
		}
	}
	return nil
}
