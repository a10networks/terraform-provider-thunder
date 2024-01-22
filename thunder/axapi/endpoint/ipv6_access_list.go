

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Ipv6AccessList struct {
	Inst struct {

    Name string `json:"name"`
    Rules []Ipv6AccessListRules `json:"rules"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"access-list"`
}


type Ipv6AccessListRules struct {
    SeqNum int `json:"seq-num"`
    Action string `json:"action"`
    Remark string `json:"remark"`
    Icmp int `json:"icmp"`
    Tcp int `json:"tcp"`
    Udp int `json:"udp"`
    Ipv6 int `json:"ipv6"`
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
    SrcObjectGroup string `json:"src-object-group"`
    SrcEq int `json:"src-eq"`
    SrcGt int `json:"src-gt"`
    SrcLt int `json:"src-lt"`
    SrcRange int `json:"src-range"`
    SrcPortEnd int `json:"src-port-end"`
    DstAny int `json:"dst-any"`
    DstHost string `json:"dst-host"`
    DstSubnet string `json:"dst-subnet"`
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
}

func (p *Ipv6AccessList) GetId() string{
    return p.Inst.Name
}

func (p *Ipv6AccessList) getPath() string{
    return "ipv6/access-list"
}

func (p *Ipv6AccessList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6AccessList::Post")
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

func (p *Ipv6AccessList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6AccessList::Get")
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
func (p *Ipv6AccessList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6AccessList::Put")
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

func (p *Ipv6AccessList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Ipv6AccessList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
