

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteSlbDevVipServerVipServerV6Stats struct {
    
    Ipv6 string `json:"ipv6"`
    Stats GslbSiteSlbDevVipServerVipServerV6StatsStats `json:"stats"`
    SiteName string 
    DeviceName string 

}
type DataGslbSiteSlbDevVipServerVipServerV6Stats struct {
    DtGslbSiteSlbDevVipServerVipServerV6Stats GslbSiteSlbDevVipServerVipServerV6Stats `json:"vip-server-v6"`
}


type GslbSiteSlbDevVipServerVipServerV6StatsStats struct {
    Dev_vip_hits int `json:"dev_vip_hits"`
    Dev_vip_recent int `json:"dev_vip_recent"`
}

func (p *GslbSiteSlbDevVipServerVipServerV6Stats) GetId() string{
    return "1"
}

func (p *GslbSiteSlbDevVipServerVipServerV6Stats) getPath() string{
    
    return "gslb/site/" +p.SiteName + "/slb-dev/" +p.DeviceName + "/vip-server/vip-server-v6/"+p.Ipv6+"/stats"
}

func (p *GslbSiteSlbDevVipServerVipServerV6Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbSiteSlbDevVipServerVipServerV6Stats,error) {
logger.Println("GslbSiteSlbDevVipServerVipServerV6Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbSiteSlbDevVipServerVipServerV6Stats
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
