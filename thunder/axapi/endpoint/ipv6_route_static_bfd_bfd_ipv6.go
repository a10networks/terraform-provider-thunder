

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6RouteStaticBfdBfdIpv6 struct {
	Inst struct {

    Action string `json:"action"`
    LocalIpv6 string `json:"local-ipv6"`
    NexthopIpv6 string `json:"nexthop-ipv6"`
    Template string `json:"template"`
    Threshold int `json:"threshold"`
    Uuid string `json:"uuid"`

	} `json:"bfd-ipv6"`
}

func (p *Ipv6RouteStaticBfdBfdIpv6) GetId() string{
    return p.Inst.LocalIpv6+"+"+p.Inst.NexthopIpv6
}

func (p *Ipv6RouteStaticBfdBfdIpv6) getPath() string{
    return "ipv6/route/static/bfd/bfd-ipv6"
}

func (p *Ipv6RouteStaticBfdBfdIpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteStaticBfdBfdIpv6::Post")
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

func (p *Ipv6RouteStaticBfdBfdIpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteStaticBfdBfdIpv6::Get")
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
func (p *Ipv6RouteStaticBfdBfdIpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteStaticBfdBfdIpv6::Put")
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

func (p *Ipv6RouteStaticBfdBfdIpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteStaticBfdBfdIpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
