

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemSpeStatusOper struct {
    
    Oper SystemSpeStatusOperOper `json:"oper"`

}
type DataSystemSpeStatusOper struct {
    DtSystemSpeStatusOper SystemSpeStatusOper `json:"spe-status"`
}


type SystemSpeStatusOperOper struct {
    Enable string `json:"enable"`
    SpeProfile string `json:"spe-profile"`
    SpeSetupStatus string `json:"spe-setup-status"`
    Ipv4_allowed int `json:"ipv4_allowed"`
    Ipv6_allowed int `json:"ipv6_allowed"`
}

func (p *SystemSpeStatusOper) GetId() string{
    return "1"
}

func (p *SystemSpeStatusOper) getPath() string{
    return "system/spe-status/oper"
}

func (p *SystemSpeStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemSpeStatusOper,error) {
logger.Println("SystemSpeStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemSpeStatusOper
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
