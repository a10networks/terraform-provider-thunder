

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VisibilityTopnStats struct {
    
    Stats VisibilityTopnStatsStats `json:"stats"`

}
type DataVisibilityTopnStats struct {
    DtVisibilityTopnStats VisibilityTopnStats `json:"topn"`
}


type VisibilityTopnStatsStats struct {
    HeapAllocSuccess int `json:"heap-alloc-success"`
    HeapAllocFailed int `json:"heap-alloc-failed"`
    HeapAllocOom int `json:"heap-alloc-oom"`
    ObjRegSuccess int `json:"obj-reg-success"`
    ObjRegFailed int `json:"obj-reg-failed"`
    ObjRegOom int `json:"obj-reg-oom"`
    HeapDeleted int `json:"heap-deleted"`
    ObjDeleted int `json:"obj-deleted"`
}

func (p *VisibilityTopnStats) GetId() string{
    return "1"
}

func (p *VisibilityTopnStats) getPath() string{
    return "visibility/topn/stats"
}

func (p *VisibilityTopnStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVisibilityTopnStats,error) {
logger.Println("VisibilityTopnStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVisibilityTopnStats
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
