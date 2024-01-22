

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceIpsStats struct {
    
    Stats DdosDstZonePortZoneServiceIpsStatsStats `json:"stats"`
    ZoneName string 
    PortNum string 
    Protocol string 

}
type DataDdosDstZonePortZoneServiceIpsStats struct {
    DtDdosDstZonePortZoneServiceIpsStats DdosDstZonePortZoneServiceIpsStats `json:"ips"`
}


type DdosDstZonePortZoneServiceIpsStatsStats struct {
    Ips_matched_total int `json:"ips_matched_total"`
    Ips_matched_action_pass int `json:"ips_matched_action_pass"`
    Ips_matched_action_drop int `json:"ips_matched_action_drop"`
    Ips_matched_action_blacklist int `json:"ips_matched_action_blacklist"`
    Ips_matched_severity_high int `json:"ips_matched_severity_high"`
    Ips_matched_severity_medium int `json:"ips_matched_severity_medium"`
    Ips_matched_severity_low int `json:"ips_matched_severity_low"`
    Src_ips_matched_action_pass int `json:"src_ips_matched_action_pass"`
    Src_ips_matched_action_drop int `json:"src_ips_matched_action_drop"`
    Src_ips_matched_action_blacklist int `json:"src_ips_matched_action_blacklist"`
    Src_ips_matched_severity_high int `json:"src_ips_matched_severity_high"`
    Src_ips_matched_severity_medium int `json:"src_ips_matched_severity_medium"`
    Src_ips_matched_severity_low int `json:"src_ips_matched_severity_low"`
}

func (p *DdosDstZonePortZoneServiceIpsStats) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceIpsStats) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service/" +p.PortNum + "+" +p.Protocol + "/ips/stats"
}

func (p *DdosDstZonePortZoneServiceIpsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceIpsStats,error) {
logger.Println("DdosDstZonePortZoneServiceIpsStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceIpsStats
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
