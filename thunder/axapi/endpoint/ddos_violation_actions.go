

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosViolationActions struct {
	Inst struct {

    Blackhole int `json:"blackhole"`
    BlacklistSrc int `json:"blacklist-src"`
    ExecuteScript string `json:"execute-script"`
    ExecuteScriptTimeout int `json:"execute-script-timeout"`
    Name string `json:"name"`
    Notification []DdosViolationActionsNotification `json:"notification"`
    SendNotificationOnly int `json:"send-notification-only"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"violation-actions"`
}


type DdosViolationActionsNotification struct {
    NotificationTemplate string `json:"notification-template"`
}

func (p *DdosViolationActions) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosViolationActions) getPath() string{
    return "ddos/violation-actions"
}

func (p *DdosViolationActions) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosViolationActions::Post")
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

func (p *DdosViolationActions) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosViolationActions::Get")
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
func (p *DdosViolationActions) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosViolationActions::Put")
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

func (p *DdosViolationActions) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosViolationActions::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
