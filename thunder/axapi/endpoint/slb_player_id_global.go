

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbPlayerIdGlobal struct {
	Inst struct {

    AbsMaxExpiration int `json:"abs-max-expiration" dval:"10"`
    Enable64bitPlayerId int `json:"enable-64bit-player-id"`
    EnforcementTimer int `json:"enforcement-timer" dval:"480"`
    ForcePassive int `json:"force-passive"`
    MinExpiration int `json:"min-expiration" dval:"1"`
    PktActivityExpiration int `json:"pkt-activity-expiration" dval:"5"`
    SamplingEnable []SlbPlayerIdGlobalSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"player-id-global"`
}


type SlbPlayerIdGlobalSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *SlbPlayerIdGlobal) GetId() string{
    return "1"
}

func (p *SlbPlayerIdGlobal) getPath() string{
    return "slb/player-id-global"
}

func (p *SlbPlayerIdGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbPlayerIdGlobal::Post")
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

func (p *SlbPlayerIdGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbPlayerIdGlobal::Get")
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
func (p *SlbPlayerIdGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbPlayerIdGlobal::Put")
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

func (p *SlbPlayerIdGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbPlayerIdGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
