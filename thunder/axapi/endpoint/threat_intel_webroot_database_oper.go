

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ThreatIntelWebrootDatabaseOper struct {
    
    Oper ThreatIntelWebrootDatabaseOperOper `json:"oper"`

}
type DataThreatIntelWebrootDatabaseOper struct {
    DtThreatIntelWebrootDatabaseOper ThreatIntelWebrootDatabaseOper `json:"webroot-database"`
}


type ThreatIntelWebrootDatabaseOperOper struct {
    Name string `json:"name"`
    Status string `json:"status"`
    Size string `json:"size"`
    Version int `json:"version"`
    LastUpdateTime string `json:"last-update-time"`
    NextUpdateTime string `json:"next-update-time"`
    ConnectionStatus string `json:"connection-status"`
    FailureReason string `json:"failure-reason"`
    LastSuccessfulConnection string `json:"last-successful-connection"`
    SpamSources int `json:"spam-sources"`
    WindowsExploits int `json:"windows-exploits"`
    WebAttacks int `json:"web-attacks"`
    Botnets int `json:"botnets"`
    Scanners int `json:"scanners"`
    DosAttacks int `json:"dos-attacks"`
    Reputation int `json:"reputation"`
    Phishing int `json:"phishing"`
    Proxy int `json:"proxy"`
    MobileThreats int `json:"mobile-threats"`
    TorProxy int `json:"tor-proxy"`
    TotalEntries int `json:"total-entries"`
}

func (p *ThreatIntelWebrootDatabaseOper) GetId() string{
    return "1"
}

func (p *ThreatIntelWebrootDatabaseOper) getPath() string{
    return "threat-intel/webroot-database/oper"
}

func (p *ThreatIntelWebrootDatabaseOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataThreatIntelWebrootDatabaseOper,error) {
logger.Println("ThreatIntelWebrootDatabaseOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataThreatIntelWebrootDatabaseOper
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
