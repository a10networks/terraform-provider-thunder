package thunder

import (
    "bytes"
    "github.com/clarketm/json"
    "io/ioutil"
    "util"
)

//based on ACOS 5_2_1-P3_70
type EndpointIpRouteStaticBfd struct { Inst IpRouteStaticBfd `json:"bfd"` }
type IpRouteStaticBfd struct {
    Action string `json:"action,omitempty"`
    LocalIp string `json:"local-ip,omitempty"`
    NexthopIp string `json:"nexthop-ip,omitempty"`
    Template string `json:"template,omitempty"`
    Threshold int `json:"threshold,omitempty"`
    Uuid string `json:"uuid,omitempty"`
}

func (p *EndpointIpRouteStaticBfd) GetId() string{
    return p.Inst.LocalIp + "+" + p.Inst.NexthopIp
}

func PostEndpointIpRouteStaticBfd(auth_token string, inst EndpointIpRouteStaticBfd, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PostEndpointIpRouteStaticBfd")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("POST", "https://" + host + "/axapi/v3/ip/route/static/bfd", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpRouteStaticBfd
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PostEndpointIpRouteStaticBfd", data)
        }
    }
    return err
}

func GetEndpointIpRouteStaticBfd(auth_token string, host string, instId string) (*EndpointIpRouteStaticBfd, error) {
    logger := util.GetLoggerInstance()
    logger.Println("GetEndpointIpRouteStaticBfd")
    headers := GetReqHeader(auth_token)

    resp, err := doHttp("GET", "https://"+ host + "/axapi/v3/ip/route/static/bfd/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpRouteStaticBfd
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("GetEndpointIpRouteStaticBfd", data)
            if err == nil {
                return &m, nil
            }
        }
    }
    return nil, err
}

func PutEndpointIpRouteStaticBfd(auth_token string, inst EndpointIpRouteStaticBfd, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PutEndpointIpRouteStaticBfd")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("PUT", "https://" + host + "/axapi/v3/ip/route/static/bfd/" + inst.GetId(), bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpRouteStaticBfd
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PutEndpointIpRouteStaticBfd", data)
        }
    }
    return err
}

func DeleteEndpointIpRouteStaticBfd(auth_token string, host string, instId string) error {
    logger := util.GetLoggerInstance()
    logger.Println("DeleteEndpointIpRouteStaticBfd")
    headers := GetReqHeader(auth_token)

    resp, err := doHttp("DELETE", "https://" + host + "/axapi/v3/ip/route/static/bfd/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpRouteStaticBfd
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("DeleteEndpointIpRouteStaticBfd", data)
        }
    }
    return err
}
