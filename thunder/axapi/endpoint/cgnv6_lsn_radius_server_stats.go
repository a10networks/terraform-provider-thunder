

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnRadiusServerStats struct {
    
    Stats Cgnv6LsnRadiusServerStatsStats `json:"stats"`

}
type DataCgnv6LsnRadiusServerStats struct {
    DtCgnv6LsnRadiusServerStats Cgnv6LsnRadiusServerStats `json:"server"`
}


type Cgnv6LsnRadiusServerStatsStats struct {
    MsisdnReceived int `json:"msisdn-received"`
    ImeiReceived int `json:"imei-received"`
    ImsiReceived int `json:"imsi-received"`
    CustomReceived int `json:"custom-received"`
    RadiusRequestReceived int `json:"radius-request-received"`
    RadiusRequestDropped int `json:"radius-request-dropped"`
    RequestBadSecretDropped int `json:"request-bad-secret-dropped"`
    RequestNoKeyVapDropped int `json:"request-no-key-vap-dropped"`
    RequestMalformedDropped int `json:"request-malformed-dropped"`
    RequestIgnored int `json:"request-ignored"`
    RadiusTableFull int `json:"radius-table-full"`
    SecretNotConfiguredDropped int `json:"secret-not-configured-dropped"`
    HaStandbyDropped int `json:"ha-standby-dropped"`
    Ipv6PrefixLengthMismatch int `json:"ipv6-prefix-length-mismatch"`
    InvalidKey int `json:"invalid-key"`
    SmpCreated int `json:"smp-created"`
    SmpDeleted int `json:"smp-deleted"`
}

func (p *Cgnv6LsnRadiusServerStats) GetId() string{
    return "1"
}

func (p *Cgnv6LsnRadiusServerStats) getPath() string{
    return "cgnv6/lsn/radius/server/stats"
}

func (p *Cgnv6LsnRadiusServerStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6LsnRadiusServerStats,error) {
logger.Println("Cgnv6LsnRadiusServerStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6LsnRadiusServerStats
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
