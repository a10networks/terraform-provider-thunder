

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats struct {
    
    ClassListName string `json:"class-list-name"`
    Stats DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsStats `json:"stats"`
    ZoneName string 
    Protocol string 
    SrcBasedPolicyName string 
    PortNum string 

}
type DataDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats struct {
    DtDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats `json:"policy-class-list"`
}


type DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStatsStats struct {
    Packet_received int `json:"packet_received"`
    Packet_dropped int `json:"packet_dropped"`
    Entry_learned int `json:"entry_learned"`
    Entry_count_overflow int `json:"entry_count_overflow"`
}

func (p *DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats) GetId() string{
    return "1"
}

func (p *DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats) getPath() string{
    
    return "ddos/dst/zone/" +p.ZoneName + "/port/zone-service/" +p.PortNum + "+" +p.Protocol + "/src-based-policy/" +p.SrcBasedPolicyName + "/policy-class-list/"+p.ClassListName+"/stats"
}

func (p *DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats,error) {
logger.Println("DdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstZonePortZoneServiceSrcBasedPolicyPolicyClassListStats
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
