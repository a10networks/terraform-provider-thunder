

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SessionsExtOper struct {
    
    Oper SessionsExtOperOper `json:"oper"`

}
type DataSessionsExtOper struct {
    DtSessionsExtOper SessionsExtOper `json:"ext"`
}


type SessionsExtOperOper struct {
    SessionExtList []SessionsExtOperOperSessionExtList `json:"session-ext-list"`
}


type SessionsExtOperOperSessionExtList struct {
    Type string `json:"type"`
    Alloc int `json:"alloc"`
    Free int `json:"free"`
    Fail int `json:"fail"`
    CpuRoundRobinFail int `json:"cpu-round-robin-fail"`
    AllocExceed int `json:"alloc-exceed"`
}

func (p *SessionsExtOper) GetId() string{
    return "1"
}

func (p *SessionsExtOper) getPath() string{
    return "sessions/ext/oper"
}

func (p *SessionsExtOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSessionsExtOper,error) {
logger.Println("SessionsExtOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSessionsExtOper
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
