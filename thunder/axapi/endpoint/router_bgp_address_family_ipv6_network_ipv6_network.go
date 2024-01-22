

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type RouterBgpAddressFamilyIpv6NetworkIpv6Network struct {
	Inst struct {

    Backdoor int `json:"backdoor"`
    CommValue string `json:"comm-value"`
    Description string `json:"description"`
    LcommValue string `json:"lcomm-value"`
    NetworkIpv6 string `json:"network-ipv6"`
    RouteMap string `json:"route-map"`
    Uuid string `json:"uuid"`
    AsNumber string 

	} `json:"ipv6-network"`
}

func (p *RouterBgpAddressFamilyIpv6NetworkIpv6Network) GetId() string{
    return url.QueryEscape(p.Inst.NetworkIpv6)
}

func (p *RouterBgpAddressFamilyIpv6NetworkIpv6Network) getPath() string{
    return "router/bgp/" +p.Inst.AsNumber + "/address-family/ipv6/network/ipv6-network"
}

func (p *RouterBgpAddressFamilyIpv6NetworkIpv6Network) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NetworkIpv6Network::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *RouterBgpAddressFamilyIpv6NetworkIpv6Network) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NetworkIpv6Network::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *RouterBgpAddressFamilyIpv6NetworkIpv6Network) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NetworkIpv6Network::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *RouterBgpAddressFamilyIpv6NetworkIpv6Network) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RouterBgpAddressFamilyIpv6NetworkIpv6Network::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
