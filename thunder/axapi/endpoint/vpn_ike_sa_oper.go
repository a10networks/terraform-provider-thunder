

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnIkeSaOper struct {
    
    Oper VpnIkeSaOperOper `json:"oper"`

}
type DataVpnIkeSaOper struct {
    DtVpnIkeSaOper VpnIkeSaOper `json:"ike-sa"`
}


type VpnIkeSaOperOper struct {
    IkeSaList []VpnIkeSaOperOperIkeSaList `json:"ike-sa-list"`
}


type VpnIkeSaOperOperIkeSaList struct {
    Name string `json:"Name"`
    InitiatorSpi string `json:"Initiator-SPI"`
    ResponderSpi string `json:"Responder-SPI"`
    LocalIp string `json:"Local-IP"`
    RemoteIp string `json:"Remote-IP"`
    Encryption string `json:"Encryption"`
    Hash string `json:"Hash"`
    Lifetime int `json:"Lifetime"`
    Status string `json:"Status"`
    NatTraversal int `json:"NAT-Traversal"`
}

func (p *VpnIkeSaOper) GetId() string{
    return "1"
}

func (p *VpnIkeSaOper) getPath() string{
    return "vpn/ike-sa/oper"
}

func (p *VpnIkeSaOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnIkeSaOper,error) {
logger.Println("VpnIkeSaOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnIkeSaOper
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
