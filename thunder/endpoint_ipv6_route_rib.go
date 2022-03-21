package thunder
import (
    "bytes"
    "github.com/clarketm/json"
    "io/ioutil"
    "util"
    "net/url"
)

//based on ACOS 5_2_1-P3_70
type EndpointIpv6RouteRib struct { Inst Ipv6RouteRib `json:"rib"` }
type Ipv6RouteRib struct {
    Ipv6Address string `json:"ipv6-address,omitempty"`
    Ipv6NexthopIpv6 []Ipv6RouteRibIpv6NexthopIpv6 `json:"ipv6-nexthop-ipv6,omitempty"`
    Ipv6NexthopTunnel []Ipv6RouteRibIpv6NexthopTunnel `json:"ipv6-nexthop-tunnel,omitempty"`
    Uuid string `json:"uuid,omitempty"`
}

type Ipv6RouteRibIpv6NexthopIpv6 struct {
    Description string `json:"description,omitempty"`
    Distance int `json:"distance,omitempty"`
    Ethernet int `json:"ethernet,omitempty"`
    Ipv6Nexthop string `json:"ipv6-nexthop,omitempty"`
    Trunk int `json:"trunk,omitempty"`
    Ve int `json:"ve,omitempty"`
}

type Ipv6RouteRibIpv6NexthopTunnel struct {
    Description string `json:"description,omitempty"`
    DistanceNexthopTunnel int `json:"distance-nexthop-tunnel,omitempty"`
    Ipv6NexthopTunnelAddr string `json:"ipv6-nexthop-tunnel-addr,omitempty"`
    Tunnel int `json:"tunnel,omitempty"`
}

func (p *EndpointIpv6RouteRib) GetId() string{
    return url.QueryEscape(p.Inst.Ipv6Address)
}

func PostEndpointIpv6RouteRib(auth_token string, inst EndpointIpv6RouteRib, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PostIpv6RouteRib")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload:", string(payloadBytes))
    resp, err := doHttp("POST", "https://" + host + "/axapi/v3/ipv6/route/rib", bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6RouteRib
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PostIpv6RouteRib", data)
        }
    }
    return err
}

func GetEndpointIpv6RouteRib(auth_token string, host string, instId string) (*EndpointIpv6RouteRib, error) {
    logger := util.GetLoggerInstance()
    logger.Println("GetIpv6RouteRib")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("GET", "https://"+ host + "/axapi/v3/ipv6/route/rib/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6RouteRib
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("GetIpv6RouteRib", data)
            if err == nil {
                return &m, nil
            }
        }
    }
    return nil, err
}

func PutEndpointIpv6RouteRib(auth_token string, inst EndpointIpv6RouteRib, host string) error {
    logger := util.GetLoggerInstance()
    logger.Println("PutIpv6RouteRib")
    headers := GetReqHeader(auth_token)
    payloadBytes, err := json.Marshal(inst)
    if err != nil {
        logger.Println("json.Marshal() failed with error", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))

    resp, err := doHttp("PUT", "https://" + host + "/axapi/v3/ipv6/route/rib/" + inst.GetId(), bytes.NewReader(payloadBytes), headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", string(data))
        var m EndpointIpv6RouteRib
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("PutIpv6RouteRib", data)
        }
    }
    return err
}

func DeleteEndpointIpv6RouteRib(auth_token string, host string, instId string) error {
    logger := util.GetLoggerInstance()
    logger.Println("DeleteIpv6RouteRib")
    headers := GetReqHeader(auth_token)
    resp, err := doHttp("DELETE", "https://" + host + "/axapi/v3/ipv6/route/rib/" + instId, nil, headers)
    if err != nil {
        logger.Println("HTTP request failed with error", err)
    } else {
        logger.Println("HTTP status", resp.Status)
        data, _ := ioutil.ReadAll(resp.Body)
        logger.Println("response:", data)
        var m EndpointIpv6RouteRib
        err = json.Unmarshal(data, &m)
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        } else {
            err = checkAxApiResponseStatus("DeleteIpv6RouteRib", data)
        }
    }
    return err
}
