package thunder
import (
    "bytes"
    "github.com/clarketm/json"
    "io/ioutil"
    "util"
)

//based on ACOS 5_2_1-P3_70
type EndpointLoggingHostIpv4addr struct { Inst LoggingHostIpv4addr `json:"ipv4addr"` }
type LoggingHostIpv4addr struct {
    HostIpv4 string `json:"host-ipv4,omitempty"`
    OverTls int `json:"over-tls,omitempty"`
    Port int `json:"port,omitempty"`
    Tcp int `json:"tcp,omitempty"`
    UseMgmtPort int `json:"use-mgmt-port,omitempty"`
    Uuid string `json:"uuid,omitempty"`
}

func (p *EndpointLoggingHostIpv4addr) GetId() string{
    return p.Inst.HostIpv4
}

func PostEndpointLoggingHostIpv4addr(auth_token string, inst EndpointLoggingHostIpv4addr, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PostLoggingHostIpv4addr")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload:", string(payloadBytes))
    resp, err := doHttp("POST", "https://" + host + "/axapi/v3/logging/host/ipv4addr", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointLoggingHostIpv4addr
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PostLoggingHostIpv4addr", data)
        }
    }
    return err
}

func GetEndpointLoggingHostIpv4addr(auth_token string, host string, instId string) (*EndpointLoggingHostIpv4addr, error) {
    logger := util.GetLoggerInstance()
    logger.Println("GetLoggingHostIpv4addr")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("GET", "https://"+ host + "/axapi/v3/logging/host/ipv4addr/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointLoggingHostIpv4addr
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("GetLoggingHostIpv4addr", data)
            if err == nil {
                return &m, nil
            }
        }
    }
    return nil, err
}

func PutEndpointLoggingHostIpv4addr(auth_token string, inst EndpointLoggingHostIpv4addr, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PutLoggingHostIpv4addr")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("PUT", "https://" + host + "/axapi/v3/logging/host/ipv4addr/" + inst.GetId(), bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointLoggingHostIpv4addr
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PutLoggingHostIpv4addr", data)
        }
    }
    return err
}

func DeleteEndpointLoggingHostIpv4addr(auth_token string, host string, instId string) error {
    logger := util.GetLoggerInstance()
    logger.Println("DeleteLoggingHostIpv4addr")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("DELETE", "https://" + host + "/axapi/v3/logging/host/ipv4addr/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", data)
        var m EndpointLoggingHostIpv4addr
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("DeleteLoggingHostIpv4addr", data)
        }
    }
    return err
}
