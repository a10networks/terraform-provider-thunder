

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnIkeGateway struct {
	Inst struct {

    AuthMethod string `json:"auth-method" dval:"preshare-key"`
    ConfigurationPayload string `json:"configuration-payload"`
    DhGroup string `json:"dh-group" dval:"1"`
    DhcpServer VpnIkeGatewayDhcpServer `json:"dhcp-server"`
    DisableRekey int `json:"disable-rekey"`
    Dpd VpnIkeGatewayDpd `json:"dpd"`
    EncCfg []VpnIkeGatewayEncCfg `json:"enc-cfg"`
    FragmentSize int `json:"fragment-size"`
    Hash string `json:"hash"`
    IkeVersion string `json:"ike-version" dval:"v2"`
    InterfaceManagement int `json:"interface-management"`
    Key string `json:"key"`
    KeyPassphrase string `json:"key-passphrase"`
    KeyPassphraseEncrypted string `json:"key-passphrase-encrypted"`
    Lifetime int `json:"lifetime" dval:"86400"`
    LocalAddress VpnIkeGatewayLocalAddress `json:"local-address"`
    LocalCert VpnIkeGatewayLocalCert `json:"local-cert"`
    LocalId string `json:"local-id"`
    Mode string `json:"mode" dval:"main"`
    Name string `json:"name"`
    NatTraversal int `json:"nat-traversal"`
    PreshareKeyEncrypted string `json:"preshare-key-encrypted"`
    PreshareKeyValue string `json:"preshare-key-value"`
    RadiusServer VpnIkeGatewayRadiusServer `json:"radius-server"`
    RemoteAddress VpnIkeGatewayRemoteAddress `json:"remote-address"`
    RemoteCaCert VpnIkeGatewayRemoteCaCert `json:"remote-ca-cert"`
    RemoteId string `json:"remote-id"`
    SamplingEnable []VpnIkeGatewaySamplingEnable `json:"sampling-enable"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Vrid VpnIkeGatewayVrid `json:"vrid"`

	} `json:"ike-gateway"`
}


type VpnIkeGatewayDhcpServer struct {
    Pri VpnIkeGatewayDhcpServerPri `json:"pri"`
    Sec VpnIkeGatewayDhcpServerSec `json:"sec"`
}


type VpnIkeGatewayDhcpServerPri struct {
    DhcpPriIpv4 string `json:"dhcp-pri-ipv4"`
}


type VpnIkeGatewayDhcpServerSec struct {
    DhcpSecIpv4 string `json:"dhcp-sec-ipv4"`
}


type VpnIkeGatewayDpd struct {
    Interval int `json:"interval"`
    Retry int `json:"retry"`
}


type VpnIkeGatewayEncCfg struct {
    Encryption string `json:"encryption"`
    Hash string `json:"hash"`
    Prf string `json:"prf"`
    Priority int `json:"priority" dval:"5"`
    Gcm_priority int `json:"gcm_priority" dval:"5"`
}


type VpnIkeGatewayLocalAddress struct {
    LocalIp string `json:"local-ip"`
    LocalIpv6 string `json:"local-ipv6"`
}


type VpnIkeGatewayLocalCert struct {
    LocalCertName string `json:"local-cert-name"`
}


type VpnIkeGatewayRadiusServer struct {
    RadiusPri string `json:"radius-pri"`
    RadiusSec string `json:"radius-sec"`
}


type VpnIkeGatewayRemoteAddress struct {
    RemoteIp string `json:"remote-ip"`
    Dns string `json:"dns"`
    RemoteIpv6 string `json:"remote-ipv6"`
}


type VpnIkeGatewayRemoteCaCert struct {
    RemoteCertName string `json:"remote-cert-name"`
}


type VpnIkeGatewaySamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type VpnIkeGatewayVrid struct {
    Default int `json:"default"`
    VridNum int `json:"vrid-num"`
}

func (p *VpnIkeGateway) GetId() string{
    return p.Inst.Name
}

func (p *VpnIkeGateway) getPath() string{
    return "vpn/ike-gateway"
}

func (p *VpnIkeGateway) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VpnIkeGateway::Post")
    headers := axapi.GenRequestHeader(authToken)
        payloadBytes, err := axapi.SerializeToJson(p)
        if err != nil {
            logger.Println("Failed to serialize struct as json", err)
            return err
        }
        logger.Println("payload:", string(payloadBytes))
        _, _, err = axapi.SendPost(host, p.getPath(), payloadBytes, headers, logger)
    return err
}

func (p *VpnIkeGateway) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VpnIkeGateway::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return err
}
func (p *VpnIkeGateway) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VpnIkeGateway::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *VpnIkeGateway) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VpnIkeGateway::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
