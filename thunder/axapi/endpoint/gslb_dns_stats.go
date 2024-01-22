

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbDnsStats struct {
    
    Stats GslbDnsStatsStats `json:"stats"`

}
type DataGslbDnsStats struct {
    DtGslbDnsStats GslbDnsStats `json:"dns"`
}


type GslbDnsStatsStats struct {
    TotalQuery int `json:"total-query"`
    TotalResponse int `json:"total-response"`
    BadPacketQuery int `json:"bad-packet-query"`
    BadPacketResponse int `json:"bad-packet-response"`
    BadHeaderQuery int `json:"bad-header-query"`
    BadHeaderResponse int `json:"bad-header-response"`
    BadFormatQuery int `json:"bad-format-query"`
    BadFormatResponse int `json:"bad-format-response"`
    BadServiceQuery int `json:"bad-service-query"`
    BadServiceResponse int `json:"bad-service-response"`
    BadClassQuery int `json:"bad-class-query"`
    BadClassResponse int `json:"bad-class-response"`
    BadTypeQuery int `json:"bad-type-query"`
    BadTypeResponse int `json:"bad-type-response"`
    No_answer int `json:"no_answer"`
    Metric_health_check int `json:"metric_health_check"`
    Metric_weighted_ip int `json:"metric_weighted_ip"`
    Metric_weighted_site int `json:"metric_weighted_site"`
    Metric_capacity int `json:"metric_capacity"`
    Metric_active_server int `json:"metric_active_server"`
    Metric_easy_rdt int `json:"metric_easy_rdt"`
    Metric_active_rdt int `json:"metric_active_rdt"`
    Metric_geographic int `json:"metric_geographic"`
    Metric_connection_load int `json:"metric_connection_load"`
    Metric_number_of_sessions int `json:"metric_number_of_sessions"`
    Metric_active_weight int `json:"metric_active_weight"`
    Metric_admin_preference int `json:"metric_admin_preference"`
    Metric_bandwidth_quality int `json:"metric_bandwidth_quality"`
    Metric_bandwidth_cost int `json:"metric_bandwidth_cost"`
    Metric_user int `json:"metric_user"`
    Metric_least_reponse int `json:"metric_least_reponse"`
    Metric_admin_ip int `json:"metric_admin_ip"`
    Metric_round_robin int `json:"metric_round_robin"`
}

func (p *GslbDnsStats) GetId() string{
    return "1"
}

func (p *GslbDnsStats) getPath() string{
    return "gslb/dns/stats"
}

func (p *GslbDnsStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbDnsStats,error) {
logger.Println("GslbDnsStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbDnsStats
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
