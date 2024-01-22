

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemCounterLibAccountingStats struct {
    
    Stats SystemCounterLibAccountingStatsStats `json:"stats"`

}
type DataSystemCounterLibAccountingStats struct {
    DtSystemCounterLibAccountingStats SystemCounterLibAccountingStats `json:"counter-lib-accounting"`
}


type SystemCounterLibAccountingStatsStats struct {
    TotalCtrAlloc int `json:"total-ctr-alloc"`
    TotalCtrInRml int `json:"total-ctr-in-rml"`
    TotalCtrFreed int `json:"total-ctr-freed"`
    TotalCtrInSystem int `json:"total-ctr-in-system"`
    TotalOperAlloc int `json:"total-oper-alloc"`
    TotalOperFree int `json:"total-oper-free"`
    TotalBlocksInHash int `json:"total-blocks-in-hash"`
    TotalBlocksInRmlHash int `json:"total-blocks-in-rml-hash"`
    TotalNodesAlloc int `json:"total-nodes-alloc"`
    TotalNodesInRml int `json:"total-nodes-in-rml"`
    TotalNodesFree int `json:"total-nodes-free"`
    TotalNodesFreeFailed int `json:"total-nodes-free-failed"`
    TotalNodesUnlinkFailed int `json:"total-nodes-unlink-failed"`
    TotalNodesInSystem int `json:"total-nodes-in-system"`
    TotalMemblocksAllocAvro int `json:"total-memblocks-alloc-avro"`
    TotalMemblocksFreeAvro int `json:"total-memblocks-free-avro"`
    TotalMemblocksReallocAvro int `json:"total-memblocks-realloc-avro"`
}

func (p *SystemCounterLibAccountingStats) GetId() string{
    return "1"
}

func (p *SystemCounterLibAccountingStats) getPath() string{
    return "system/counter-lib-accounting/stats"
}

func (p *SystemCounterLibAccountingStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemCounterLibAccountingStats,error) {
logger.Println("SystemCounterLibAccountingStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemCounterLibAccountingStats
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
