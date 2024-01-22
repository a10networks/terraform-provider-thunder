

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGeolocNameHelperStats struct {
    
    Stats SystemGeolocNameHelperStatsStats `json:"stats"`

}
type DataSystemGeolocNameHelperStats struct {
    DtSystemGeolocNameHelperStats SystemGeolocNameHelperStats `json:"geoloc-name-helper"`
}


type SystemGeolocNameHelperStatsStats struct {
}

func (p *SystemGeolocNameHelperStats) GetId() string{
    return "1"
}

func (p *SystemGeolocNameHelperStats) getPath() string{
    return "system/geoloc-name-helper/stats"
}

func (p *SystemGeolocNameHelperStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemGeolocNameHelperStats,error) {
logger.Println("SystemGeolocNameHelperStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemGeolocNameHelperStats
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
