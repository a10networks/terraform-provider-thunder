

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnDefaultOper struct {
    
    Oper VpnDefaultOperOper `json:"oper"`

}
type DataVpnDefaultOper struct {
    DtVpnDefaultOper VpnDefaultOper `json:"default"`
}


type VpnDefaultOperOper struct {
    IkeVersion string `json:"ike-version"`
    IkeMode string `json:"ike-mode"`
    IkeDhGroup string `json:"ike-dh-group"`
    IkeAuthMethod string `json:"ike-auth-method"`
    IkeEncryption string `json:"ike-encryption"`
    IkeHash string `json:"ike-hash"`
    IkePriority int `json:"ike-priority"`
    IkeLifetime int `json:"ike-lifetime"`
    IkeNatTraversal string `json:"ike-nat-traversal"`
    IkeLocalAddress string `json:"ike-local-address"`
    IkeRemoteAddress string `json:"ike-remote-address"`
    IkeDpdInterval int `json:"ike-dpd-interval"`
    IpsecMode string `json:"IPsec-mode"`
    IpsecProtocol string `json:"IPsec-protocol"`
    IpsecDhGroup string `json:"IPsec-dh-group"`
    IpsecEncryption string `json:"IPsec-encryption"`
    IpsecHash string `json:"IPsec-hash"`
    IpsecPriority int `json:"IPsec-priority"`
    IpsecLifetime int `json:"IPsec-lifetime"`
    IpsecLifebytes int `json:"IPsec-lifebytes"`
    IpsecTrafficSelector string `json:"IPsec-traffic-selector"`
    IpsecLocalSubnet string `json:"IPsec-local-subnet"`
    IpsecLocalPort int `json:"IPsec-local-port"`
    IpsecLocalProtocol int `json:"IPsec-local-protocol"`
    IpsecRemoteSubnet string `json:"IPsec-remote-subnet"`
    IpsecRemotePort int `json:"IPsec-remote-port"`
    IpsecRemoteProtocol int `json:"IPsec-remote-protocol"`
    IpsecAntiReplayWindow int `json:"IPsec-anti-replay-window"`
}

func (p *VpnDefaultOper) GetId() string{
    return "1"
}

func (p *VpnDefaultOper) getPath() string{
    return "vpn/default/oper"
}

func (p *VpnDefaultOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataVpnDefaultOper,error) {
logger.Println("VpnDefaultOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataVpnDefaultOper
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
