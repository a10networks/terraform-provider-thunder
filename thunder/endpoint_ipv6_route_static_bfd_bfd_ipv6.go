package thunder

import (
    "bytes"
    "github.com/clarketm/json"
    "io/ioutil"
    "util"
)

//based on ACOS 5_2_1-P3_70
type EndpointIpv6RouteStaticBfdBfdIpv6 struct { Inst Ipv6RouteStaticBfdBfdIpv6 `json:"bfd-ipv6"` }
type Ipv6RouteStaticBfdBfdIpv6 struct {
    Action string `json:"action,omitempty"`
    LocalIpv6 string `json:"local-ipv6,omitempty"`
    NexthopIpv6 string `json:"nexthop-ipv6,omitempty"`
    Template string `json:"template,omitempty"`
    Threshold int `json:"threshold,omitempty"`
    Uuid string `json:"uuid,omitempty"`
}
func (p *EndpointIpv6RouteStaticBfdBfdIpv6) GetId() string{
    return p.Inst.LocalIpv6 + "+" + p.Inst.NexthopIpv6
}

func PostEndpointIpv6RouteStaticBfdBfdIpv6(auth_token string, inst EndpointIpv6RouteStaticBfdBfdIpv6, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PostEndpointIpv6RouteStaticBfdBfdIpv6")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("POST", "https://" + host + "/axapi/v3/ipv6/route/static/bfd/bfd-ipv6", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6RouteStaticBfdBfdIpv6
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PostEndpointIpv6RouteStaticBfdBfdIpv6", data)
        }
    }
    return err
}

func GetEndpointIpv6RouteStaticBfdBfdIpv6(auth_token string, host string, instId string) (*EndpointIpv6RouteStaticBfdBfdIpv6, error) {
    logger := util.GetLoggerInstance()
    logger.Println("GetEndpointIpv6RouteStaticBfdBfdIpv6")
    headers := GetReqHeader(auth_token)

    resp, err := doHttp("GET", "https://"+ host + "/axapi/v3/ipv6/route/static/bfd/bfd-ipv6/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6RouteStaticBfdBfdIpv6
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("GetEndpointIpv6RouteStaticBfdBfdIpv6", data)
            if err == nil {
                return &m, nil
            }
        }
    }
    return nil, err
}

func PutEndpointIpv6RouteStaticBfdBfdIpv6(auth_token string, inst EndpointIpv6RouteStaticBfdBfdIpv6, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PutEndpointIpv6RouteStaticBfdBfdIpv6")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("PUT", "https://" + host + "/axapi/v3/ipv6/route/static/bfd/bfd-ipv6/" + inst.GetId(), bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6RouteStaticBfdBfdIpv6
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PutEndpointIpv6RouteStaticBfdBfdIpv6", data)
        }
    }
    return err
}

func DeleteEndpointIpv6RouteStaticBfdBfdIpv6(auth_token string, host string, instId string) error {
    logger := util.GetLoggerInstance()
    logger.Println("DeleteEndpointIpv6RouteStaticBfdBfdIpv6")
    headers := GetReqHeader(auth_token)

    resp, err := doHttp("DELETE", "https://" + host + "/axapi/v3/ipv6/route/static/bfd/bfd-ipv6/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6RouteStaticBfdBfdIpv6
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("DeleteEndpointIpv6RouteStaticBfdBfdIpv6", data)
        }
    }
    return err
}
