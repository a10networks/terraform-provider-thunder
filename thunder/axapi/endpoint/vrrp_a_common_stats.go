

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VrrpACommonStats struct {
    
    Stats VrrpACommonStatsStats `json:"stats"`

}
type DataVrrpACommonStats struct {
    DtVrrpACommonStats VrrpACommonStats `json:"common"`
}


type VrrpACommonStatsStats struct {
    Vrrp_common_dummy int `json:"vrrp_common_dummy"`
}

func (p *VrrpACommonStats) GetId() string{
    return "1"
}

func (p *VrrpACommonStats) getPath() string{
    return "vrrp-a/common/stats"
}

func (p *VrrpACommonStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVrrpACommonStats,error) {
logger.Println("VrrpACommonStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVrrpACommonStats
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
