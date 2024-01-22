

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6NatIcmpv6 struct {
	Inst struct {

    RespondToPing int `json:"respond-to-ping"`
    Uuid string `json:"uuid"`

	} `json:"icmpv6"`
}

func (p *Ipv6NatIcmpv6) GetId() string{
    return "1"
}

func (p *Ipv6NatIcmpv6) getPath() string{
    return "ipv6/nat/icmpv6"
}

func (p *Ipv6NatIcmpv6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NatIcmpv6::Post")
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

func (p *Ipv6NatIcmpv6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NatIcmpv6::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *Ipv6NatIcmpv6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NatIcmpv6::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *Ipv6NatIcmpv6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6NatIcmpv6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
