

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateServerSshStats struct {
    
    Name string `json:"name"`
    Stats SlbTemplateServerSshStatsStats `json:"stats"`

}
type DataSlbTemplateServerSshStats struct {
    DtSlbTemplateServerSshStats SlbTemplateServerSshStats `json:"server-ssh"`
}


type SlbTemplateServerSshStatsStats struct {
    Successful_handshakes int `json:"successful_handshakes"`
    Failed_handshakes int `json:"failed_handshakes"`
    Forwarding_errors int `json:"forwarding_errors"`
}

func (p *SlbTemplateServerSshStats) GetId() string{
    return "1"
}

func (p *SlbTemplateServerSshStats) getPath() string{
    
    return "slb/template/server-ssh/"+p.Name+"/stats"
}

func (p *SlbTemplateServerSshStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbTemplateServerSshStats,error) {
logger.Println("SlbTemplateServerSshStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbTemplateServerSshStats
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
