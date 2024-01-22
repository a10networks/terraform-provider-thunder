

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateTcpProgressionTracking struct {
	Inst struct {

    ConnectionTracking DdosTemplateTcpProgressionTrackingConnectionTracking300 `json:"connection-tracking"`
    FirstRequestMaxTime int `json:"first-request-max-time"`
    ProfilingConnectionLifeModel int `json:"profiling-connection-life-model"`
    ProfilingRequestResponseModel int `json:"profiling-request-response-model"`
    ProfilingTimeWindowModel int `json:"profiling-time-window-model"`
    ProgressionTrackingAction string `json:"progression-tracking-action" dval:"drop"`
    ProgressionTrackingActionListName string `json:"progression-tracking-action-list-name"`
    ProgressionTrackingEnabled string `json:"progression-tracking-enabled"`
    RequestLengthMax int `json:"request-length-max"`
    RequestLengthMin int `json:"request-length-min"`
    RequestResponseModel string `json:"request-response-model" dval:"enable"`
    RequestToResponseMaxTime int `json:"request-to-response-max-time"`
    ResponseLengthMax int `json:"response-length-max"`
    ResponseLengthMin int `json:"response-length-min"`
    ResponseRequestMaxRatio int `json:"response-request-max-ratio"`
    ResponseRequestMinRatio int `json:"response-request-min-ratio"`
    ResponseToRequestMaxTime int `json:"response-to-request-max-time"`
    TimeWindowTracking DdosTemplateTcpProgressionTrackingTimeWindowTracking301 `json:"time-window-tracking"`
    Uuid string `json:"uuid"`
    Violation int `json:"violation"`
    Name string 

	} `json:"progression-tracking"`
}


type DdosTemplateTcpProgressionTrackingConnectionTracking300 struct {
    ProgressionTrackingConnEnabled string `json:"progression-tracking-conn-enabled"`
    ConnSentMax int `json:"conn-sent-max"`
    ConnSentMin int `json:"conn-sent-min"`
    ConnRcvdMax int `json:"conn-rcvd-max"`
    ConnRcvdMin int `json:"conn-rcvd-min"`
    ConnRcvdSentRatioMin int `json:"conn-rcvd-sent-ratio-min"`
    ConnRcvdSentRatioMax int `json:"conn-rcvd-sent-ratio-max"`
    ConnDurationMax int `json:"conn-duration-max"`
    ConnDurationMin int `json:"conn-duration-min"`
    ConnViolation int `json:"conn-violation"`
    ProgressionTrackingConnActionListName string `json:"progression-tracking-conn-action-list-name"`
    ProgressionTrackingConnAction string `json:"progression-tracking-conn-action" dval:"drop"`
    Uuid string `json:"uuid"`
}


type DdosTemplateTcpProgressionTrackingTimeWindowTracking301 struct {
    ProgressionTrackingWinEnabled string `json:"progression-tracking-win-enabled"`
    WindowSentMax int `json:"window-sent-max"`
    WindowSentMin int `json:"window-sent-min"`
    WindowRcvdMax int `json:"window-rcvd-max"`
    WindowRcvdMin int `json:"window-rcvd-min"`
    WindowRcvdSentRatioMin int `json:"window-rcvd-sent-ratio-min"`
    WindowRcvdSentRatioMax int `json:"window-rcvd-sent-ratio-max"`
    WindowViolation int `json:"window-violation"`
    ProgressionTrackingWindowsActionListName string `json:"progression-tracking-windows-action-list-name"`
    ProgressionTrackingWindowsAction string `json:"progression-tracking-windows-action" dval:"drop"`
    Uuid string `json:"uuid"`
}

func (p *DdosTemplateTcpProgressionTracking) GetId() string{
    return "1"
}

func (p *DdosTemplateTcpProgressionTracking) getPath() string{
    return "ddos/template/tcp/" +p.Inst.Name + "/progression-tracking"
}

func (p *DdosTemplateTcpProgressionTracking) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcpProgressionTracking::Post")
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

func (p *DdosTemplateTcpProgressionTracking) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcpProgressionTracking::Get")
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
func (p *DdosTemplateTcpProgressionTracking) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcpProgressionTracking::Put")
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

func (p *DdosTemplateTcpProgressionTracking) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateTcpProgressionTracking::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
