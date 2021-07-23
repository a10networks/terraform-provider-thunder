package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type TCPProxy struct {
	UUID TCPProxyInstance `json:"tcp-proxy,omitempty"`
}

type TCPProxyInstance struct {
	Qos                     int    `json:"qos,omitempty"`
	InitCwnd                int    `json:"init-cwnd,omitempty"`
	IdleTimeout             int    `json:"idle-timeout,omitempty"`
	FinTimeout              int    `json:"fin-timeout,omitempty"`
	HalfOpenIdleTimeout     int    `json:"half-open-idle-timeout,omitempty"`
	Reno                    int    `json:"reno,omitempty"`
	Down                    int    `json:"down,omitempty"`
	EarlyRetransmit         int    `json:"early-retransmit,omitempty"`
	ServerDownAction        string `json:"server-down-action,omitempty"`
	Timewait                int    `json:"timewait,omitempty"`
	MinRto                  int    `json:"min-rto,omitempty"`
	DynamicBufferAllocation int    `json:"dynamic-buffer-allocation,omitempty"`
	LimitedSlowstart        int    `json:"limited-slowstart,omitempty"`
	DisableSack             int    `json:"disable-sack,omitempty"`
	DisableWindowScale      int    `json:"disable-window-scale,omitempty"`
	AliveIfActive           int    `json:"alive-if-active,omitempty"`
	Mss                     int    `json:"mss,omitempty"`
	KeepaliveInterval       int    `json:"keepalive-interval,omitempty"`
	RetransmitRetries       int    `json:"retransmit-retries,omitempty"`
	InsertClientIP          int    `json:"insert-client-ip,omitempty"`
	TransmitBuffer          int    `json:"transmit-buffer,omitempty"`
	Nagle                   int    `json:"nagle,omitempty"`
	ForceDeleteTimeout100Ms int    `json:"force-delete-timeout-100ms,omitempty"`
	InitialWindowSize       int    `json:"initial-window-size,omitempty"`
	KeepaliveProbes         int    `json:"keepalive-probes,omitempty"`
	PshFlagOptimization     int    `json:"psh-flag-optimization,omitempty"`
	AckAggressiveness       string `json:"ack-aggressiveness,omitempty"`
	BackendWscale           int    `json:"backend-wscale,omitempty"`
	Disable                 int    `json:"disable,omitempty"`
	ResetRev                int    `json:"reset-rev,omitempty"`
	Maxburst                int    `json:"maxburst,omitempty"`
	UUID                    string `json:"uuid,omitempty"`
	ReceiveBuffer           int    `json:"receive-buffer,omitempty"`
	DelSessionOnServerDown  int    `json:"del-session-on-server-down,omitempty"`
	Name                    string `json:"name,omitempty"`
	ReassemblyTimeout       int    `json:"reassembly-timeout,omitempty"`
	ResetFwd                int    `json:"reset-fwd,omitempty"`
	DisableTCPTimestamps    int    `json:"disable-tcp-timestamps,omitempty"`
	SynRetries              int    `json:"syn-retries,omitempty"`
	ForceDeleteTimeout      int    `json:"force-delete-timeout,omitempty"`
	UserTag                 string `json:"user-tag,omitempty"`
	ReassemblyLimit         int    `json:"reassembly-limit,omitempty"`
	InvalidRateLimit        int    `json:"invalid-rate-limit,omitempty"`
	DisableAbc              int    `json:"disable-abc,omitempty"`
	HalfCloseIdleTimeout    int    `json:"half-close-idle-timeout,omitempty"`
}

func PostSlbTemplateTcpProxy(id string, inst TCPProxy, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateTcpProxy")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/tcp-proxy", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m TCPProxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateTcpProxy REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateTcpProxy", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func GetSlbTemplateTcpProxy(id string, name string, host string) (*TCPProxy, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateTcpProxy")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/tcp-proxy/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m TCPProxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateTcpProxy REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateTcpProxy", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateTcpProxy(id string, name string, inst TCPProxy, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateTcpProxy")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/tcp-proxy/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m TCPProxy
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateTcpProxy REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateTcpProxy", data)
			if err != nil {
				return err
			}

		}
	}
return err
}

func DeleteSlbTemplateTcpProxy(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateTcpProxy")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/tcp-proxy/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m TCPProxy
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
