

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbCommonConnRateLimitSrcIp struct {
	Inst struct {

    DisableIpv6Support int `json:"disable-ipv6-support"`
    ExceedAction int `json:"exceed-action"`
    Limit int `json:"limit"`
    LimitPeriod string `json:"limit-period"`
    LockOut int `json:"lock-out"`
    Log int `json:"log"`
    Protocol string `json:"protocol"`
    Shared int `json:"shared"`
    Uuid string `json:"uuid"`

	} `json:"src-ip"`
}

func (p *SlbCommonConnRateLimitSrcIp) GetId() string{
    return strconv.Itoa(p.Inst.DisableIpv6Support)+"+"+p.Inst.Protocol
}

func (p *SlbCommonConnRateLimitSrcIp) getPath() string{
    return "slb/common/conn-rate-limit/src-ip"
}

func (p *SlbCommonConnRateLimitSrcIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonConnRateLimitSrcIp::Post")
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

func (p *SlbCommonConnRateLimitSrcIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonConnRateLimitSrcIp::Get")
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
func (p *SlbCommonConnRateLimitSrcIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonConnRateLimitSrcIp::Put")
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

func (p *SlbCommonConnRateLimitSrcIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbCommonConnRateLimitSrcIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
