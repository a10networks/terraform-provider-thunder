package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateTcpProxy struct {
	SlbTemplateTCPProxyInstanceName SlbTemplateTCPProxyInstance `json:"tcp-proxy,omitempty"`
}

type SlbTemplateTCPProxyInstance struct {
	SlbTemplateTCPProxyInstanceAckAggressiveness            string                                 `json:"ack-aggressiveness,omitempty"`
	SlbTemplateTCPProxyInstanceAliveIfActive                int                                    `json:"alive-if-active,omitempty"`
	SlbTemplateTCPProxyInstanceBackendWscale                int                                    `json:"backend-wscale,omitempty"`
	SlbTemplateTCPProxyInstanceDelSessionOnServerDown       int                                    `json:"del-session-on-server-down,omitempty"`
	SlbTemplateTCPProxyInstanceDisable                      int                                    `json:"disable,omitempty"`
	SlbTemplateTCPProxyInstanceDisableAbc                   int                                    `json:"disable-abc,omitempty"`
	SlbTemplateTCPProxyInstanceDisableSack                  int                                    `json:"disable-sack,omitempty"`
	SlbTemplateTCPProxyInstanceDisableTCPTimestamps         int                                    `json:"disable-tcp-timestamps,omitempty"`
	SlbTemplateTCPProxyInstanceDisableWindowScale           int                                    `json:"disable-window-scale,omitempty"`
	SlbTemplateTCPProxyInstanceDown                         int                                    `json:"down,omitempty"`
	SlbTemplateTCPProxyInstanceDynamicBufferAllocation      int                                    `json:"dynamic-buffer-allocation,omitempty"`
	SlbTemplateTCPProxyInstanceEarlyRetransmit              int                                    `json:"early-retransmit,omitempty"`
	SlbTemplateTCPProxyInstanceFinTimeout                   int                                    `json:"fin-timeout,omitempty"`
	SlbTemplateTCPProxyInstanceForceDeleteTimeout           int                                    `json:"force-delete-timeout,omitempty"`
	SlbTemplateTCPProxyInstanceForceDeleteTimeout100Ms      int                                    `json:"force-delete-timeout-100ms,omitempty"`
	SlbTemplateTCPProxyInstanceHalfCloseIdleTimeout         int                                    `json:"half-close-idle-timeout,omitempty"`
	SlbTemplateTCPProxyInstanceHalfOpenIdleTimeout          int                                    `json:"half-open-idle-timeout,omitempty"`
	SlbTemplateTCPProxyInstanceIdleTimeout                  int                                    `json:"idle-timeout,omitempty"`
	SlbTemplateTCPProxyInstanceInitCwnd                     int                                    `json:"init-cwnd,omitempty"`
	SlbTemplateTCPProxyInstanceInitialWindowSize            int                                    `json:"initial-window-size,omitempty"`
	SlbTemplateTCPProxyInstanceInsertClientIP               int                                    `json:"insert-client-ip,omitempty"`
	SlbTemplateTCPProxyInstanceInvalidRateLimit             int                                    `json:"invalid-rate-limit,omitempty"`
	SlbTemplateTCPProxyInstanceKeepaliveInterval            int                                    `json:"keepalive-interval,omitempty"`
	SlbTemplateTCPProxyInstanceKeepaliveProbes              int                                    `json:"keepalive-probes,omitempty"`
	SlbTemplateTCPProxyInstanceLimitedSlowstart             int                                    `json:"limited-slowstart,omitempty"`
	SlbTemplateTCPProxyInstanceMaxburst                     int                                    `json:"maxburst,omitempty"`
	SlbTemplateTCPProxyInstanceMinRto                       int                                    `json:"min-rto,omitempty"`
	SlbTemplateTCPProxyInstanceMss                          int                                    `json:"mss,omitempty"`
	SlbTemplateTCPProxyInstanceNagle                        int                                    `json:"nagle,omitempty"`
	SlbTemplateTCPProxyInstanceName                         string                                 `json:"name,omitempty"`
	SlbTemplateTCPProxyInstanceProxyHeaderProxyHeaderAction SlbTemplateTCPProxyInstanceProxyHeader `json:"proxy-header,omitempty"`
	SlbTemplateTCPProxyInstancePshFlagOptimization          int                                    `json:"psh-flag-optimization,omitempty"`
	SlbTemplateTCPProxyInstanceQos                          int                                    `json:"qos,omitempty"`
	SlbTemplateTCPProxyInstanceReassemblyLimit              int                                    `json:"reassembly-limit,omitempty"`
	SlbTemplateTCPProxyInstanceReassemblyTimeout            int                                    `json:"reassembly-timeout,omitempty"`
	SlbTemplateTCPProxyInstanceReceiveBuffer                int                                    `json:"receive-buffer,omitempty"`
	SlbTemplateTCPProxyInstanceReno                         int                                    `json:"reno,omitempty"`
	SlbTemplateTCPProxyInstanceResetFwd                     int                                    `json:"reset-fwd,omitempty"`
	SlbTemplateTCPProxyInstanceResetRev                     int                                    `json:"reset-rev,omitempty"`
	SlbTemplateTCPProxyInstanceRetransmitRetries            int                                    `json:"retransmit-retries,omitempty"`
	SlbTemplateTCPProxyInstanceServerDownAction             string                                 `json:"server-down-action,omitempty"`
	SlbTemplateTCPProxyInstanceSynRetries                   int                                    `json:"syn-retries,omitempty"`
	SlbTemplateTCPProxyInstanceTimewait                     int                                    `json:"timewait,omitempty"`
	SlbTemplateTCPProxyInstanceTransmitBuffer               int                                    `json:"transmit-buffer,omitempty"`
	SlbTemplateTCPProxyInstanceUUID                         string                                 `json:"uuid,omitempty"`
	SlbTemplateTCPProxyInstanceUserTag                      string                                 `json:"user-tag,omitempty"`
}

type SlbTemplateTCPProxyInstanceProxyHeader struct {
	SlbTemplateTCPProxyInstanceProxyHeaderProxyHeaderAction string `json:"proxy-header-action,omitempty"`
	SlbTemplateTCPProxyInstanceProxyHeaderVersion           string `json:"version,omitempty"`
}

func PostSlbTemplateTcpProxy(id string, inst SlbTemplateTcpProxy, host string) error {

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
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/tcp-proxy", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateTcpProxy
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateTcpProxy", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateTcpProxy(id string, name1 string, host string) (*SlbTemplateTcpProxy, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateTcpProxy")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/tcp-proxy/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateTcpProxy
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateTcpProxy", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateTcpProxy(id string, name1 string, inst SlbTemplateTcpProxy, host string) error {

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
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/tcp-proxy/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateTcpProxy
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateTcpProxy", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateTcpProxy(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateTcpProxy")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/tcp-proxy/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateTcpProxy
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateTcpProxy", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
