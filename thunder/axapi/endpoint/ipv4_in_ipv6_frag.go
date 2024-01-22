

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv4InIpv6Frag struct {
	Inst struct {

    SamplingEnable []Ipv4InIpv6FragSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"frag"`
}


type Ipv4InIpv6FragSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Ipv4InIpv6Frag) GetId() string{
    return "1"
}

func (p *Ipv4InIpv6Frag) getPath() string{
    return "ipv4-in-ipv6/frag"
}

func (p *Ipv4InIpv6Frag) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv4InIpv6Frag::Post")
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

func (p *Ipv4InIpv6Frag) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv4InIpv6Frag::Get")
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
func (p *Ipv4InIpv6Frag) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv4InIpv6Frag::Put")
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

func (p *Ipv4InIpv6Frag) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv4InIpv6Frag::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
