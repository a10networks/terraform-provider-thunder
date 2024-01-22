

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteSlbDevVipServerVipServerNameStats struct {
    
    Stats GslbSiteSlbDevVipServerVipServerNameStatsStats `json:"stats"`
    VipName string `json:"vip-name"`
    SiteName string 
    DeviceName string 

}
type DataGslbSiteSlbDevVipServerVipServerNameStats struct {
    DtGslbSiteSlbDevVipServerVipServerNameStats GslbSiteSlbDevVipServerVipServerNameStats `json:"vip-server-name"`
}


type GslbSiteSlbDevVipServerVipServerNameStatsStats struct {
    Dev_vip_hits int `json:"dev_vip_hits"`
    Dev_vip_recent int `json:"dev_vip_recent"`
}

func (p *GslbSiteSlbDevVipServerVipServerNameStats) GetId() string{
    return "1"
}

func (p *GslbSiteSlbDevVipServerVipServerNameStats) getPath() string{
    
    return "gslb/site/" +p.SiteName + "/slb-dev/" +p.DeviceName + "/vip-server/vip-server-name/"+p.VipName+"/stats"
}

func (p *GslbSiteSlbDevVipServerVipServerNameStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbSiteSlbDevVipServerVipServerNameStats,error) {
logger.Println("GslbSiteSlbDevVipServerVipServerNameStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbSiteSlbDevVipServerVipServerNameStats
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
