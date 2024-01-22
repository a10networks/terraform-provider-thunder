

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpFrag struct {
	Inst struct {

    Buff int `json:"buff"`
    CpuThreshold IpFragCpuThreshold `json:"cpu-threshold"`
    MaxPacketsPerReassembly int `json:"max-packets-per-reassembly"`
    MaxReassemblySessions int `json:"max-reassembly-sessions"`
    SamplingEnable []IpFragSamplingEnable `json:"sampling-enable"`
    Timeout int `json:"timeout" dval:"60000"`
    Uuid string `json:"uuid"`

	} `json:"frag"`
}


type IpFragCpuThreshold struct {
    High int `json:"high" dval:"75"`
    Low int `json:"low" dval:"60"`
}


type IpFragSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *IpFrag) GetId() string{
    return "1"
}

func (p *IpFrag) getPath() string{
    return "ip/frag"
}

func (p *IpFrag) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpFrag::Post")
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

func (p *IpFrag) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpFrag::Get")
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
func (p *IpFrag) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpFrag::Put")
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

func (p *IpFrag) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpFrag::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
