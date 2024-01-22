

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6FixedNatInsideIplist struct {
	Inst struct {

    DestRuleList string `json:"dest-rule-list"`
    DynamicPoolSize int `json:"dynamic-pool-size"`
    InsideIpList string `json:"inside-ip-list"`
    Method string `json:"method" dval:"use-least-nat-ips"`
    NatEndAddress string `json:"nat-end-address"`
    NatIpList string `json:"nat-ip-list"`
    NatNetmask string `json:"nat-netmask"`
    NatStartAddress string `json:"nat-start-address"`
    Offset Cgnv6FixedNatInsideIplistOffset `json:"offset"`
    Partition string `json:"partition"`
    PortsPerUser int `json:"ports-per-user"`
    RespondToUserMac int `json:"respond-to-user-mac"`
    SessionQuota int `json:"session-quota"`
    SkipPortsOnRollover int `json:"skip-ports-on-rollover"`
    UsableNatPorts Cgnv6FixedNatInsideIplistUsableNatPorts `json:"usable-nat-ports"`
    Uuid string `json:"uuid"`
    Vrid int `json:"vrid"`

	} `json:"iplist"`
}


type Cgnv6FixedNatInsideIplistOffset struct {
    Random int `json:"random"`
    NumericOffset int `json:"numeric-offset"`
}


type Cgnv6FixedNatInsideIplistUsableNatPorts struct {
    UsableStartPort int `json:"usable-start-port"`
    UsableEndPort int `json:"usable-end-port"`
}

func (p *Cgnv6FixedNatInsideIplist) GetId() string{
    return url.QueryEscape(p.Inst.InsideIpList)+"+"+p.Inst.Partition
}

func (p *Cgnv6FixedNatInsideIplist) getPath() string{
    return "cgnv6/fixed-nat/inside/iplist"
}

func (p *Cgnv6FixedNatInsideIplist) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatInsideIplist::Post")
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

func (p *Cgnv6FixedNatInsideIplist) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatInsideIplist::Get")
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
func (p *Cgnv6FixedNatInsideIplist) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatInsideIplist::Put")
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

func (p *Cgnv6FixedNatInsideIplist) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6FixedNatInsideIplist::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
