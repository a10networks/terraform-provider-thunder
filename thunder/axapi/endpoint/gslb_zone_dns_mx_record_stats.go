

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbZoneDnsMxRecordStats struct {
    
    MxName string `json:"mx-name"`
    Stats GslbZoneDnsMxRecordStatsStats `json:"stats"`
    Name string 

}
type DataGslbZoneDnsMxRecordStats struct {
    DtGslbZoneDnsMxRecordStats GslbZoneDnsMxRecordStats `json:"dns-mx-record"`
}


type GslbZoneDnsMxRecordStatsStats struct {
    Hits int `json:"hits"`
}

func (p *GslbZoneDnsMxRecordStats) GetId() string{
    return "1"
}

func (p *GslbZoneDnsMxRecordStats) getPath() string{
    
    return "gslb/zone/" +p.Name + "/dns-mx-record/"+p.MxName+"/stats"
}

func (p *GslbZoneDnsMxRecordStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataGslbZoneDnsMxRecordStats,error) {
logger.Println("GslbZoneDnsMxRecordStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataGslbZoneDnsMxRecordStats
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
