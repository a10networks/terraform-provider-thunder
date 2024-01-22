

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemSslStatusOper struct {
    
    Oper SystemSslStatusOperOper `json:"oper"`

}
type DataSystemSslStatusOper struct {
    DtSystemSslStatusOper SystemSslStatusOper `json:"ssl-status"`
}


type SystemSslStatusOperOper struct {
    Enable string `json:"enable"`
    SslSetupStatus string `json:"ssl-setup-status"`
    Num_chip int `json:"num_chip"`
    Num_aes int `json:"num_aes"`
    CryptoSupport string `json:"crypto-support"`
}

func (p *SystemSslStatusOper) GetId() string{
    return "1"
}

func (p *SystemSslStatusOper) getPath() string{
    return "system/ssl-status/oper"
}

func (p *SystemSslStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataSystemSslStatusOper,error) {
logger.Println("SystemSslStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataSystemSslStatusOper
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
