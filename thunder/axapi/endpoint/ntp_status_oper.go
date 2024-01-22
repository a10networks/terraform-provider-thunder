

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NtpStatusOper struct {
    
    Oper NtpStatusOperOper `json:"oper"`

}
type DataNtpStatusOper struct {
    DtNtpStatusOper NtpStatusOper `json:"ntp-status"`
}


type NtpStatusOperOper struct {
    Server []NtpStatusOperOperServer `json:"server"`
}


type NtpStatusOperOperServer struct {
    NtpServer string `json:"ntp-server"`
    Status string `json:"status"`
    Is_preferred string `json:"is_preferred"`
    Mode string `json:"mode"`
    Auth string `json:"auth"`
}

func (p *NtpStatusOper) GetId() string{
    return "1"
}

func (p *NtpStatusOper) getPath() string{
    return "ntp-status/oper"
}

func (p *NtpStatusOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataNtpStatusOper,error) {
logger.Println("NtpStatusOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataNtpStatusOper
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
