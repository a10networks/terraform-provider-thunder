

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwDdosProtection struct {
	Inst struct {

    Action FwDdosProtectionAction `json:"action"`
    DynamicBlacklist FwDdosProtectionDynamicBlacklist `json:"dynamic-blacklist"`
    Logging FwDdosProtectionLogging `json:"logging"`
    SamplingEnable []FwDdosProtectionSamplingEnable `json:"sampling-enable"`
    Uuid string `json:"uuid"`

	} `json:"ddos-protection"`
}


type FwDdosProtectionAction struct {
    ActionType string `json:"action-type" dval:"drop"`
    RouteMap string `json:"route-map"`
    Expiration int `json:"expiration" dval:"5"`
    ExpirationRoute int `json:"expiration-route" dval:"60"`
    TimerMultiplyMax int `json:"timer-multiply-max" dval:"6"`
    RemoveWaitTimer int `json:"remove-wait-timer" dval:"300"`
}


type FwDdosProtectionDynamicBlacklist struct {
    DynamicBlacklistAction string `json:"dynamic-blacklist-action" dval:"disable"`
    Dir string `json:"dir" dval:"both"`
    Timeout int `json:"timeout" dval:"5"`
    CpuThreshold int `json:"cpu-threshold" dval:"60"`
}


type FwDdosProtectionLogging struct {
    LoggingAction string `json:"logging-action" dval:"enable"`
    EnableAction string `json:"enable-action" dval:"local"`
}


type FwDdosProtectionSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *FwDdosProtection) GetId() string{
    return "1"
}

func (p *FwDdosProtection) getPath() string{
    return "fw/ddos-protection"
}

func (p *FwDdosProtection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwDdosProtection::Post")
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

func (p *FwDdosProtection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwDdosProtection::Get")
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
func (p *FwDdosProtection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwDdosProtection::Put")
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

func (p *FwDdosProtection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwDdosProtection::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
