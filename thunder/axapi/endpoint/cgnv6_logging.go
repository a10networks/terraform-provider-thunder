

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Logging struct {
	Inst struct {

    NatQuotaExceeded Cgnv6LoggingNatQuotaExceeded84 `json:"nat-quota-exceeded"`
    NatResourceExhausted Cgnv6LoggingNatResourceExhausted85 `json:"nat-resource-exhausted"`
    SamplingEnable []Cgnv6LoggingSamplingEnable `json:"sampling-enable"`
    SourceAddress Cgnv6LoggingSourceAddress86 `json:"source-address"`
    TcpSvrStatus Cgnv6LoggingTcpSvrStatus87 `json:"tcp-svr-status"`
    Uuid string `json:"uuid"`

	} `json:"logging"`
}


type Cgnv6LoggingNatQuotaExceeded84 struct {
    Level string `json:"level" dval:"warning"`
    Uuid string `json:"uuid"`
}


type Cgnv6LoggingNatResourceExhausted85 struct {
    Level string `json:"level" dval:"critical"`
    Uuid string `json:"uuid"`
}


type Cgnv6LoggingSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type Cgnv6LoggingSourceAddress86 struct {
    Uuid string `json:"uuid"`
}


type Cgnv6LoggingTcpSvrStatus87 struct {
    Uuid string `json:"uuid"`
}

func (p *Cgnv6Logging) GetId() string{
    return "1"
}

func (p *Cgnv6Logging) getPath() string{
    return "cgnv6/logging"
}

func (p *Cgnv6Logging) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Logging::Post")
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

func (p *Cgnv6Logging) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Logging::Get")
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
func (p *Cgnv6Logging) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Logging::Put")
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

func (p *Cgnv6Logging) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Logging::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
