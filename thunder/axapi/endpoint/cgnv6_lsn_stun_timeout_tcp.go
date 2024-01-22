

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnStunTimeoutTcp struct {
	Inst struct {

    PortEnd int `json:"port-end"`
    PortStart int `json:"port-start"`
    Timeout int `json:"timeout"`
    Uuid string `json:"uuid"`

	} `json:"tcp"`
}

func (p *Cgnv6LsnStunTimeoutTcp) GetId() string{
    return strconv.Itoa(p.Inst.PortStart)+"+"+strconv.Itoa(p.Inst.PortEnd)
}

func (p *Cgnv6LsnStunTimeoutTcp) getPath() string{
    return "cgnv6/lsn/stun-timeout/tcp"
}

func (p *Cgnv6LsnStunTimeoutTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnStunTimeoutTcp::Post")
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

func (p *Cgnv6LsnStunTimeoutTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnStunTimeoutTcp::Get")
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
func (p *Cgnv6LsnStunTimeoutTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnStunTimeoutTcp::Put")
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

func (p *Cgnv6LsnStunTimeoutTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnStunTimeoutTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
