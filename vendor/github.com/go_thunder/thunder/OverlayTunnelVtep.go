package go_thunder

import (
	"bytes"
	//"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"

	"github.com/clarketm/json" // go get -u github.com/clarketm/json  ===> to avoide null dictionary in json
)

type OverlayTunnelVtep struct {
	ID OverlayTunnelVtepInstance `json:"vtep,omitempty"`
}

type OverlayTunnelVtepInstance struct {
	Encap       string                               `json:"encap,omitempty"`
	IPAddr      []OverlayTunnelHostList              `json:"host-list,omitempty"`
	ID          int                                  `json:"id,omitempty"`
	IPAddress   OverlayTunnelLocalIPAddress          `json:"local-ip-address,omitempty"`
	Ipv6Address OverlayTunnelLocalIpv6Address        `json:"local-ipv6-address,omitempty"`
	GreKey      []OverlayTunnelRemoteIPAddressList   `json:"remote-ip-address-list,omitempty"`
	Partition   []OverlayTunnelRemoteIpv6AddressList `json:"remote-ipv6-address-list,omitempty"`
	Counters1   []OverlayTunnelSamplingEnable        `json:"sampling-enable,omitempty"`
	UUID        string                               `json:"uuid,omitempty"`
	UserTag     string                               `json:"user-tag,omitempty"`
}

type OverlayTunnelHostList struct {
	IPAddr         string `json:"ip-addr,omitempty"`
	OverlayMacAddr string `json:"overlay-mac-addr,omitempty"`
	RemoteVtep     string `json:"remote-vtep,omitempty"`
	UUID           string `json:"uuid,omitempty"`
	Vni            int    `json:"vni,omitempty"`
}

type OverlayTunnelLocalIPAddress struct {
	IPAddress string                 `json:"ip-address,omitempty"`
	UUID      string                 `json:"uuid,omitempty"`
	Segment   []OverlayTunnelVniList `json:"vni-list,omitempty"`
}

type OverlayTunnelLocalIpv6Address struct {
	Ipv6Address string `json:"ipv6-address,omitempty"`
	UUID        string `json:"uuid,omitempty"`
}

type OverlayTunnelRemoteIPAddressList struct {
	Encap     string                    `json:"encap,omitempty"`
	RetryTime OverlayTunnelGreKeepalive `json:"gre-keepalive,omitempty"`
	IPAddress string                    `json:"ip-address,omitempty"`
	UUID      string                    `json:"uuid,omitempty"`
	GreKey    OverlayTunnelUseGreKey    `json:"use-gre-key,omitempty"`
	Partition OverlayTunnelUseLif       `json:"use-lif,omitempty"`
	UserTag   string                    `json:"user-tag,omitempty"`
	Segment   []OverlayTunnelVniList    `json:"vni-list,omitempty"`
}

type OverlayTunnelRemoteIpv6AddressList struct {
	Ipv6Address string              `json:"ipv6-address,omitempty"`
	UUID        string              `json:"uuid,omitempty"`
	Partition   OverlayTunnelUseLif `json:"use-lif,omitempty"`
	UserTag     string              `json:"user-tag,omitempty"`
}

type OverlayTunnelSamplingEnable struct {
	Counters1 string `json:"counters1,omitempty"`
}

type OverlayTunnelVniList struct {
	Segment int    `json:"segment,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}

type OverlayTunnelGreKeepalive struct {
	RetryCount int    `json:"retry-count,omitempty"`
	RetryTime  int    `json:"retry-time,omitempty"`
	UUID       string `json:"uuid,omitempty"`
}

type OverlayTunnelUseGreKey struct {
	GreKey int    `json:"gre-key,omitempty"`
	UUID   string `json:"uuid,omitempty"`
}

type OverlayTunnelUseLif struct {
	Lif       string `json:"lif,omitempty"`
	Partition string `json:"partition,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

func PostOverlayTunnelVtep(id string, inst OverlayTunnelVtep, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostOverlayTunnelVtep")

	payloadBytes, err := json.Marshal(inst)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/overlay-tunnel/vtep", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m OverlayTunnelVtep
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostOverlayTunnelVtep", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetOverlayTunnelVtep(id string, name1 string, host string) (*OverlayTunnelVtep, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetOverlayTunnelVtep")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/overlay-tunnel/vtep/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m OverlayTunnelVtep
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetOverlayTunnelVtep", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutOverlayTunnelVtep(id string, name1 string, inst OverlayTunnelVtep, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutOverlayTunnelVtep")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/overlay-tunnel/vtep/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m OverlayTunnelVtep
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutOverlayTunnelVtep", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteOverlayTunnelVtep(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteOverlayTunnelVtep")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/overlay-tunnel/vtep/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m OverlayTunnelVtep
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteOverlayTunnelVtep", data)
			if err != nil {
				return err
			}

		}
	}
	return nil
}
