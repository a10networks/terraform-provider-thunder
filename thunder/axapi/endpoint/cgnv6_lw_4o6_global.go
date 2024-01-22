

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Lw4o6Global struct {
	Inst struct {

    Hairpinning string `json:"hairpinning" dval:"filter-none"`
    IcmpInbound string `json:"icmp-inbound" dval:"handle"`
    InsideSrcAccessList int `json:"inside-src-access-list"`
    NatPrefixList string `json:"nat-prefix-list"`
    NoForwardMatch Cgnv6Lw4o6GlobalNoForwardMatch `json:"no-forward-match"`
    NoReverseMatch Cgnv6Lw4o6GlobalNoReverseMatch `json:"no-reverse-match"`
    SamplingEnable []Cgnv6Lw4o6GlobalSamplingEnable `json:"sampling-enable"`
    UseBindingTable string `json:"use-binding-table"`
    Uuid string `json:"uuid"`

	} `json:"global"`
}


type Cgnv6Lw4o6GlobalNoForwardMatch struct {
    SendIcmpv6 int `json:"send-icmpv6"`
}


type Cgnv6Lw4o6GlobalNoReverseMatch struct {
    SendIcmp int `json:"send-icmp"`
}


type Cgnv6Lw4o6GlobalSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6Lw4o6Global) GetId() string{
    return "1"
}

func (p *Cgnv6Lw4o6Global) getPath() string{
    return "cgnv6/lw-4o6/global"
}

func (p *Cgnv6Lw4o6Global) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6Global::Post")
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

func (p *Cgnv6Lw4o6Global) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6Global::Get")
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
func (p *Cgnv6Lw4o6Global) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6Global::Put")
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

func (p *Cgnv6Lw4o6Global) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Lw4o6Global::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
