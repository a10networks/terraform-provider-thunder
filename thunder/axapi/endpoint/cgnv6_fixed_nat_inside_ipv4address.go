

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatInsideIpv4address struct {
	Inst struct {

    DestRuleList string `json:"dest-rule-list"`
    DynamicPoolSize int `json:"dynamic-pool-size"`
    InsideEndAddress string `json:"inside-end-address"`
    InsideNetmask string `json:"inside-netmask"`
    InsideStartAddress string `json:"inside-start-address"`
    Method string `json:"method" dval:"use-least-nat-ips"`
    NatEndAddress string `json:"nat-end-address"`
    NatIpList string `json:"nat-ip-list"`
    NatNetmask string `json:"nat-netmask"`
    NatStartAddress string `json:"nat-start-address"`
    Offset Cgnv6FixedNatInsideIpv4addressOffset `json:"offset"`
    Partition string `json:"partition"`
    PortsPerUser int `json:"ports-per-user"`
    RespondToUserMac int `json:"respond-to-user-mac"`
    SessionQuota int `json:"session-quota"`
    SkipPortsOnRollover int `json:"skip-ports-on-rollover"`
    UsableNatPorts Cgnv6FixedNatInsideIpv4addressUsableNatPorts `json:"usable-nat-ports"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"ipv4address"`
}


type Cgnv6FixedNatInsideIpv4addressOffset struct {
    Random int `json:"random"`
    NumericOffset int `json:"numeric-offset"`
}


type Cgnv6FixedNatInsideIpv4addressUsableNatPorts struct {
    UsableStartPort int `json:"usable-start-port"`
    UsableEndPort int `json:"usable-end-port"`
}

func (p *Cgnv6FixedNatInsideIpv4address) GetId() string{
    return p.Inst.InsideStartAddress+"+"+p.Inst.InsideEndAddress+"+"+url.QueryEscape(p.Inst.InsideNetmask)+"+"+p.Inst.Partition
}

func (p *Cgnv6FixedNatInsideIpv4address) getPath() string{
    return "cgnv6/fixed-nat/inside/ipv4address"
}

func (p *Cgnv6FixedNatInsideIpv4address) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatInsideIpv4address::Post")
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

func (p *Cgnv6FixedNatInsideIpv4address) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatInsideIpv4address::Get")
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
func (p *Cgnv6FixedNatInsideIpv4address) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatInsideIpv4address::Put")
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

func (p *Cgnv6FixedNatInsideIpv4address) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatInsideIpv4address::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
