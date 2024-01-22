

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpAccessList struct {
	Inst struct {

    Name string `json:"name"`
    Rules []IpAccessListRules `json:"rules"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"access-list"`
}


type IpAccessListRules struct {
    SeqNum int `json:"seq-num"`
    Action string `json:"action"`
    Remark string `json:"remark"`
    Icmp int `json:"icmp"`
    Tcp int `json:"tcp"`
    Udp int `json:"udp"`
    Ip int `json:"ip"`
    ServiceObjGroup string `json:"service-obj-group"`
    GeoLocation string `json:"geo-location"`
    IcmpType int `json:"icmp-type"`
    AnyType int `json:"any-type"`
    SpecialType string `json:"special-type"`
    AnyCode int `json:"any-code"`
    IcmpCode int `json:"icmp-code"`
    SpecialCode string `json:"special-code"`
    SrcAny int `json:"src-any"`
    SrcHost string `json:"src-host"`
    SrcSubnet string `json:"src-subnet"`
    SrcMask string `json:"src-mask"`
    SrcObjectGroup string `json:"src-object-group"`
    SrcEq int `json:"src-eq"`
    SrcGt int `json:"src-gt"`
    SrcLt int `json:"src-lt"`
    SrcRange int `json:"src-range"`
    SrcPortEnd int `json:"src-port-end"`
    DstAny int `json:"dst-any"`
    DstHost string `json:"dst-host"`
    DstSubnet string `json:"dst-subnet"`
    DstMask string `json:"dst-mask"`
    DstObjectGroup string `json:"dst-object-group"`
    DstEq int `json:"dst-eq"`
    DstGt int `json:"dst-gt"`
    DstLt int `json:"dst-lt"`
    DstRange int `json:"dst-range"`
    DstPortEnd int `json:"dst-port-end"`
    Fragments int `json:"fragments"`
    Vlan int `json:"vlan"`
    Ethernet int `json:"ethernet"`
    Trunk int `json:"trunk"`
    Dscp int `json:"dscp"`
    Established int `json:"established"`
    AclLog int `json:"acl-log"`
    TransparentSessionOnly int `json:"transparent-session-only"`
}

func (p *IpAccessList) GetId() string{
    return p.Inst.Name
}

func (p *IpAccessList) getPath() string{
    return "ip/access-list"
}

func (p *IpAccessList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAccessList::Post")
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

func (p *IpAccessList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAccessList::Get")
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
func (p *IpAccessList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAccessList::Put")
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

func (p *IpAccessList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAccessList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
