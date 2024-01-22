

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemRadiusServerStats struct {
    
    Stats SystemRadiusServerStatsStats `json:"stats"`

}
type DataSystemRadiusServerStats struct {
    DtSystemRadiusServerStats SystemRadiusServerStats `json:"server"`
}


type SystemRadiusServerStatsStats struct {
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
}

func (p *SystemRadiusServerStats) GetId() string{
    return "1"
}

func (p *SystemRadiusServerStats) getPath() string{
    return "system/radius/server/stats"
}

func (p *SystemRadiusServerStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemRadiusServerStats,error) {
logger.Println("SystemRadiusServerStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemRadiusServerStats
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
