

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ThreatIntelWebrootGlobalStats struct {
    
    Stats ThreatIntelWebrootGlobalStatsStats `json:"stats"`

}
type DataThreatIntelWebrootGlobalStats struct {
    DtThreatIntelWebrootGlobalStats ThreatIntelWebrootGlobalStats `json:"webroot-global"`
}


type ThreatIntelWebrootGlobalStatsStats struct {
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
    RtuLookup int `json:"rtu-lookup"`
    DatabaseLookup int `json:"database-lookup"`
    NonMaliciousIps int `json:"non-malicious-ips"`
}

func (p *ThreatIntelWebrootGlobalStats) GetId() string{
    return "1"
}

func (p *ThreatIntelWebrootGlobalStats) getPath() string{
    return "threat-intel/webroot-global/stats"
}

func (p *ThreatIntelWebrootGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataThreatIntelWebrootGlobalStats,error) {
logger.Println("ThreatIntelWebrootGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataThreatIntelWebrootGlobalStats
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
