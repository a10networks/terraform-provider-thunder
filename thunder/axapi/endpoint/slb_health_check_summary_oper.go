

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHealthCheckSummaryOper struct {
    
    Oper SlbHealthCheckSummaryOperOper `json:"oper"`

}
type DataSlbHealthCheckSummaryOper struct {
    DtSlbHealthCheckSummaryOper SlbHealthCheckSummaryOper `json:"health-check-summary"`
}


type SlbHealthCheckSummaryOperOper struct {
    ServerUp int `json:"server-up"`
    ServerDown int `json:"server-down"`
    ServerPortUp int `json:"server-port-up"`
    ServerPortDown int `json:"server-port-down"`
}

func (p *SlbHealthCheckSummaryOper) GetId() string{
    return "1"
}

func (p *SlbHealthCheckSummaryOper) getPath() string{
    return "slb/health-check-summary/oper"
}

func (p *SlbHealthCheckSummaryOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHealthCheckSummaryOper,error) {
logger.Println("SlbHealthCheckSummaryOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHealthCheckSummaryOper
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
