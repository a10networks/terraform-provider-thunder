

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHealthDownReasonOper struct {
    
    Oper SlbHealthDownReasonOperOper `json:"oper"`

}
type DataSlbHealthDownReasonOper struct {
    DtSlbHealthDownReasonOper SlbHealthDownReasonOper `json:"health-down-reason"`
}


type SlbHealthDownReasonOperOper struct {
    Down_id int `json:"down_id"`
    Down_reason string `json:"down_reason"`
}

func (p *SlbHealthDownReasonOper) GetId() string{
    return "1"
}

func (p *SlbHealthDownReasonOper) getPath() string{
    return "slb/health-down-reason/oper"
}

func (p *SlbHealthDownReasonOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHealthDownReasonOper,error) {
logger.Println("SlbHealthDownReasonOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHealthDownReasonOper
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
