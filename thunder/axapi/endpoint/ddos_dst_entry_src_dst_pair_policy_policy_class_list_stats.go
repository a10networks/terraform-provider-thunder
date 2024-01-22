

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntrySrcDstPairPolicyPolicyClassListStats struct {
    
    ClassListName string `json:"class-list-name"`
    Stats DdosDstEntrySrcDstPairPolicyPolicyClassListStatsStats `json:"stats"`
    SrcBasedPolicyName string 
    DstEntryName string 

}
type DataDdosDstEntrySrcDstPairPolicyPolicyClassListStats struct {
    DtDdosDstEntrySrcDstPairPolicyPolicyClassListStats DdosDstEntrySrcDstPairPolicyPolicyClassListStats `json:"policy-class-list"`
}


type DdosDstEntrySrcDstPairPolicyPolicyClassListStatsStats struct {
    Packet_received int `json:"packet_received"`
    Packet_dropped int `json:"packet_dropped"`
    Entry_learned int `json:"entry_learned"`
    Entry_count_overflow int `json:"entry_count_overflow"`
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListStats) GetId() string{
    return "1"
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListStats) getPath() string{
    
    return "ddos/dst/entry/" +p.DstEntryName + "/src-dst-pair-policy/" +p.SrcBasedPolicyName + "/policy-class-list/"+p.ClassListName+"/stats"
}

func (p *DdosDstEntrySrcDstPairPolicyPolicyClassListStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataDdosDstEntrySrcDstPairPolicyPolicyClassListStats,error) {
logger.Println("DdosDstEntrySrcDstPairPolicyPolicyClassListStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataDdosDstEntrySrcDstPairPolicyPolicyClassListStats
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
