

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnIpsecSaOper struct {
    
    Oper VpnIpsecSaOperOper `json:"oper"`

}
type DataVpnIpsecSaOper struct {
    DtVpnIpsecSaOper VpnIpsecSaOper `json:"ipsec-sa"`
}


type VpnIpsecSaOperOper struct {
    IpsecSaList []VpnIpsecSaOperOperIpsecSaList `json:"ipsec-sa-list"`
}


type VpnIpsecSaOperOperIpsecSaList struct {
    IpsecSaName string `json:"ipsec-sa-name"`
    IkeGatewayName string `json:"ike-gateway-name"`
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

func (p *VpnIpsecSaOper) GetId() string{
    return "1"
}

func (p *VpnIpsecSaOper) getPath() string{
    return "vpn/ipsec-sa/oper"
}

func (p *VpnIpsecSaOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnIpsecSaOper,error) {
logger.Println("VpnIpsecSaOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnIpsecSaOper
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
