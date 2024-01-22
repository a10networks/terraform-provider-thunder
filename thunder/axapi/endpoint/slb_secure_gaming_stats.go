

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbSecureGamingStats struct {
    
    Stats SlbSecureGamingStatsStats `json:"stats"`

}
type DataSlbSecureGamingStats struct {
    DtSlbSecureGamingStats SlbSecureGamingStats `json:"secure-gaming"`
}


type SlbSecureGamingStatsStats struct {
    Secure_gaming_pass int `json:"secure_gaming_pass"`
    Secure_gaming_drop int `json:"secure_gaming_drop"`
}

func (p *SlbSecureGamingStats) GetId() string{
    return "1"
}

func (p *SlbSecureGamingStats) getPath() string{
    return "slb/secure-gaming/stats"
}

func (p *SlbSecureGamingStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbSecureGamingStats,error) {
logger.Println("SlbSecureGamingStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbSecureGamingStats
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
