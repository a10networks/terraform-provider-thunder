

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6MapTranslationDomainStats struct {
    
    Name string `json:"name"`
    Stats Cgnv6MapTranslationDomainStatsStats `json:"stats"`

}
type DataCgnv6MapTranslationDomainStats struct {
    DtCgnv6MapTranslationDomainStats Cgnv6MapTranslationDomainStats `json:"domain"`
}


type Cgnv6MapTranslationDomainStatsStats struct {
    Inbound_packet_received int `json:"inbound_packet_received"`
    Inbound_frag_packet_received int `json:"inbound_frag_packet_received"`
    Inbound_addr_port_validation_failed int `json:"inbound_addr_port_validation_failed"`
    Inbound_rev_lookup_failed int `json:"inbound_rev_lookup_failed"`
    Inbound_dest_unreachable int `json:"inbound_dest_unreachable"`
    Outbound_packet_received int `json:"outbound_packet_received"`
    Outbound_frag_packet_received int `json:"outbound_frag_packet_received"`
    Outbound_addr_validation_failed int `json:"outbound_addr_validation_failed"`
    Outbound_rev_lookup_failed int `json:"outbound_rev_lookup_failed"`
    Outbound_dest_unreachable int `json:"outbound_dest_unreachable"`
    Packet_mtu_exceeded int `json:"packet_mtu_exceeded"`
    Frag_icmp_sent int `json:"frag_icmp_sent"`
    Interface_not_configured int `json:"interface_not_configured"`
    Bmr_prefixrules_configured int `json:"bmr_prefixrules_configured"`
}

func (p *Cgnv6MapTranslationDomainStats) GetId() string{
    return "1"
}

func (p *Cgnv6MapTranslationDomainStats) getPath() string{
    
    return "cgnv6/map/translation/domain/"+p.Name+"/stats"
}

func (p *Cgnv6MapTranslationDomainStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataCgnv6MapTranslationDomainStats,error) {
logger.Println("Cgnv6MapTranslationDomainStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataCgnv6MapTranslationDomainStats
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
