

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type ObjectGroupService struct {
	Inst struct {

    Description string `json:"description"`
    Rules []ObjectGroupServiceRules `json:"rules"`
    SvcName string `json:"svc-name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"service"`
}


type ObjectGroupServiceRules struct {
    SeqNum int `json:"seq-num"`
    ProtocolId int `json:"protocol-id"`
    TcpUdp string `json:"tcp-udp"`
    Icmp int `json:"icmp"`
    Icmpv6 int `json:"icmpv6"`
    IcmpType int `json:"icmp-type"`
    AnyType int `json:"any-type"`
    SpecialType string `json:"special-type"`
    AnyCode int `json:"any-code"`
    IcmpCode int `json:"icmp-code"`
    SpecialCode string `json:"special-code"`
    Icmpv6Type int `json:"icmpv6-type"`
    V6AnyType int `json:"v6-any-type"`
    SpecialV6Type string `json:"special-v6-type"`
    V6AnyCode int `json:"v6-any-code"`
    Icmpv6Code int `json:"icmpv6-code"`
    SpecialV6Code string `json:"special-v6-code"`
    Source int `json:"source"`
    EqSrc int `json:"eq-src"`
    GtSrc int `json:"gt-src"`
    LtSrc int `json:"lt-src"`
    RangeSrc int `json:"range-src"`
    PortNumEndSrc int `json:"port-num-end-src"`
    EqDst int `json:"eq-dst"`
    GtDst int `json:"gt-dst"`
    LtDst int `json:"lt-dst"`
    RangeDst int `json:"range-dst"`
    PortNumEndDst int `json:"port-num-end-dst"`
    Alg string `json:"alg"`
}

func (p *ObjectGroupService) GetId() string{
    return url.QueryEscape(p.Inst.SvcName)
}

func (p *ObjectGroupService) getPath() string{
    return "object-group/service"
}

func (p *ObjectGroupService) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ObjectGroupService::Post")
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

func (p *ObjectGroupService) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ObjectGroupService::Get")
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
func (p *ObjectGroupService) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ObjectGroupService::Put")
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

func (p *ObjectGroupService) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ObjectGroupService::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
