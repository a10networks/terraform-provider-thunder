

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTcp struct {
	Inst struct {

    RateLimitResetUnknownConn SystemTcpRateLimitResetUnknownConn1646 `json:"rate-limit-reset-unknown-conn"`
    SamplingEnable []SystemTcpSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"tcp"`
}


type SystemTcpRateLimitResetUnknownConn1646 struct {
    PktRateForResetUnknownConn int `json:"pkt-rate-for-reset-unknown-conn"`
    LogForResetUnknownConn int `json:"log-for-reset-unknown-conn"`
    Uuid string `json:"uuid"`
}


type SystemTcpSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SystemTcp) GetId() string{
    return "1"
}

func (p *SystemTcp) getPath() string{
    return "system/tcp"
}

func (p *SystemTcp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTcp::Post")
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

func (p *SystemTcp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTcp::Get")
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
func (p *SystemTcp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTcp::Put")
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

func (p *SystemTcp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTcp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
