

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats struct {
    
    ClassListName string `json:"class-list-name"`
    Stats DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsStats `json:"stats"`
    SrcBasedPolicyName string 
    ZoneName string 
    ProtocolNum string 

}
type DataDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats struct {
    DtDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats `json:"policy-class-list"`
}


type DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStatsStats struct {
    Packet_received int `json:"packet_received"`
    Packet_dropped int `json:"packet_dropped"`
    Entry_learned int `json:"entry_learned"`
    Entry_count_overflow int `json:"entry_count_overflow"`
}

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-number/" +p.ProtocolNum + "/src-based-policy/" +p.SrcBasedPolicyName + "/policy-class-list/"+p.ClassListName+"/stats"
}

func (p *DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats,error) {
logger.Println("DdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNumberSrcBasedPolicyPolicyClassListStats
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
