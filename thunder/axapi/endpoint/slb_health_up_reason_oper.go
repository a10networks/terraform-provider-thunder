

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbHealthUpReasonOper struct {
    
    Oper SlbHealthUpReasonOperOper `json:"oper"`

}
type DataSlbHealthUpReasonOper struct {
    DtSlbHealthUpReasonOper SlbHealthUpReasonOper `json:"health-up-reason"`
}


type SlbHealthUpReasonOperOper struct {
    Up_id int `json:"up_id"`
    Up_reason string `json:"up_reason"`
}

func (p *SlbHealthUpReasonOper) GetId() string{
    return "1"
}

func (p *SlbHealthUpReasonOper) getPath() string{
    return "slb/health-up-reason/oper"
}

func (p *SlbHealthUpReasonOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbHealthUpReasonOper,error) {
logger.Println("SlbHealthUpReasonOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbHealthUpReasonOper
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
