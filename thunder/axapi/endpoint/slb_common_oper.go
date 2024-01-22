

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbCommonOper struct {
    
    Oper SlbCommonOperOper `json:"oper"`

}
type DataSlbCommonOper struct {
    DtSlbCommonOper SlbCommonOper `json:"common"`
}


type SlbCommonOperOper struct {
    ServerAutoReselect int `json:"server-auto-reselect"`
}

func (p *SlbCommonOper) GetId() string{
    return "1"
}

func (p *SlbCommonOper) getPath() string{
    return "slb/common/oper"
}

func (p *SlbCommonOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSlbCommonOper,error) {
logger.Println("SlbCommonOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSlbCommonOper
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
