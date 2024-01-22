

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking struct {
	Inst struct {

    ProgressionTrackingWinEnabled string `json:"progression-tracking-win-enabled"`
    ProgressionTrackingWindowsAction string `json:"progression-tracking-windows-action" dval:"drop"`
    ProgressionTrackingWindowsActionListName string `json:"progression-tracking-windows-action-list-name"`
    Uuid string `json:"uuid"`
    WindowRcvdMax int `json:"window-rcvd-max"`
    WindowRcvdMin int `json:"window-rcvd-min"`
    WindowRcvdSentRatioMax int `json:"window-rcvd-sent-ratio-max"`
    WindowRcvdSentRatioMin int `json:"window-rcvd-sent-ratio-min"`
    WindowSentMax int `json:"window-sent-max"`
    WindowSentMin int `json:"window-sent-min"`
    WindowViolation int `json:"window-violation"`
    Name string 

	} `json:"time-window-tracking"`
}

func (p *DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking) GetId() string{
    return "1"
}

func (p *DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking) getPath() string{
    return "ddos/zone-template/tcp/" +p.Inst.Name + "/progression-tracking/time-window-tracking"
}

func (p *DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking::Post")
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

func (p *DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking::Get")
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
func (p *DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking::Put")
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

func (p *DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateTcpProgressionTrackingTimeWindowTracking::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
