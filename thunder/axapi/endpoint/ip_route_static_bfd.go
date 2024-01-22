

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpRouteStaticBfd struct {
	Inst struct {

    Action string `json:"action"`
    LocalIp string `json:"local-ip"`
    NexthopIp string `json:"nexthop-ip"`
    Template string `json:"template"`
    Threshold int `json:"threshold"`
    Uuid string `json:"uuid"`

	} `json:"bfd"`
}

func (p *IpRouteStaticBfd) GetId() string{
    return p.Inst.LocalIp+"+"+p.Inst.NexthopIp
}

func (p *IpRouteStaticBfd) getPath() string{
    return "ip/route/static/bfd"
}

func (p *IpRouteStaticBfd) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpRouteStaticBfd::Post")
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

func (p *IpRouteStaticBfd) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpRouteStaticBfd::Get")
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
func (p *IpRouteStaticBfd) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpRouteStaticBfd::Put")
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

func (p *IpRouteStaticBfd) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpRouteStaticBfd::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
