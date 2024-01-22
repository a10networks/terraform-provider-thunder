

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnIkeSaClientsOper struct {
    
    Oper VpnIkeSaClientsOperOper `json:"oper"`

}
type DataVpnIkeSaClientsOper struct {
    DtVpnIkeSaClientsOper VpnIkeSaClientsOper `json:"ike-sa-clients"`
}


type VpnIkeSaClientsOperOper struct {
    Name string `json:"name"`
    IkeSaClientsLocalIp string `json:"ike-sa-clients-local-ip"`
    IkeSaClientsRemoteGw []VpnIkeSaClientsOperOperIkeSaClientsRemoteGw `json:"ike-sa-clients-remote-gw"`
}


type VpnIkeSaClientsOperOperIkeSaClientsRemoteGw struct {
    IkeSaClientsRemoteGwIp string `json:"ike-sa-clients-remote-gw-ip"`
    IkeSaClientsRemoteGwRemoteId string `json:"ike-sa-clients-remote-gw-remote-id"`
    IkeSaClientsRemoteGwUserId string `json:"ike-sa-clients-remote-gw-user-id"`
    IkeSaClientsRemoteGwIdleTime string `json:"ike-sa-clients-remote-gw-idle-time"`
    IkeSaClientsRemoteGwSessionTime string `json:"ike-sa-clients-remote-gw-session-time"`
    IkeSaClientsRemoteGwBytes string `json:"ike-sa-clients-remote-gw-bytes"`
}

func (p *VpnIkeSaClientsOper) GetId() string{
    return "1"
}

func (p *VpnIkeSaClientsOper) getPath() string{
    return "vpn/ike-sa-clients/oper"
}

func (p *VpnIkeSaClientsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnIkeSaClientsOper,error) {
logger.Println("VpnIkeSaClientsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnIkeSaClientsOper
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
