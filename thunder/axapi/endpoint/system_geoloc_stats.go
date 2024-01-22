

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGeolocStats struct {
    
    Stats SystemGeolocStatsStats `json:"stats"`

}
type DataSystemGeolocStats struct {
    DtSystemGeolocStats SystemGeolocStats `json:"geoloc"`
}


type SystemGeolocStatsStats struct {
}

func (p *SystemGeolocStats) GetId() string{
    return "1"
}

func (p *SystemGeolocStats) getPath() string{
    return "system/geoloc/stats"
}

func (p *SystemGeolocStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemGeolocStats,error) {
logger.Println("SystemGeolocStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemGeolocStats
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
