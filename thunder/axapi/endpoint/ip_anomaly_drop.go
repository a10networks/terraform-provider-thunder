

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpAnomalyDrop struct {
	Inst struct {

    BadContent int `json:"bad-content"`
    DropAll int `json:"drop-all"`
    Frag int `json:"frag"`
    IpOption int `json:"ip-option"`
    Ipv6ExtHeader IpAnomalyDropIpv6ExtHeader `json:"ipv6-ext-header"`
    LandAttack int `json:"land-attack"`
    OutOfSequence int `json:"out-of-sequence"`
    PacketDeformity IpAnomalyDropPacketDeformity `json:"packet-deformity"`
    PingOfDeath int `json:"ping-of-death"`
    SamplingEnable []IpAnomalyDropSamplingEnable `json:"sampling-enable"`
    SecurityAttack IpAnomalyDropSecurityAttack `json:"security-attack"`
    TcpNoFlag int `json:"tcp-no-flag"`
    TcpSynFin int `json:"tcp-syn-fin"`
    TcpSynFrag int `json:"tcp-syn-frag"`
    Uuid string `json:"uuid"`
    ZeroWindow int `json:"zero-window"`

	} `json:"anomaly-drop"`
}


type IpAnomalyDropIpv6ExtHeader struct {
    Ipv6EhFrag int `json:"ipv6-eh-frag"`
    Ipv6EhAuth int `json:"ipv6-eh-auth"`
    Ipv6EhEsp int `json:"ipv6-eh-esp"`
    Ipv6EhMobility int `json:"ipv6-eh-mobility"`
    Ipv6EhNonext int `json:"ipv6-eh-nonext"`
    Ipv6EhMalformed int `json:"ipv6-eh-malformed"`
    Ipv6EhHbh int `json:"ipv6-eh-hbh"`
    Ipv6EhDest int `json:"ipv6-eh-dest"`
    Ipv6EhRouting int `json:"ipv6-eh-routing"`
    HbhOptionList []IpAnomalyDropIpv6ExtHeaderHbhOptionList `json:"hbh-option-list"`
    DstOptionList []IpAnomalyDropIpv6ExtHeaderDstOptionList `json:"dst-option-list"`
    RoutingOptionList []IpAnomalyDropIpv6ExtHeaderRoutingOptionList `json:"routing-option-list"`
    UnknownExtHeaderList []IpAnomalyDropIpv6ExtHeaderUnknownExtHeaderList `json:"unknown-ext-header-list"`
}


type IpAnomalyDropIpv6ExtHeaderHbhOptionList struct {
    HbhOtypeFrom int `json:"hbh-otype-from"`
    HbhOtypeTo int `json:"hbh-otype-to"`
}


type IpAnomalyDropIpv6ExtHeaderDstOptionList struct {
    DstOtypeFrom int `json:"dst-otype-from"`
    DstOtypeTo int `json:"dst-otype-to"`
}


type IpAnomalyDropIpv6ExtHeaderRoutingOptionList struct {
    RoutingOtypeFrom int `json:"routing-otype-from"`
    RoutingOtypeTo int `json:"routing-otype-to"`
}


type IpAnomalyDropIpv6ExtHeaderUnknownExtHeaderList struct {
    EhTypeFrom int `json:"eh-type-from"`
    EhTypeTo int `json:"eh-type-to"`
}


type IpAnomalyDropPacketDeformity struct {
    PacketDeformityLayer3 int `json:"packet-deformity-layer-3"`
    PacketDeformityLayer4 int `json:"packet-deformity-layer-4"`
}


type IpAnomalyDropSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type IpAnomalyDropSecurityAttack struct {
    SecurityAttackLayer3 int `json:"security-attack-layer-3"`
    SecurityAttackLayer4 int `json:"security-attack-layer-4"`
}

func (p *IpAnomalyDrop) GetId() string{
    return "1"
}

func (p *IpAnomalyDrop) getPath() string{
    return "ip/anomaly-drop"
}

func (p *IpAnomalyDrop) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAnomalyDrop::Post")
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

func (p *IpAnomalyDrop) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAnomalyDrop::Get")
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
func (p *IpAnomalyDrop) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAnomalyDrop::Put")
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

func (p *IpAnomalyDrop) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAnomalyDrop::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
