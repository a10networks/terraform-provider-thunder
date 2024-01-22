

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VcsShowdebugOper struct {
    
    Oper VcsShowdebugOperOper `json:"oper"`

}
type DataVcsShowdebugOper struct {
    DtVcsShowdebugOper VcsShowdebugOper `json:"showdebug"`
}


type VcsShowdebugOperOper struct {
    DebuggingSwitches string `json:"debugging-switches"`
    DebuggingBuffSize int `json:"debugging-buff-size"`
}

func (p *VcsShowdebugOper) GetId() string{
    return "1"
}

func (p *VcsShowdebugOper) getPath() string{
    return "vcs/showdebug/oper"
}

func (p *VcsShowdebugOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVcsShowdebugOper,error) {
logger.Println("VcsShowdebugOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVcsShowdebugOper
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
