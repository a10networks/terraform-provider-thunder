

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnIkeSaBriefOper struct {
    
    Oper VpnIkeSaBriefOperOper `json:"oper"`

}
type DataVpnIkeSaBriefOper struct {
    DtVpnIkeSaBriefOper VpnIkeSaBriefOper `json:"ike-sa-brief"`
}


type VpnIkeSaBriefOperOper struct {
    Name string `json:"name"`
    LocalIp string `json:"local-ip"`
    IkeSaBriefRemoteGw []VpnIkeSaBriefOperOperIkeSaBriefRemoteGw `json:"ike-sa-brief-remote-gw"`
}


type VpnIkeSaBriefOperOperIkeSaBriefRemoteGw struct {
    IkeSaBriefRemoteGwIp string `json:"ike-sa-brief-remote-gw-ip"`
    IkeSaBriefRemoteGwId string `json:"ike-sa-brief-remote-gw-id"`
    IkeSaBriefRemoteGwLifetime string `json:"ike-sa-brief-remote-gw-lifetime"`
    IkeSaBriefRemoteGwStatus string `json:"ike-sa-brief-remote-gw-status"`
}

func (p *VpnIkeSaBriefOper) GetId() string{
    return "1"
}

func (p *VpnIkeSaBriefOper) getPath() string{
    return "vpn/ike-sa-brief/oper"
}

func (p *VpnIkeSaBriefOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnIkeSaBriefOper,error) {
logger.Println("VpnIkeSaBriefOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnIkeSaBriefOper
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
