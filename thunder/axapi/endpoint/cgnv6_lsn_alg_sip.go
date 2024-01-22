

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnAlgSip struct {
	Inst struct {

    RtpStunTimeout int `json:"rtp-stun-timeout" dval:"5"`
    SamplingEnable []Cgnv6LsnAlgSipSamplingEnable `json:"sampling-enable"`
    SipValue string `json:"sip-value"`
    Uuid string `json:"uuid"`

	} `json:"sip"`
}


type Cgnv6LsnAlgSipSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Cgnv6LsnAlgSip) GetId() string{
    return "1"
}

func (p *Cgnv6LsnAlgSip) getPath() string{
    return "cgnv6/lsn/alg/sip"
}

func (p *Cgnv6LsnAlgSip) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnAlgSip::Post")
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

func (p *Cgnv6LsnAlgSip) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnAlgSip::Get")
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
func (p *Cgnv6LsnAlgSip) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnAlgSip::Put")
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

func (p *Cgnv6LsnAlgSip) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnAlgSip::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
