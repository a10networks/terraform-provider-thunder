

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpsGlobalStats struct {
    
    Stats IpsGlobalStatsStats `json:"stats"`

}
type DataIpsGlobalStats struct {
    DtIpsGlobalStats IpsGlobalStats `json:"global"`
}


type IpsGlobalStatsStats struct {
    Ips_matched_total int `json:"ips_matched_total"`
    Ips_matched_http int `json:"ips_matched_http"`
    Ips_matched_dns int `json:"ips_matched_dns"`
    Ips_matched_other int `json:"ips_matched_other"`
    Ips_matched_action_pass int `json:"ips_matched_action_pass"`
    Ips_matched_action_drop int `json:"ips_matched_action_drop"`
    Ips_matched_action_blacklist int `json:"ips_matched_action_blacklist"`
    Ips_matched_severity_high int `json:"ips_matched_severity_high"`
    Ips_matched_severity_medium int `json:"ips_matched_severity_medium"`
    Ips_matched_severity_low int `json:"ips_matched_severity_low"`
}

func (p *IpsGlobalStats) GetId() string{
    return "1"
}

func (p *IpsGlobalStats) getPath() string{
    return "ips/global/stats"
}

func (p *IpsGlobalStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataIpsGlobalStats,error) {
logger.Println("IpsGlobalStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataIpsGlobalStats
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
