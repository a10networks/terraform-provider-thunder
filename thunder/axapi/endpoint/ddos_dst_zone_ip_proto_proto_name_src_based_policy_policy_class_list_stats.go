

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListStats struct {
    
    ClassListName string `json:"class-list-name"`
    Stats DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListStatsStats `json:"stats"`
    SrcBasedPolicyName string 
    ZoneName string 
    Protocol string 

}
type DataDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListStats struct {
    DtDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListStats DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListStats `json:"policy-class-list"`
}


type DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListStatsStats struct {
    Packet_received int `json:"packet_received"`
    Packet_dropped int `json:"packet_dropped"`
    Entry_learned int `json:"entry_learned"`
    Entry_count_overflow int `json:"entry_count_overflow"`
}

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListStats) GetId() string{
    return "1"
}

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListStats) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/ip-proto/proto-name/" +p.Protocol + "/src-based-policy/" +p.SrcBasedPolicyName + "/policy-class-list/"+p.ClassListName+"/stats"
}

func (p *DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListStats,error) {
logger.Println("DdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZoneIpProtoProtoNameSrcBasedPolicyPolicyClassListStats
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
