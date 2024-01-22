

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemGeolocListStats struct {
    
    Name string `json:"name"`
    Stats SystemGeolocListStatsStats `json:"stats"`

}
type DataSystemGeolocListStats struct {
    DtSystemGeolocListStats SystemGeolocListStats `json:"geoloc-list"`
}


type SystemGeolocListStatsStats struct {
    HitCount int `json:"hit-count"`
    TotalGeoloc int `json:"total-geoloc"`
    TotalActive int `json:"total-active"`
}

func (p *SystemGeolocListStats) GetId() string{
    return "1"
}

func (p *SystemGeolocListStats) getPath() string{
    
    return "system/geoloc-list/"+p.Name+"/stats"
}

func (p *SystemGeolocListStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemGeolocListStats,error) {
logger.Println("SystemGeolocListStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemGeolocListStats
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
