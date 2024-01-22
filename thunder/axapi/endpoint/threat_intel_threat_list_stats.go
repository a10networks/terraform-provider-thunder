

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ThreatIntelThreatListStats struct {
    
    Name string `json:"name"`
    Stats ThreatIntelThreatListStatsStats `json:"stats"`

}
type DataThreatIntelThreatListStats struct {
    DtThreatIntelThreatListStats ThreatIntelThreatListStats `json:"threat-list"`
}


type ThreatIntelThreatListStatsStats struct {
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
    TotalHits int `json:"total-hits"`
}

func (p *ThreatIntelThreatListStats) GetId() string{
    return "1"
}

func (p *ThreatIntelThreatListStats) getPath() string{
    
    return "threat-intel/threat-list/"+p.Name+"/stats"
}

func (p *ThreatIntelThreatListStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataThreatIntelThreatListStats,error) {
logger.Println("ThreatIntelThreatListStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataThreatIntelThreatListStats
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
