

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateTcpProgressionTrackingConnectionTracking struct {
	Inst struct {

    ConnDurationMax int `json:"conn-duration-max"`
    ConnDurationMin int `json:"conn-duration-min"`
    ConnRcvdMax int `json:"conn-rcvd-max"`
    ConnRcvdMin int `json:"conn-rcvd-min"`
    ConnRcvdSentRatioMax int `json:"conn-rcvd-sent-ratio-max"`
    ConnRcvdSentRatioMin int `json:"conn-rcvd-sent-ratio-min"`
    ConnSentMax int `json:"conn-sent-max"`
    ConnSentMin int `json:"conn-sent-min"`
    ConnViolation int `json:"conn-violation"`
    ProgressionTrackingConnAction string `json:"progression-tracking-conn-action" dval:"drop"`
    ProgressionTrackingConnActionListName string `json:"progression-tracking-conn-action-list-name"`
    ProgressionTrackingConnEnabled string `json:"progression-tracking-conn-enabled"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"connection-tracking"`
}

func (p *DdosTemplateTcpProgressionTrackingConnectionTracking) GetId() string{
    return "1"
}

func (p *DdosTemplateTcpProgressionTrackingConnectionTracking) getPath() string{
    return "ddos/template/tcp/" +p.Inst.Name + "/progression-tracking/connection-tracking"
}

func (p *DdosTemplateTcpProgressionTrackingConnectionTracking) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcpProgressionTrackingConnectionTracking::Post")
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

func (p *DdosTemplateTcpProgressionTrackingConnectionTracking) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcpProgressionTrackingConnectionTracking::Get")
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
func (p *DdosTemplateTcpProgressionTrackingConnectionTracking) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcpProgressionTrackingConnectionTracking::Put")
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

func (p *DdosTemplateTcpProgressionTrackingConnectionTracking) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcpProgressionTrackingConnectionTracking::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
