

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbDnsCacheStats struct {
    
    Stats SlbDnsCacheStatsStats `json:"stats"`

}
type DataSlbDnsCacheStats struct {
    DtSlbDnsCacheStats SlbDnsCacheStats `json:"dns-cache"`
}


type SlbDnsCacheStatsStats struct {
    Total_q int `json:"total_q"`
    Total_r int `json:"total_r"`
    Hit int `json:"hit"`
    Bad_q int `json:"bad_q"`
    Encode_q int `json:"encode_q"`
    Multiple_q int `json:"multiple_q"`
    Oversize_q int `json:"oversize_q"`
    Bad_r int `json:"bad_r"`
    Oversize_r int `json:"oversize_r"`
    Encode_r int `json:"encode_r"`
    Multiple_r int `json:"multiple_r"`
    Answer_r int `json:"answer_r"`
    Ttl_r int `json:"ttl_r"`
    Ageout int `json:"ageout"`
    Bad_answer int `json:"bad_answer"`
    Ageout_weight int `json:"ageout_weight"`
    Total_log int `json:"total_log"`
    Total_alloc int `json:"total_alloc"`
    Total_freed int `json:"total_freed"`
    Current_allocate int `json:"current_allocate"`
    Current_data_allocate int `json:"current_data_allocate"`
    Resolver_queue_full int `json:"resolver_queue_full"`
    Truncated_r int `json:"truncated_r"`
}

func (p *SlbDnsCacheStats) GetId() string{
    return "1"
}

func (p *SlbDnsCacheStats) getPath() string{
    return "slb/dns-cache/stats"
}

func (p *SlbDnsCacheStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbDnsCacheStats,error) {
logger.Println("SlbDnsCacheStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbDnsCacheStats
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
