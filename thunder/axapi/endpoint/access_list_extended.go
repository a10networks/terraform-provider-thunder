

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type AccessListExtended struct {
	Inst struct {

    Extd int `json:"extd"`
    Rules []AccessListExtendedRules `json:"rules"`
    Uuid string `json:"uuid"`

	} `json:"extended"`
}


type AccessListExtendedRules struct {
    ExtdSeqNum int `json:"extd-seq-num"`
    ExtdRemark string `json:"extd-remark"`
    ExtdAction string `json:"extd-action"`
    Icmp int `json:"icmp"`
    Tcp int `json:"tcp"`
    Udp int `json:"udp"`
    Ip int `json:"ip"`
    ServiceObjGroup string `json:"service-obj-group"`
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

func (p *AccessListExtended) GetId() string{
    return strconv.Itoa(p.Inst.Extd)
}

func (p *AccessListExtended) getPath() string{
    return "access-list/extended"
}

func (p *AccessListExtended) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AccessListExtended::Post")
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

func (p *AccessListExtended) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AccessListExtended::Get")
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
func (p *AccessListExtended) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("AccessListExtended::Put")
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

func (p *AccessListExtended) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("AccessListExtended::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
