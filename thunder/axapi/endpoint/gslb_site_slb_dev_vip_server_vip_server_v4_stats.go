

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteSlbDevVipServerVipServerV4Stats struct {
    
    Ipv4 string `json:"ipv4"`
    Stats GslbSiteSlbDevVipServerVipServerV4StatsStats `json:"stats"`
    SiteName string 
    DeviceName string 

}
type DataGslbSiteSlbDevVipServerVipServerV4Stats struct {
    DtGslbSiteSlbDevVipServerVipServerV4Stats GslbSiteSlbDevVipServerVipServerV4Stats `json:"vip-server-v4"`
}


type GslbSiteSlbDevVipServerVipServerV4StatsStats struct {
    Dev_vip_hits int `json:"dev_vip_hits"`
    Dev_vip_recent int `json:"dev_vip_recent"`
}

func (p *GslbSiteSlbDevVipServerVipServerV4Stats) GetId() string{
    return "1"
}

func (p *GslbSiteSlbDevVipServerVipServerV4Stats) getPath() string{
    
    return "gslb/site/" +p.SiteName + "/slb-dev/" +p.DeviceName + "/vip-server/vip-server-v4/"+p.Ipv4+"/stats"
}

func (p *GslbSiteSlbDevVipServerVipServerV4Stats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbSiteSlbDevVipServerVipServerV4Stats,error) {
logger.Println("GslbSiteSlbDevVipServerVipServerV4Stats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbSiteSlbDevVipServerVipServerV4Stats
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
