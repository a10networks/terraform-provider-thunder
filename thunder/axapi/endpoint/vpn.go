

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Vpn struct {
	Inst struct {

    AsymmetricFlowSupport int `json:"asymmetric-flow-support"`
    Crl VpnCrl3614 `json:"crl"`
    Default VpnDefault3615 `json:"default"`
    EnableVpnMetrics int `json:"enable-vpn-metrics"`
    Error VpnError3616 `json:"error"`
    Errordump VpnErrordump3617 `json:"errordump"`
    ExtendedMatching int `json:"extended-matching"`
    FragmentAfterEncap int `json:"fragment-after-encap"`
    GroupList VpnGroupList3618 `json:"group-list"`
    IkeAccEnable int `json:"ike-acc-enable"`
    IkeGatewayList []VpnIkeGatewayList `json:"ike-gateway-list"`
    IkeLoggingEnable int `json:"ike-logging-enable"`
    IkeSa VpnIkeSa3619 `json:"ike-sa"`
    IkeSaBrief VpnIkeSaBrief3620 `json:"ike-sa-brief"`
    IkeSaClients VpnIkeSaClients3621 `json:"ike-sa-clients"`
    IkeSaTimeout int `json:"ike-sa-timeout" dval:"600"`
    IkeStatsByGw VpnIkeStatsByGw3622 `json:"ike-stats-by-gw"`
    IkeStatsGlobal VpnIkeStatsGlobal3623 `json:"ike-stats-global"`
    IpsecCipherCheck int `json:"ipsec-cipher-check"`
    IpsecErrorDump int `json:"ipsec-error-dump"`
    IpsecGroupList []VpnIpsecGroupList `json:"ipsec-group-list"`
    IpsecList []VpnIpsecList `json:"ipsec-list"`
    IpsecMgmtDefaultPolicyDrop int `json:"ipsec-mgmt-default-policy-drop"`
    IpsecSa VpnIpsecSa3625 `json:"ipsec-sa"`
    IpsecSaByGw VpnIpsecSaByGw3626 `json:"ipsec-sa-by-gw"`
    IpsecSaClients VpnIpsecSaClients3627 `json:"ipsec-sa-clients"`
    IpsecSaStatsList []VpnIpsecSaStatsList `json:"ipsec-sa-stats-list"`
    JumboFragment int `json:"jumbo-fragment"`
    Log VpnLog3628 `json:"log"`
    NatTraversalFlowAffinity int `json:"nat-traversal-flow-affinity"`
    Ocsp VpnOcsp3629 `json:"ocsp"`
    RevocationList []VpnRevocationList `json:"revocation-list"`
    SamplingEnable []VpnSamplingEnable `json:"sampling-enable"`
    SignatureAuthentication int `json:"signature-authentication"`
    StatefulMode int `json:"stateful-mode"`
    TcpMssAdjustDisable int `json:"tcp-mss-adjust-disable"`
    Uuid string `json:"uuid"`

	} `json:"vpn"`
}


type VpnCrl3614 struct {
    Uuid string `json:"uuid"`
}


type VpnDefault3615 struct {
    Uuid string `json:"uuid"`
}


type VpnError3616 struct {
    Uuid string `json:"uuid"`
}


type VpnErrordump3617 struct {
    Uuid string `json:"uuid"`
}


type VpnGroupList3618 struct {
    Uuid string `json:"uuid"`
}


type VpnIkeGatewayList struct {
    Name string `json:"name"`
    IkeVersion string `json:"ike-version" dval:"v2"`
    Mode string `json:"mode" dval:"main"`
    AuthMethod string `json:"auth-method" dval:"preshare-key"`
    PreshareKeyValue string `json:"preshare-key-value"`
    PreshareKeyEncrypted string `json:"preshare-key-encrypted"`
    Hash string `json:"hash"`
    InterfaceManagement int `json:"interface-management"`
    Key string `json:"key"`
    KeyPassphrase string `json:"key-passphrase"`
    KeyPassphraseEncrypted string `json:"key-passphrase-encrypted"`
    Vrid VpnIkeGatewayListVrid `json:"vrid"`
    LocalCert VpnIkeGatewayListLocalCert `json:"local-cert"`
    RemoteCaCert VpnIkeGatewayListRemoteCaCert `json:"remote-ca-cert"`
    LocalId string `json:"local-id"`
    RemoteId string `json:"remote-id"`
    EncCfg []VpnIkeGatewayListEncCfg `json:"enc-cfg"`
    DhGroup string `json:"dh-group" dval:"1"`
    LocalAddress VpnIkeGatewayListLocalAddress `json:"local-address"`
    RemoteAddress VpnIkeGatewayListRemoteAddress `json:"remote-address"`
    Lifetime int `json:"lifetime" dval:"86400"`
    FragmentSize int `json:"fragment-size"`
    NatTraversal int `json:"nat-traversal"`
    Dpd VpnIkeGatewayListDpd `json:"dpd"`
    DisableRekey int `json:"disable-rekey"`
    ConfigurationPayload string `json:"configuration-payload"`
    DhcpServer VpnIkeGatewayListDhcpServer `json:"dhcp-server"`
    RadiusServer VpnIkeGatewayListRadiusServer `json:"radius-server"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []VpnIkeGatewayListSamplingEnable `json:"sampling-enable"`
}


type VpnIkeGatewayListVrid struct {
    Default int `json:"default"`
    VridNum int `json:"vrid-num"`
}


type VpnIkeGatewayListLocalCert struct {
    LocalCertName string `json:"local-cert-name"`
}


type VpnIkeGatewayListRemoteCaCert struct {
    RemoteCertName string `json:"remote-cert-name"`
}


type VpnIkeGatewayListEncCfg struct {
    Encryption string `json:"encryption"`
    Hash string `json:"hash"`
    Prf string `json:"prf"`
    Priority int `json:"priority" dval:"5"`
    Gcm_priority int `json:"gcm_priority" dval:"5"`
}


type VpnIkeGatewayListLocalAddress struct {
    LocalIp string `json:"local-ip"`
    LocalIpv6 string `json:"local-ipv6"`
}


type VpnIkeGatewayListRemoteAddress struct {
    RemoteIp string `json:"remote-ip"`
    Dns string `json:"dns"`
    RemoteIpv6 string `json:"remote-ipv6"`
}


type VpnIkeGatewayListDpd struct {
    Interval int `json:"interval"`
    Retry int `json:"retry"`
}


type VpnIkeGatewayListDhcpServer struct {
    Pri VpnIkeGatewayListDhcpServerPri `json:"pri"`
    Sec VpnIkeGatewayListDhcpServerSec `json:"sec"`
}


type VpnIkeGatewayListDhcpServerPri struct {
    DhcpPriIpv4 string `json:"dhcp-pri-ipv4"`
}


type VpnIkeGatewayListDhcpServerSec struct {
    DhcpSecIpv4 string `json:"dhcp-sec-ipv4"`
}


type VpnIkeGatewayListRadiusServer struct {
    RadiusPri string `json:"radius-pri"`
    RadiusSec string `json:"radius-sec"`
}


type VpnIkeGatewayListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type VpnIkeSa3619 struct {
    Uuid string `json:"uuid"`
}


type VpnIkeSaBrief3620 struct {
    Uuid string `json:"uuid"`
}


type VpnIkeSaClients3621 struct {
    Uuid string `json:"uuid"`
}


type VpnIkeStatsByGw3622 struct {
    Uuid string `json:"uuid"`
}


type VpnIkeStatsGlobal3623 struct {
    Uuid string `json:"uuid"`
    SamplingEnable []VpnIkeStatsGlobalSamplingEnable3624 `json:"sampling-enable"`
}


type VpnIkeStatsGlobalSamplingEnable3624 struct {
    Counters1 string `json:"counters1"`
}


type VpnIpsecGroupList struct {
    Name string `json:"name"`
    IpsecgroupCfg []VpnIpsecGroupListIpsecgroupCfg `json:"ipsecgroup-cfg"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type VpnIpsecGroupListIpsecgroupCfg struct {
    Ipsec string `json:"ipsec"`
    Priority int `json:"priority"`
}


type VpnIpsecList struct {
    Name string `json:"name"`
    Mode string `json:"mode" dval:"tunnel"`
    Dscp string `json:"dscp"`
    Proto string `json:"proto" dval:"esp"`
    DhGroup string `json:"dh-group" dval:"0"`
    EncCfg []VpnIpsecListEncCfg `json:"enc-cfg"`
    Lifetime int `json:"lifetime" dval:"28800"`
    Lifebytes int `json:"lifebytes"`
    AntiReplayWindow string `json:"anti-replay-window" dval:"0"`
    Up int `json:"up"`
    SequenceNumberDisable int `json:"sequence-number-disable"`
    TrafficSelector VpnIpsecListTrafficSelector `json:"traffic-selector"`
    EnforceTrafficSelector int `json:"enforce-traffic-selector"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    SamplingEnable []VpnIpsecListSamplingEnable `json:"sampling-enable"`
    BindTunnel VpnIpsecListBindTunnel `json:"bind-tunnel"`
    IpsecGateway VpnIpsecListIpsecGateway `json:"ipsec-gateway"`
}


type VpnIpsecListEncCfg struct {
    Encryption string `json:"encryption"`
    Hash string `json:"hash"`
    Priority int `json:"priority" dval:"5"`
    Gcm_priority int `json:"gcm_priority" dval:"5"`
}


type VpnIpsecListTrafficSelector struct {
    Ipv4 VpnIpsecListTrafficSelectorIpv4 `json:"ipv4"`
    Ipv6 VpnIpsecListTrafficSelectorIpv6 `json:"ipv6"`
}


type VpnIpsecListTrafficSelectorIpv4 struct {
    Local string `json:"local"`
    Local_netmask string `json:"local_netmask"`
    Local_port int `json:"local_port"`
    RemoteIpv4Assigned int `json:"remote-ipv4-assigned"`
    RemoteIp string `json:"remote-ip"`
    Remote_netmask string `json:"remote_netmask"`
    Remote_port int `json:"remote_port"`
    Protocol int `json:"protocol"`
}


type VpnIpsecListTrafficSelectorIpv6 struct {
    Localv6 string `json:"localv6"`
    Local_portv6 int `json:"local_portv6"`
    RemoteIpv6Assigned int `json:"remote-ipv6-assigned"`
    RemoteIpv6 string `json:"remote-ipv6"`
    Remote_portv6 int `json:"remote_portv6"`
    Protocolv6 int `json:"protocolv6"`
}


type VpnIpsecListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type VpnIpsecListBindTunnel struct {
    Tunnel int `json:"tunnel"`
    NextHop string `json:"next-hop"`
    NextHopV6 string `json:"next-hop-v6"`
    Uuid string `json:"uuid"`
}


type VpnIpsecListIpsecGateway struct {
    IkeGateway string `json:"ike-gateway"`
    Uuid string `json:"uuid"`
}


type VpnIpsecSa3625 struct {
    Uuid string `json:"uuid"`
}


type VpnIpsecSaByGw3626 struct {
    Uuid string `json:"uuid"`
}


type VpnIpsecSaClients3627 struct {
    Uuid string `json:"uuid"`
}


type VpnIpsecSaStatsList struct {
    SamplingEnable []VpnIpsecSaStatsListSamplingEnable `json:"sampling-enable"`
}


type VpnIpsecSaStatsListSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type VpnLog3628 struct {
    Uuid string `json:"uuid"`
}


type VpnOcsp3629 struct {
    Uuid string `json:"uuid"`
}


type VpnRevocationList struct {
    Name string `json:"name"`
    Ca string `json:"ca"`
    Crl VpnRevocationListCrl `json:"crl"`
    Ocsp VpnRevocationListOcsp `json:"ocsp"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type VpnRevocationListCrl struct {
    CrlPri string `json:"crl-pri"`
    CrlSec string `json:"crl-sec"`
}


type VpnRevocationListOcsp struct {
    OcspPri string `json:"ocsp-pri"`
    OcspSec string `json:"ocsp-sec"`
}


type VpnSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *Vpn) GetId() string{
    return "1"
}

func (p *Vpn) getPath() string{
    return "vpn"
}

func (p *Vpn) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Vpn::Post")
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

func (p *Vpn) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Vpn::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
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
func (p *Vpn) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Vpn::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), "", payloadBytes, headers, logger)
    return err
}

func (p *Vpn) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Vpn::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
