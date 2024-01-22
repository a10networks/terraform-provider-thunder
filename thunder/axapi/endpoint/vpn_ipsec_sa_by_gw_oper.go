

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnIpsecSaByGwOper struct {
    
    Oper VpnIpsecSaByGwOperOper `json:"oper"`

}
type DataVpnIpsecSaByGwOper struct {
    DtVpnIpsecSaByGwOper VpnIpsecSaByGwOper `json:"ipsec-sa-by-gw"`
}


type VpnIpsecSaByGwOperOper struct {
    IkeGatewayName string `json:"ike-gateway-name"`
    LocalIp string `json:"local-ip"`
    PeerIp string `json:"peer-ip"`
    IpsecSaList []VpnIpsecSaByGwOperOperIpsecSaList `json:"ipsec-sa-list"`
}


type VpnIpsecSaByGwOperOperIpsecSaList struct {
    IpsecSaName string `json:"ipsec-sa-name"`
    LocalTs string `json:"local-ts"`
    RemoteTs string `json:"remote-ts"`
    InSpi string `json:"in-spi"`
    OutSpi string `json:"out-spi"`
    Protocol string `json:"protocol"`
    Mode string `json:"mode"`
    Encryption string `json:"encryption"`
    Hash string `json:"hash"`
    Lifetime int `json:"lifetime"`
    Lifebytes string `json:"lifebytes"`
}

func (p *VpnIpsecSaByGwOper) GetId() string{
    return "1"
}

func (p *VpnIpsecSaByGwOper) getPath() string{
    return "vpn/ipsec-sa-by-gw/oper"
}

func (p *VpnIpsecSaByGwOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnIpsecSaByGwOper,error) {
logger.Println("VpnIpsecSaByGwOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnIpsecSaByGwOper
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
