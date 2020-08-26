package go_vthunder

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "util"
)

type Fw struct {
	Ftp FwAlgInstance  `json:"alg,omitempty"`
}

type FwAlgInstance struct {
	DefaultPortDisable FwDNS  `json:"dns,omitempty"`
	DefaultPortDisable FwFtp  `json:"ftp,omitempty"`
	Disable FwIcmp  `json:"icmp,omitempty"`
	DefaultPortDisable FwFtp  `json:"pptp,omitempty"`
	DefaultPortDisable FwFtp  `json:"rtsp,omitempty"`
	DefaultPortDisable FwFtp  `json:"sip,omitempty"`
	DefaultPortDisable FwFtp  `json:"tftp,omitempty"`
	UUID string  `json:"uuid,omitempty"`
}

type FwDNS struct {
	DefaultPortDisable string  `json:"default-port-disable,omitempty"`
	UUID string  `json:"uuid,omitempty"`
}

type FwFtp struct {
	DefaultPortDisable string  `json:"default-port-disable,omitempty"`
	Counters1 []FwSamplingEnable  `json:"sampling-enable,omitempty"`
	UUID string  `json:"uuid,omitempty"`
}

type FwIcmp struct {
	Disable string  `json:"disable,omitempty"`
	UUID string  `json:"uuid,omitempty"`
}

type FwSamplingEnable struct {
	Counters1 string  `json:"counters1,omitempty"`
}

func PostFwAlg(id string, inst Fw, host string)  { 

    logger := util.GetLoggerInstance()

    var headers = make(map[string]string)
    headers["Accept"] = "application/json"
    headers["Content-Type"] = "application/json"
    headers["Authorization"] = id
    logger.Println("[INFO] Inside PostFwAlg")
    payloadBytes, err := json.Marshal(inst)
    logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
    if err != nil {
        logger.Println("[INFO] Marshalling failed with error ", err)
    }

    resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/alg", bytes.NewReader(payloadBytes), headers)

    if err != nil {
        logger.Println("The HTTP request failed with error ", err)
        
    } else {
        data, _ := ioutil.ReadAll(resp.Body)
        var m Fw
        erro := json.Unmarshal(data, &m)
        if erro != nil {
            logger.Println("Unmarshal error ", err)
            
        } else {
            logger.Println("[INFO] GET REQ RES..........................", m)
            
        }
    }
    
}

func GetFwAlg(id string, name string, host string) (*Fw, error) { 

    logger := util.GetLoggerInstance()

    var headers = make(map[string]string)
    headers["Accept"] = "application/json"
    headers["Content-Type"] = "application/json"
    headers["Authorization"] = id
    logger.Println("[INFO] Inside GetFwAlg")
    
    resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/alg/", nil, headers)

    if err != nil {
        logger.Println("The HTTP request failed with error ", err)
        return nil, err
    } else {
        data, _ := ioutil.ReadAll(resp.Body)
        var m Fw
        erro := json.Unmarshal(data, &m)
        if erro != nil {
            logger.Println("Unmarshal error ", err)
            return nil, err
        } else {
            logger.Println("[INFO] GET REQ RES..........................", m)
            return &m,nil
        }
    }
    
}
