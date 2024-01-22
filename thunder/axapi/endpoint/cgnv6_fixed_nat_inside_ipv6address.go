

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatInsideIpv6address struct {
	Inst struct {

    DestRuleList string `json:"dest-rule-list"`
    DynamicPoolSize int `json:"dynamic-pool-size"`
    InsideEndAddress string `json:"inside-end-address"`
    InsideNetmask int `json:"inside-netmask"`
    InsideStartAddress string `json:"inside-start-address"`
    Method string `json:"method" dval:"use-least-nat-ips"`
    NatEndAddress string `json:"nat-end-address"`
    NatIpList string `json:"nat-ip-list"`
    NatNetmask string `json:"nat-netmask"`
    NatStartAddress string `json:"nat-start-address"`
    Offset Cgnv6FixedNatInsideIpv6addressOffset `json:"offset"`
    Partition string `json:"partition"`
    PortsPerUser int `json:"ports-per-user"`
    RespondToUserMac int `json:"respond-to-user-mac"`
    SessionQuota int `json:"session-quota"`
    SkipPortsOnRollover int `json:"skip-ports-on-rollover"`
    UsableNatPorts Cgnv6FixedNatInsideIpv6addressUsableNatPorts `json:"usable-nat-ports"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"ipv6address"`
}


type Cgnv6FixedNatInsideIpv6addressOffset struct {
    Random int `json:"random"`
    NumericOffset int `json:"numeric-offset"`
}


type Cgnv6FixedNatInsideIpv6addressUsableNatPorts struct {
    UsableStartPort int `json:"usable-start-port"`
    UsableEndPort int `json:"usable-end-port"`
}

func (p *Cgnv6FixedNatInsideIpv6address) GetId() string{
    return p.Inst.InsideStartAddress+"+"+p.Inst.InsideEndAddress+"+"+strconv.Itoa(p.Inst.InsideNetmask)+"+"+p.Inst.Partition
}

func (p *Cgnv6FixedNatInsideIpv6address) getPath() string{
    return "cgnv6/fixed-nat/inside/ipv6address"
}

func (p *Cgnv6FixedNatInsideIpv6address) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatInsideIpv6address::Post")
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

func (p *Cgnv6FixedNatInsideIpv6address) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatInsideIpv6address::Get")
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
func (p *Cgnv6FixedNatInsideIpv6address) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatInsideIpv6address::Put")
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

func (p *Cgnv6FixedNatInsideIpv6address) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatInsideIpv6address::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
