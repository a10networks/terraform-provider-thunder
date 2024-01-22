

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ThreatIntelWebrootLogOper struct {
    
    Oper ThreatIntelWebrootLogOperOper `json:"oper"`

}
type DataThreatIntelWebrootLogOper struct {
    DtThreatIntelWebrootLogOper ThreatIntelWebrootLogOper `json:"webroot-log"`
}


type ThreatIntelWebrootLogOperOper struct {
    WebrootLogList []ThreatIntelWebrootLogOperOperWebrootLogList `json:"webroot-log-list"`
    WebrootLogOffset int `json:"webroot-log-offset"`
    WebrootLogOver int `json:"webroot-log-over"`
    Follow int `json:"follow"`
    FromStart int `json:"from-start"`
    NumLines int `json:"num-lines"`
}


type ThreatIntelWebrootLogOperOperWebrootLogList struct {
    WebrootLogData string `json:"webroot-log-data"`
}

func (p *ThreatIntelWebrootLogOper) GetId() string{
    return "1"
}

func (p *ThreatIntelWebrootLogOper) getPath() string{
    return "threat-intel/webroot-log/oper"
}

func (p *ThreatIntelWebrootLogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataThreatIntelWebrootLogOper,error) {
logger.Println("ThreatIntelWebrootLogOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataThreatIntelWebrootLogOper
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
