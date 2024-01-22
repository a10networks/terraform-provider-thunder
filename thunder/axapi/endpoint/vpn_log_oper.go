

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnLogOper struct {
    
    Oper VpnLogOperOper `json:"oper"`

}
type DataVpnLogOper struct {
    DtVpnLogOper VpnLogOper `json:"log"`
}


type VpnLogOperOper struct {
    VpnLogList []VpnLogOperOperVpnLogList `json:"vpn-log-list"`
    VpnLogOffset int `json:"vpn-log-offset"`
    VpnLogOver int `json:"vpn-log-over"`
    Follow int `json:"follow"`
    FromStart int `json:"from-start"`
    NumLines int `json:"num-lines"`
}


type VpnLogOperOperVpnLogList struct {
    VpnLogData string `json:"vpn-log-data"`
}

func (p *VpnLogOper) GetId() string{
    return "1"
}

func (p *VpnLogOper) getPath() string{
    return "vpn/log/oper"
}

func (p *VpnLogOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnLogOper,error) {
logger.Println("VpnLogOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnLogOper
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
