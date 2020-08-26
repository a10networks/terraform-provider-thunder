package go_vthunder

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "util"
)

type FwAlg struct {
	DefaultPortDisable FwAlgFtpInstance  `json:"ftp,omitempty"`
}

type FwAlgFtpInstance struct {
	DefaultPortDisable string  `json:"default-port-disable,omitempty"`
	Counters1 []FwAlgSamplingEnable  `json:"sampling-enable,omitempty"`
	UUID string  `json:"uuid,omitempty"`
}

type FwAlgSamplingEnable struct {
	Counters1 string  `json:"counters1,omitempty"`
}

func PostFwAlgFtp(id string, inst FwAlg, host string)  { 

    logger := util.GetLoggerInstance()

    var headers = make(map[string]string)
    headers["Accept"] = "application/json"
    headers["Content-Type"] = "application/json"
    headers["Authorization"] = id
    logger.Println("[INFO] Inside PostFwAlgFtp")
    payloadBytes, err := json.Marshal(inst)
    logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
    if err != nil {
        logger.Println("[INFO] Marshalling failed with error ", err)
    }

    resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/alg/ftp", bytes.NewReader(payloadBytes), headers)

    if err != nil {
        logger.Println("The HTTP request failed with error ", err)
        
    } else {
        data, _ := ioutil.ReadAll(resp.Body)
        var m FwAlg
        erro := json.Unmarshal(data, &m)
        if erro != nil {
            logger.Println("Unmarshal error ", err)
            
        } else {
            logger.Println("[INFO] GET REQ RES..........................", m)
            
        }
    }
    
}

func GetFwAlgFtp(id string, name string, host string) (*FwAlg, error) { 

    logger := util.GetLoggerInstance()

    var headers = make(map[string]string)
    headers["Accept"] = "application/json"
    headers["Content-Type"] = "application/json"
    headers["Authorization"] = id
    logger.Println("[INFO] Inside GetFwAlgFtp")
    
    resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/alg/ftp/", nil, headers)

    if err != nil {
        logger.Println("The HTTP request failed with error ", err)
        return nil, err
    } else {
        data, _ := ioutil.ReadAll(resp.Body)
        var m FwAlg
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
