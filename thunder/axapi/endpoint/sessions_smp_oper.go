

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SessionsSmpOper struct {
    
    Oper SessionsSmpOperOper `json:"oper"`

}
type DataSessionsSmpOper struct {
    DtSessionsSmpOper SessionsSmpOper `json:"smp"`
}


type SessionsSmpOperOper struct {
    SessionSmpList []SessionsSmpOperOperSessionSmpList `json:"session-smp-list"`
}


type SessionsSmpOperOperSessionSmpList struct {
    Type string `json:"type"`
    Alloc int `json:"alloc"`
    Free int `json:"free"`
    AllocFail int `json:"alloc-fail"`
}

func (p *SessionsSmpOper) GetId() string{
    return "1"
}

func (p *SessionsSmpOper) getPath() string{
    return "sessions/smp/oper"
}

func (p *SessionsSmpOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSessionsSmpOper,error) {
logger.Println("SessionsSmpOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSessionsSmpOper
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
