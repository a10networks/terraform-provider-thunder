

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosIpProtoStats struct {
    
    Stats DdosIpProtoStatsStats `json:"stats"`

}
type DataDdosIpProtoStats struct {
    DtDdosIpProtoStats DdosIpProtoStats `json:"ip-proto"`
}


type DdosIpProtoStatsStats struct {
    Dst_ipproto_rcvd int `json:"dst_ipproto_rcvd"`
    Dst_ipproto_drop int `json:"dst_ipproto_drop"`
    Dst_ipproto_bl int `json:"dst_ipproto_bl"`
    Dst_ipproto_exceed_drop_any int `json:"dst_ipproto_exceed_drop_any"`
    Dst_ipproto_pkt_rate_exceed int `json:"dst_ipproto_pkt_rate_exceed"`
    Dst_ipproto_kbit_rate_exceed int `json:"dst_ipproto_kbit_rate_exceed"`
    Dst_ipproto_frag_rate_exceed int `json:"dst_ipproto_frag_rate_exceed"`
}

func (p *DdosIpProtoStats) GetId() string{
    return "1"
}

func (p *DdosIpProtoStats) getPath() string{
    return "ddos/ip-proto/stats"
}

func (p *DdosIpProtoStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosIpProtoStats,error) {
logger.Println("DdosIpProtoStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosIpProtoStats
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
