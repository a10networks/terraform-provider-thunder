

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbDnsPersistentCacheStats struct {
    
    Stats SlbDnsPersistentCacheStatsStats `json:"stats"`

}
type DataSlbDnsPersistentCacheStats struct {
    DtSlbDnsPersistentCacheStats SlbDnsPersistentCacheStats `json:"dns-persistent-cache"`
}


type SlbDnsPersistentCacheStatsStats struct {
    Total_entry int `json:"total_entry"`
    Entry_saved int `json:"entry_saved"`
    Entry_deleted int `json:"entry_deleted"`
    Database_busy int `json:"database_busy"`
    Database_error int `json:"database_error"`
}

func (p *SlbDnsPersistentCacheStats) GetId() string{
    return "1"
}

func (p *SlbDnsPersistentCacheStats) getPath() string{
    return "slb/dns-persistent-cache/stats"
}

func (p *SlbDnsPersistentCacheStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbDnsPersistentCacheStats,error) {
logger.Println("SlbDnsPersistentCacheStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbDnsPersistentCacheStats
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
