

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnIkeGatewayOper struct {
    
    Name string `json:"name"`
    Oper VpnIkeGatewayOperOper `json:"oper"`

}
type DataVpnIkeGatewayOper struct {
    DtVpnIkeGatewayOper VpnIkeGatewayOper `json:"ike-gateway"`
}


type VpnIkeGatewayOperOper struct {
    RemoteIpFilter string `json:"remote-ip-filter"`
    RemoteIdFilter string `json:"remote-id-filter"`
    BriefFilter string `json:"brief-filter"`
    SaList []VpnIkeGatewayOperOperSaList `json:"SA-List"`
}


type VpnIkeGatewayOperOperSaList struct {
    InitiatorSpi string `json:"Initiator-SPI"`
    ResponderSpi string `json:"Responder-SPI"`
    LocalIp string `json:"Local-IP"`
    RemoteIp string `json:"Remote-IP"`
    Encryption string `json:"Encryption"`
    Hash string `json:"Hash"`
    SignHash string `json:"Sign-hash"`
    Lifetime int `json:"Lifetime"`
    Status string `json:"Status"`
    NatTraversal int `json:"NAT-Traversal"`
    RemoteId string `json:"Remote-ID"`
    DhGroup int `json:"DH-Group"`
    FragmentMessageGenerated int `json:"Fragment-message-generated"`
    FragmentMessageReceived int `json:"Fragment-message-received"`
    FragmentationError int `json:"Fragmentation-error"`
    FragmentReassembleError int `json:"Fragment-reassemble-error"`
}

func (p *VpnIkeGatewayOper) GetId() string{
    return "1"
}

func (p *VpnIkeGatewayOper) getPath() string{
    
    return "vpn/ike-gateway/"+p.Name+"/oper"
}

func (p *VpnIkeGatewayOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnIkeGatewayOper,error) {
logger.Println("VpnIkeGatewayOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnIkeGatewayOper
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
