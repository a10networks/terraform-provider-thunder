

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpNatAlgPptp struct {
	Inst struct {

    Pptp string `json:"pptp" dval:"disable"`
    SamplingEnable []IpNatAlgPptpSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"pptp"`
}


type IpNatAlgPptpSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *IpNatAlgPptp) GetId() string{
    return "1"
}

func (p *IpNatAlgPptp) getPath() string{
    return "ip/nat/alg/pptp"
}

func (p *IpNatAlgPptp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatAlgPptp::Post")
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

func (p *IpNatAlgPptp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatAlgPptp::Get")
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
func (p *IpNatAlgPptp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatAlgPptp::Put")
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

func (p *IpNatAlgPptp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpNatAlgPptp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
