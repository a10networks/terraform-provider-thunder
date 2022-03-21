package thunder

import (
    "bytes"
    "github.com/clarketm/json"
    "io/ioutil"
    "util"
    "strconv"
)

//based on ACOS 5_2_1-P3_70
type EndpointIpv6RouteStaticBfdVe struct { Inst Ipv6RouteStaticBfdVe `json:"ve"` }
type Ipv6RouteStaticBfdVe struct {
    Action string `json:"action,omitempty"`
    VeNum int `json:"ve-num,omitempty"`
    NexthopIpv6Ll string `json:"nexthop-ipv6-ll,omitempty"`
    Template string `json:"template,omitempty"`
    Threshold int `json:"threshold,omitempty"`
    Uuid string `json:"uuid,omitempty"`
}
func (p *EndpointIpv6RouteStaticBfdVe) GetId() string{
    return strconv.Itoa(p.Inst.VeNum) + "+" + p.Inst.NexthopIpv6Ll
}

func PostEndpointIpv6RouteStaticBfdVe(auth_token string, inst EndpointIpv6RouteStaticBfdVe, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PostIpv6RouteStaticBfdVe")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload:", string(payloadBytes))
    resp, err := doHttp("POST", "https://" + host + "/axapi/v3/ipv6/route/static/bfd/ve", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6RouteStaticBfdVe
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PostIpv6RouteStaticBfdVe", data)
        }
    }
    return err
}

func GetEndpointIpv6RouteStaticBfdVe(auth_token string, host string, instId string) (*EndpointIpv6RouteStaticBfdVe, error) {
    logger := util.GetLoggerInstance()
    logger.Println("GetIpv6RouteStaticBfdVe")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("GET", "https://"+ host + "/axapi/v3/ipv6/route/static/bfd/ve/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6RouteStaticBfdVe
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("GetIpv6RouteStaticBfdVe", data)
            if err == nil {
                return &m, nil
            }
        }
    }
    return nil, err
}

func PutEndpointIpv6RouteStaticBfdVe(auth_token string, inst EndpointIpv6RouteStaticBfdVe, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PutIpv6RouteStaticBfdVe")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("PUT", "https://" + host + "/axapi/v3/ipv6/route/static/bfd/ve/" + inst.GetId(), bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6RouteStaticBfdVe
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PutIpv6RouteStaticBfdVe", data)
        }
    }
    return err
}

func DeleteEndpointIpv6RouteStaticBfdVe(auth_token string, host string, instId string) error {
    logger := util.GetLoggerInstance()
    logger.Println("DeleteIpv6RouteStaticBfdVe")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("DELETE", "https://" + host + "/axapi/v3/ipv6/route/static/bfd/ve/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", data)
        var m EndpointIpv6RouteStaticBfdVe
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("DeleteIpv6RouteStaticBfdVe", data)
        }
    }
    return err
}
