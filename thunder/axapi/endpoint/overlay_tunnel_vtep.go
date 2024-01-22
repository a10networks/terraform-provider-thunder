

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type OverlayTunnelVtep struct {
	Inst struct {

    DestPort int `json:"dest-port"`
    Encap string `json:"encap"`
    HostList []OverlayTunnelVtepHostList `json:"host-list"`
    Id1 int `json:"id"`
    LocalIpAddress OverlayTunnelVtepLocalIpAddress1085 `json:"local-ip-address"`
    LocalIpv6Address OverlayTunnelVtepLocalIpv6Address1087 `json:"local-ipv6-address"`
    RemoteIpAddressList []OverlayTunnelVtepRemoteIpAddressList `json:"remote-ip-address-list"`
    RemoteIpv6AddressList []OverlayTunnelVtepRemoteIpv6AddressList `json:"remote-ipv6-address-list"`
    SamplingEnable []OverlayTunnelVtepSamplingEnable `json:"sampling-enable"`
    SrcPortRange OverlayTunnelVtepSrcPortRange1088 `json:"src-port-range"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"vtep"`
}


type OverlayTunnelVtepHostList struct {
    IpAddr string `json:"ip-addr"`
    OverlayMacAddr string `json:"overlay-mac-addr"`
    Vni int `json:"vni"`
    RemoteVtep string `json:"remote-vtep"`
    Uuid string `json:"uuid"`
}


type OverlayTunnelVtepLocalIpAddress1085 struct {
    IpAddress string `json:"ip-address"`
    Uuid string `json:"uuid"`
    VniList []OverlayTunnelVtepLocalIpAddressVniList1086 `json:"vni-list"`
}


type OverlayTunnelVtepLocalIpAddressVniList1086 struct {
    Segment int `json:"segment"`
    Partition string `json:"partition"`
    Gateway int `json:"gateway"`
    Lif string `json:"lif"`
    Uuid string `json:"uuid"`
}


type OverlayTunnelVtepLocalIpv6Address1087 struct {
    Ipv6Address string `json:"ipv6-address"`
    Uuid string `json:"uuid"`
}


type OverlayTunnelVtepRemoteIpAddressList struct {
    IpAddress string `json:"ip-address"`
    Encap string `json:"encap"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    UseLif OverlayTunnelVtepRemoteIpAddressListUseLif `json:"use-lif"`
    GreKeepalive OverlayTunnelVtepRemoteIpAddressListGreKeepalive `json:"gre-keepalive"`
    UseGreKey OverlayTunnelVtepRemoteIpAddressListUseGreKey `json:"use-gre-key"`
    VniList []OverlayTunnelVtepRemoteIpAddressListVniList `json:"vni-list"`
}


type OverlayTunnelVtepRemoteIpAddressListUseLif struct {
    Partition string `json:"partition"`
    Lif string `json:"lif"`
    Uuid string `json:"uuid"`
}


type OverlayTunnelVtepRemoteIpAddressListGreKeepalive struct {
    RetryTime int `json:"retry-time" dval:"10"`
    RetryCount int `json:"retry-count"`
    Uuid string `json:"uuid"`
}


type OverlayTunnelVtepRemoteIpAddressListUseGreKey struct {
    GreKey int `json:"gre-key"`
    Uuid string `json:"uuid"`
}


type OverlayTunnelVtepRemoteIpAddressListVniList struct {
    Segment int `json:"segment"`
    Uuid string `json:"uuid"`
}


type OverlayTunnelVtepRemoteIpv6AddressList struct {
    Ipv6Address string `json:"ipv6-address"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    UseLif OverlayTunnelVtepRemoteIpv6AddressListUseLif `json:"use-lif"`
}


type OverlayTunnelVtepRemoteIpv6AddressListUseLif struct {
    Partition string `json:"partition"`
    Lif string `json:"lif"`
    Uuid string `json:"uuid"`
}


type OverlayTunnelVtepSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type OverlayTunnelVtepSrcPortRange1088 struct {
    MinPort int `json:"min-port" dval:"1"`
    MaxPort int `json:"max-port" dval:"65535"`
    Uuid string `json:"uuid"`
}

func (p *OverlayTunnelVtep) GetId() string{
    return strconv.Itoa(p.Inst.Id1)
}

func (p *OverlayTunnelVtep) getPath() string{
    return "overlay-tunnel/vtep"
}

func (p *OverlayTunnelVtep) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtep::Post")
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

func (p *OverlayTunnelVtep) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtep::Get")
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
func (p *OverlayTunnelVtep) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtep::Put")
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

func (p *OverlayTunnelVtep) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("OverlayTunnelVtep::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
