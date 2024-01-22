

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Ipv6RouteStaticBfdTrunk struct {
	Inst struct {

    Action string `json:"action"`
    NexthopIpv6Ll string `json:"nexthop-ipv6-ll"`
    Template string `json:"template"`
    Threshold int `json:"threshold"`
    TrunkNum int `json:"trunk-num"`
    Uuid string `json:"uuid"`

	} `json:"trunk"`
}

func (p *Ipv6RouteStaticBfdTrunk) GetId() string{
    return strconv.Itoa(p.Inst.TrunkNum)+"+"+p.Inst.NexthopIpv6Ll
}

func (p *Ipv6RouteStaticBfdTrunk) getPath() string{
    return "ipv6/route/static/bfd/trunk"
}

func (p *Ipv6RouteStaticBfdTrunk) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteStaticBfdTrunk::Post")
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

func (p *Ipv6RouteStaticBfdTrunk) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteStaticBfdTrunk::Get")
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
func (p *Ipv6RouteStaticBfdTrunk) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteStaticBfdTrunk::Put")
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

func (p *Ipv6RouteStaticBfdTrunk) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6RouteStaticBfdTrunk::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
