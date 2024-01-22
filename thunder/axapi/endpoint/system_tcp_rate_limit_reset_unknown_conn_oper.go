

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTcpRateLimitResetUnknownConnOper struct {
    
    Oper SystemTcpRateLimitResetUnknownConnOperOper `json:"oper"`

}
type DataSystemTcpRateLimitResetUnknownConnOper struct {
    DtSystemTcpRateLimitResetUnknownConnOper SystemTcpRateLimitResetUnknownConnOper `json:"rate-limit-reset-unknown-conn"`
}


type SystemTcpRateLimitResetUnknownConnOperOper struct {
    UnknownConnRateLimit int `json:"unknown-conn-rate-limit"`
    UnknownConnCurrentRate int `json:"unknown-conn-current-rate"`
    UnknownConnRateLimitDrop int `json:"unknown-conn-rate-limit-drop"`
}

func (p *SystemTcpRateLimitResetUnknownConnOper) GetId() string{
    return "1"
}

func (p *SystemTcpRateLimitResetUnknownConnOper) getPath() string{
    return "system/tcp/rate-limit-reset-unknown-conn/oper"
}

func (p *SystemTcpRateLimitResetUnknownConnOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemTcpRateLimitResetUnknownConnOper,error) {
logger.Println("SystemTcpRateLimitResetUnknownConnOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemTcpRateLimitResetUnknownConnOper
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
