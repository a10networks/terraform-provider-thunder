package thunder
import (
    "bytes"
    "github.com/clarketm/json"
    "io/ioutil"
    "util"
    "net/url"
)

//based on ACOS 5_2_1-P3_70
type EndpointIpv6NatPool struct { Inst Ipv6NatPool `json:"pool"` }
type Ipv6NatPool struct {
    EndAddress string `json:"end-address,omitempty"`
    Gateway string `json:"gateway,omitempty"`
    IpRr int `json:"ip-rr,omitempty"`
    Netmask int `json:"netmask,omitempty"`
    PoolName string `json:"pool-name,omitempty"`
    PortOverload int `json:"port-overload,omitempty"`
    SamplingEnable []Ipv6NatPoolSamplingEnable `json:"sampling-enable,omitempty"`
    ScaleoutDeviceId int `json:"scaleout-device-id,omitempty"`
    StartAddress string `json:"start-address,omitempty"`
    Uuid string `json:"uuid,omitempty"`
    Vrid int `json:"vrid,omitempty"`
}

func (p *EndpointIpv6NatPool) GetId() string{
    return url.QueryEscape(p.Inst.PoolName)
}

type Ipv6NatPoolSamplingEnable struct {
    Counters1 string `json:"counters1,omitempty"`
}

func PostEndpointIpv6NatPool(auth_token string, inst EndpointIpv6NatPool, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PostIpv6NatPool")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload:", string(payloadBytes))
    resp, err := doHttp("POST", "https://" + host + "/axapi/v3/ipv6/nat/pool", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6NatPool
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PostIpv6NatPool", data)
        }
    }
    return err
}

func GetEndpointIpv6NatPool(auth_token string, host string, instId string) (*EndpointIpv6NatPool, error) {
    logger := util.GetLoggerInstance()
    logger.Println("GetIpv6NatPool")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("GET", "https://"+ host + "/axapi/v3/ipv6/nat/pool/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6NatPool
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("GetIpv6NatPool", data)
            if err == nil {
                return &m, nil
            }
        }
    }
    return nil, err
}

func PutEndpointIpv6NatPool(auth_token string, inst EndpointIpv6NatPool, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PutIpv6NatPool")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("PUT", "https://" + host + "/axapi/v3/ipv6/nat/pool/" + inst.GetId(), bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6NatPool
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PutIpv6NatPool", data)
        }
    }
    return err
}

func DeleteEndpointIpv6NatPool(auth_token string, host string, instId string) error {
    logger := util.GetLoggerInstance()
    logger.Println("DeleteIpv6NatPool")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("DELETE", "https://" + host + "/axapi/v3/ipv6/nat/pool/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", data)
        var m EndpointIpv6NatPool
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("DeleteIpv6NatPool", data)
        }
    }
    return err
}
