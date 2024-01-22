

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbSiteIpServerStats struct {
    
    IpServerName string `json:"ip-server-name"`
    Stats GslbSiteIpServerStatsStats `json:"stats"`
    SiteName string 

}
type DataGslbSiteIpServerStats struct {
    DtGslbSiteIpServerStats GslbSiteIpServerStats `json:"ip-server"`
}


type GslbSiteIpServerStatsStats struct {
    Hits int `json:"hits"`
    Recent int `json:"recent"`
}

func (p *GslbSiteIpServerStats) GetId() string{
    return "1"
}

func (p *GslbSiteIpServerStats) getPath() string{
    
    return "gslb/site/" +p.SiteName + "/ip-server/"+p.IpServerName+"/stats"
}

func (p *GslbSiteIpServerStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbSiteIpServerStats,error) {
logger.Println("GslbSiteIpServerStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbSiteIpServerStats
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
