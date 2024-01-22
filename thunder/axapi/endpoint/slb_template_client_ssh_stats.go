

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateClientSshStats struct {
    
    Name string `json:"name"`
    Stats SlbTemplateClientSshStatsStats `json:"stats"`

}
type DataSlbTemplateClientSshStats struct {
    DtSlbTemplateClientSshStats SlbTemplateClientSshStats `json:"client-ssh"`
}


type SlbTemplateClientSshStatsStats struct {
    Successful_handshakes int `json:"successful_handshakes"`
    Failed_handshakes int `json:"failed_handshakes"`
    Forwarding_errors int `json:"forwarding_errors"`
}

func (p *SlbTemplateClientSshStats) GetId() string{
    return "1"
}

func (p *SlbTemplateClientSshStats) getPath() string{
    
    return "slb/template/client-ssh/"+p.Name+"/stats"
}

func (p *SlbTemplateClientSshStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplateClientSshStats,error) {
logger.Println("SlbTemplateClientSshStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplateClientSshStats
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
