

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type VrrpAVridStats struct {
    
    Stats VrrpAVridStatsStats `json:"stats"`
    VridVal int `json:"vrid-val"`

}
type DataVrrpAVridStats struct {
    DtVrrpAVridStats VrrpAVridStats `json:"vrid"`
}


type VrrpAVridStatsStats struct {
    Associated_vip_count int `json:"associated_vip_count"`
    Associated_vport_count int `json:"associated_vport_count"`
    Associated_natpool_count int `json:"associated_natpool_count"`
}

func (p *VrrpAVridStats) GetId() string{
    return "1"
}

func (p *VrrpAVridStats) getPath() string{
    
    return "vrrp-a/vrid/" +strconv.Itoa(p.VridVal)+"/stats"
}

func (p *VrrpAVridStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVrrpAVridStats,error) {
logger.Println("VrrpAVridStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVrrpAVridStats
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
