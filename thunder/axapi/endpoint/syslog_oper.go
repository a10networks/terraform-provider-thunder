

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SyslogOper struct {
    
    Oper SyslogOperOper `json:"oper"`

}
type DataSyslogOper struct {
    DtSyslogOper SyslogOper `json:"syslog"`
}


type SyslogOperOper struct {
    SystemLog []SyslogOperOperSystemLog `json:"system-log"`
    NextMsgIdx int `json:"next-msg-idx"`
}


type SyslogOperOperSystemLog struct {
    LogData string `json:"log-data"`
    AddSlotInfo int `json:"add-slot-info"`
    LogDataSearch string `json:"log-data-search"`
}

func (p *SyslogOper) GetId() string{
    return "1"
}

func (p *SyslogOper) getPath() string{
    return "syslog/oper"
}

func (p *SyslogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSyslogOper,error) {
logger.Println("SyslogOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSyslogOper
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
