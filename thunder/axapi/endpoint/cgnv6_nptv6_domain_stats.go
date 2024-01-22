

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Nptv6DomainStats struct {
    
    Name string `json:"name"`
    Stats Cgnv6Nptv6DomainStatsStats `json:"stats"`

}
type DataCgnv6Nptv6DomainStats struct {
    DtCgnv6Nptv6DomainStats Cgnv6Nptv6DomainStats `json:"domain"`
}


type Cgnv6Nptv6DomainStatsStats struct {
    OutboundPackets int `json:"outbound-packets"`
    InboundPackets int `json:"inbound-packets"`
    HairpinPackets int `json:"hairpin-packets"`
    AddressNotValidForTranslation int `json:"address-not-valid-for-translation"`
    InboundPacketsNoMap int `json:"inbound-packets-no-map"`
    PacketsDestUnreachable int `json:"packets-dest-unreachable"`
}

func (p *Cgnv6Nptv6DomainStats) GetId() string{
    return "1"
}

func (p *Cgnv6Nptv6DomainStats) getPath() string{
    
    return "cgnv6/nptv6/domain/"+p.Name+"/stats"
}

func (p *Cgnv6Nptv6DomainStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6Nptv6DomainStats,error) {
logger.Println("Cgnv6Nptv6DomainStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6Nptv6DomainStats
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
