

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnIpsecSaClientsOper struct {
    
    Oper VpnIpsecSaClientsOperOper `json:"oper"`

}
type DataVpnIpsecSaClientsOper struct {
    DtVpnIpsecSaClientsOper VpnIpsecSaClientsOper `json:"ipsec-sa-clients"`
}


type VpnIpsecSaClientsOperOper struct {
    IpsecClients []VpnIpsecSaClientsOperOperIpsecClients `json:"ipsec-clients"`
}


type VpnIpsecSaClientsOperOperIpsecClients struct {
    IpsecClientsIp string `json:"ipsec-clients-ip"`
    SaList []VpnIpsecSaClientsOperOperIpsecClientsSaList `json:"sa-list"`
}


type VpnIpsecSaClientsOperOperIpsecClientsSaList struct {
    Name string `json:"name"`
    LocalTs string `json:"local-ts"`
    InSpi string `json:"in-spi"`
    OutSpi string `json:"out-spi"`
    Lifetime string `json:"lifetime"`
    Lifebytes string `json:"lifebytes"`
}

func (p *VpnIpsecSaClientsOper) GetId() string{
    return "1"
}

func (p *VpnIpsecSaClientsOper) getPath() string{
    return "vpn/ipsec-sa-clients/oper"
}

func (p *VpnIpsecSaClientsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnIpsecSaClientsOper,error) {
logger.Println("VpnIpsecSaClientsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnIpsecSaClientsOper
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
