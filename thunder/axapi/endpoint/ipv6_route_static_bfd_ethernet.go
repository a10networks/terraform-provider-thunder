

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Ipv6RouteStaticBfdEthernet struct {
	Inst struct {

    Action string `json:"action"`
    EthNum int `json:"eth-num"`
    NexthopIpv6Ll string `json:"nexthop-ipv6-ll"`
    Template string `json:"template"`
    Threshold int `json:"threshold"`
    Uuid string `json:"uuid"`

	} `json:"ethernet"`
}

func (p *Ipv6RouteStaticBfdEthernet) GetId() string{
    return strconv.Itoa(p.Inst.EthNum)+"+"+p.Inst.NexthopIpv6Ll
}

func (p *Ipv6RouteStaticBfdEthernet) getPath() string{
    return "ipv6/route/static/bfd/ethernet"
}

func (p *Ipv6RouteStaticBfdEthernet) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteStaticBfdEthernet::Post")
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

func (p *Ipv6RouteStaticBfdEthernet) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteStaticBfdEthernet::Get")
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
func (p *Ipv6RouteStaticBfdEthernet) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteStaticBfdEthernet::Put")
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

func (p *Ipv6RouteStaticBfdEthernet) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteStaticBfdEthernet::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
