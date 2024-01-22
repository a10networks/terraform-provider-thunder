

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6SixrdDomain struct {
	Inst struct {

    BrIpv4Address string `json:"br-ipv4-address"`
    CeIpv4Netmask string `json:"ce-ipv4-netmask"`
    CeIpv4Network string `json:"ce-ipv4-network"`
    Ipv6Prefix string `json:"ipv6-prefix"`
    Mtu int `json:"mtu"`
    Name string `json:"name"`
    SamplingEnable []Cgnv6SixrdDomainSamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"domain"`
}


type Cgnv6SixrdDomainSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6SixrdDomain) GetId() string{
    return p.Inst.Name
}

func (p *Cgnv6SixrdDomain) getPath() string{
    return "cgnv6/sixrd/domain"
}

func (p *Cgnv6SixrdDomain) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SixrdDomain::Post")
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

func (p *Cgnv6SixrdDomain) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SixrdDomain::Get")
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
func (p *Cgnv6SixrdDomain) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SixrdDomain::Put")
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

func (p *Cgnv6SixrdDomain) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6SixrdDomain::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
