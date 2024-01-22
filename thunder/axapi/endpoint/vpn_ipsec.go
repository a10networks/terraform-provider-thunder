

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type VpnIpsec struct {
	Inst struct {

    AntiReplayWindow string `json:"anti-replay-window" dval:"0"`
    BindTunnel VpnIpsecBindTunnel3612 `json:"bind-tunnel"`
    DhGroup string `json:"dh-group" dval:"0"`
    Dscp string `json:"dscp"`
    EncCfg []VpnIpsecEncCfg `json:"enc-cfg"`
    EnforceTrafficSelector int `json:"enforce-traffic-selector"`
    IpsecGateway VpnIpsecIpsecGateway3613 `json:"ipsec-gateway"`
    Lifebytes int `json:"lifebytes"`
    Lifetime int `json:"lifetime" dval:"28800"`
    Mode string `json:"mode" dval:"tunnel"`
    Name string `json:"name"`
    Proto string `json:"proto" dval:"esp"`
    SamplingEnable []VpnIpsecSamplingEnable `json:"sampling-enable"`
    SequenceNumberDisable int `json:"sequence-number-disable"`
    TrafficSelector VpnIpsecTrafficSelector `json:"traffic-selector"`
    Up int `json:"up"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"ipsec"`
}


type VpnIpsecBindTunnel3612 struct {
    Tunnel int `json:"tunnel"`
    NextHop string `json:"next-hop"`
    NextHopV6 string `json:"next-hop-v6"`
    Uuid string `json:"uuid"`
}


type VpnIpsecEncCfg struct {
    Encryption string `json:"encryption"`
    Hash string `json:"hash"`
    Priority int `json:"priority" dval:"5"`
    Gcm_priority int `json:"gcm_priority" dval:"5"`
}


type VpnIpsecIpsecGateway3613 struct {
    IkeGateway string `json:"ike-gateway"`
    Uuid string `json:"uuid"`
}


type VpnIpsecSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type VpnIpsecTrafficSelector struct {
    Ipv4 VpnIpsecTrafficSelectorIpv4 `json:"ipv4"`
    Ipv6 VpnIpsecTrafficSelectorIpv6 `json:"ipv6"`
}


type VpnIpsecTrafficSelectorIpv4 struct {
    Local string `json:"local"`
    Local_netmask string `json:"local_netmask"`
    Local_port int `json:"local_port"`
    RemoteIpv4Assigned int `json:"remote-ipv4-assigned"`
    RemoteIp string `json:"remote-ip"`
    Remote_netmask string `json:"remote_netmask"`
    Remote_port int `json:"remote_port"`
    Protocol int `json:"protocol"`
}


type VpnIpsecTrafficSelectorIpv6 struct {
    Localv6 string `json:"localv6"`
    Local_portv6 int `json:"local_portv6"`
    RemoteIpv6Assigned int `json:"remote-ipv6-assigned"`
    RemoteIpv6 string `json:"remote-ipv6"`
    Remote_portv6 int `json:"remote_portv6"`
    Protocolv6 int `json:"protocolv6"`
}

func (p *VpnIpsec) GetId() string{
    return p.Inst.Name
}

func (p *VpnIpsec) getPath() string{
    return "vpn/ipsec"
}

func (p *VpnIpsec) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VpnIpsec::Post")
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

func (p *VpnIpsec) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VpnIpsec::Get")
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
func (p *VpnIpsec) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("VpnIpsec::Put")
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

func (p *VpnIpsec) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("VpnIpsec::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
