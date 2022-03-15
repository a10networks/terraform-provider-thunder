package thunder

import (
    "bytes"
    "github.com/clarketm/json"
    "io/ioutil"
    "util"
    "net/url"
)

//based on ACOS 5_2_1-P3_70
type EndpointIpRouteRib struct { Inst IpRouteRib `json:"rib"` }
type IpRouteRib struct {
    IpDestAddr string `json:"ip-dest-addr,omitempty"`
    IpMask string `json:"ip-mask,omitempty"`
    IpNexthopIpv4 []IpRouteRibIpNexthopIpv4 `json:"ip-nexthop-ipv4,omitempty"`
    IpNexthopLif []IpRouteRibIpNexthopLif `json:"ip-nexthop-lif,omitempty"`
    IpNexthopPartition []IpRouteRibIpNexthopPartition `json:"ip-nexthop-partition,omitempty"`
    IpNexthopTunnel []IpRouteRibIpNexthopTunnel `json:"ip-nexthop-tunnel,omitempty"`
    Uuid string `json:"uuid,omitempty"`
}
func (p *EndpointIpRouteRib) GetId() string {
    //axapi does not want '+' be QueryEscaped
    return p.Inst.IpDestAddr + "+" + url.QueryEscape(p.Inst.IpMask)
}

type IpRouteRibIpNexthopIpv4 struct {
    DescriptionNexthopIp string `json:"description-nexthop-ip,omitempty"`
    DistanceNexthopIp int `json:"distance-nexthop-ip,omitempty"`
    IpNextHop string `json:"ip-next-hop,omitempty"`
}

type IpRouteRibIpNexthopLif struct {
    DescriptionNexthopLif string `json:"description-nexthop-lif,omitempty"`
    Lif string `json:"lif,omitempty"`
}

type IpRouteRibIpNexthopPartition struct {
    DescriptionNexthopPartition string `json:"description-nexthop-partition,omitempty"`
    DescriptionPartitionVrid string `json:"description-partition-vrid,omitempty"`
    PartitionName string `json:"partition-name,omitempty"`
    VridNumInPartition int `json:"vrid-num-in-partition,omitempty"`
}

type IpRouteRibIpNexthopTunnel struct {
    DescriptionNexthopTunnel string `json:"description-nexthop-tunnel,omitempty"`
    DistanceNexthopTunnel int `json:"distance-nexthop-tunnel,omitempty"`
    IpNextHopTunnel string `json:"ip-next-hop-tunnel,omitempty"`
    Tunnel int `json:"tunnel,omitempty"`
}

func PostEndpointIpRouteRib(auth_token string, inst EndpointIpRouteRib, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PostIpRouteRib")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload:", string(payloadBytes))
    resp, err := doHttp("POST", "https://" + host + "/axapi/v3/ip/route/rib", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpRouteRib
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PostIpRouteRib", data)
        }
    }
    return err
}

func GetEndpointIpRouteRib(auth_token string, host string, instId string) (*EndpointIpRouteRib, error) {
    logger := util.GetLoggerInstance()
    logger.Println("GetIpRouteRib")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("GET", "https://"+ host + "/axapi/v3/ip/route/rib/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpRouteRib
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("GetIpRouteRib", data)
            if err == nil {
                return &m, nil
            }
        }
    }
    return nil, err
}

func PutEndpointIpRouteRib(auth_token string, inst EndpointIpRouteRib, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PutIpRouteRib")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("PUT", "https://" + host + "/axapi/v3/ip/route/rib/" + inst.GetId(), bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpRouteRib
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PutIpRouteRib", data)
        }
    }
    return err
}

func DeleteEndpointIpRouteRib(auth_token string, host string, instId string) error {
    logger := util.GetLoggerInstance()
    logger.Println("DeleteIpRouteRib")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("DELETE", "https://" + host + "/axapi/v3/ip/route/rib/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpRouteRib
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("DeleteIpRouteRib", data)
        }
    }
    return err
}
