

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type AcosEventsMessagesOper struct {
    
    Oper AcosEventsMessagesOperOper `json:"oper"`

}
type DataAcosEventsMessagesOper struct {
    DtAcosEventsMessagesOper AcosEventsMessagesOper `json:"messages"`
}


type AcosEventsMessagesOperOper struct {
    AllMsgs int `json:"all-msgs"`
    ActiveMsgs int `json:"active-msgs"`
    DebugMsgs int `json:"debug-msgs"`
    MessageId string `json:"message-id"`
    MessageIdScope string `json:"message-id-scope"`
    TotalLogCount int `json:"total-log-count"`
    ActiveLogCount int `json:"active-log-count"`
    MessagesList []AcosEventsMessagesOperOperMessagesList `json:"messages-list"`
}


type AcosEventsMessagesOperOperMessagesList struct {
    MessageId string `json:"message-id"`
    Syslog string `json:"syslog"`
    SyslogCollectorGrpsList []AcosEventsMessagesOperOperMessagesListSyslogCollectorGrpsList `json:"syslog-collector-grps-list"`
    Cef string `json:"cef"`
    CefCollectorGrpsList []AcosEventsMessagesOperOperMessagesListCefCollectorGrpsList `json:"cef-collector-grps-list"`
    Leef string `json:"leef"`
    LeefCollectorGrpsList []AcosEventsMessagesOperOperMessagesListLeefCollectorGrpsList `json:"leef-collector-grps-list"`
    MessageRoute string `json:"message-route"`
    Enabled int `json:"enabled"`
    Disabled int `json:"disabled"`
    NotDefined int `json:"not-defined"`
    Local string `json:"local"`
    Remote string `json:"remote"`
    CollectorGroup string `json:"collector-group"`
    DescriptiveName string `json:"descriptive-name"`
    Signature string `json:"signature"`
    TriggerReason string `json:"trigger-reason"`
    DefinedSeverity int `json:"defined-severity"`
    ConfiguredSeverity int `json:"configured-severity"`
    CefImportance int `json:"cef-importance"`
    RateLimiting string `json:"rate-limiting"`
    DefaultRoute string `json:"default-route"`
    ConfiguredRoute string `json:"configured-route"`
    LogDetailList []AcosEventsMessagesOperOperMessagesListLogDetailList `json:"log-detail-list"`
}


type AcosEventsMessagesOperOperMessagesListSyslogCollectorGrpsList struct {
    Name string `json:"name"`
}


type AcosEventsMessagesOperOperMessagesListCefCollectorGrpsList struct {
    Name string `json:"name"`
}


type AcosEventsMessagesOperOperMessagesListLeefCollectorGrpsList struct {
    Name string `json:"name"`
}


type AcosEventsMessagesOperOperMessagesListLogDetailList struct {
    LogName string `json:"log-name"`
    LocalLoggingConf string `json:"local-logging-conf"`
    RemoteLoggingConf string `json:"remote-logging-conf"`
    CollectorGroupsList []AcosEventsMessagesOperOperMessagesListLogDetailListCollectorGroupsList `json:"collector-groups-list"`
    LogFormat string `json:"log-format"`
}


type AcosEventsMessagesOperOperMessagesListLogDetailListCollectorGroupsList struct {
    Name string `json:"name"`
}

func (p *AcosEventsMessagesOper) GetId() string{
    return "1"
}

func (p *AcosEventsMessagesOper) getPath() string{
    return "acos-events/messages/oper"
}

func (p *AcosEventsMessagesOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataAcosEventsMessagesOper,error) {
logger.Println("AcosEventsMessagesOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataAcosEventsMessagesOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
