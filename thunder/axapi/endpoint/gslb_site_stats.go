

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteStats struct {
    
    IpServerList []GslbSiteStatsIpServerList `json:"ip-server-list"`
    SiteName string `json:"site-name"`
    Stats GslbSiteStatsStats `json:"stats"`

}
type DataGslbSiteStats struct {
    DtGslbSiteStats GslbSiteStats `json:"site"`
}


type GslbSiteStatsIpServerList struct {
    IpServerName string `json:"ip-server-name"`
    Stats GslbSiteStatsIpServerListStats `json:"stats"`
}


type GslbSiteStatsIpServerListStats struct {
    Hits int `json:"hits"`
    Recent int `json:"recent"`
}


type GslbSiteStatsStats struct {
    Hits int `json:"hits"`
}

func (p *GslbSiteStats) GetId() string{
    return "1"
}

func (p *GslbSiteStats) getPath() string{
    
    return "gslb/site/"+p.SiteName+"/stats"
}

func (p *GslbSiteStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbSiteStats,error) {
logger.Println("GslbSiteStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbSiteStats
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
