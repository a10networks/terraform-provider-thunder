

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbResetUnknownConnOper struct {
    
    Oper SlbResetUnknownConnOperOper `json:"oper"`

}
type DataSlbResetUnknownConnOper struct {
    DtSlbResetUnknownConnOper SlbResetUnknownConnOper `json:"reset-unknown-conn"`
}


type SlbResetUnknownConnOperOper struct {
    UnknownConnRateLimit int `json:"unknown-conn-rate-limit"`
    UnknownConnCurrentRate int `json:"unknown-conn-current-rate"`
    UnknownConnRateLimitDrop int `json:"unknown-conn-rate-limit-drop"`
}

func (p *SlbResetUnknownConnOper) GetId() string{
    return "1"
}

func (p *SlbResetUnknownConnOper) getPath() string{
    return "slb/reset-unknown-conn/oper"
}

func (p *SlbResetUnknownConnOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbResetUnknownConnOper,error) {
logger.Println("SlbResetUnknownConnOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbResetUnknownConnOper
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
