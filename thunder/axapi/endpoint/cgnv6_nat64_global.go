

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6Nat64Global struct {
	Inst struct {

    All int `json:"all"`
    ForceNonZeroIpv4Id int `json:"force-non-zero-ipv4-id"`
    Icmp Cgnv6Nat64GlobalIcmp `json:"icmp"`
    Inside Cgnv6Nat64GlobalInside `json:"inside"`
    SamplingEnable []Cgnv6Nat64GlobalSamplingEnable `json:"sampling-enable"`
    Tcp Cgnv6Nat64GlobalTcp `json:"tcp"`
    UserQuotaPrefixLength int `json:"user-quota-prefix-length" dval:"128"`
    Uuid string `json:"uuid"`

	} `json:"global"`
}


type Cgnv6Nat64GlobalIcmp struct {
    SendOnPortUnavailable string `json:"send-on-port-unavailable" dval:"disable"`
    SendOnUserQuotaExceeded string `json:"send-on-user-quota-exceeded" dval:"admin-filtered"`
}


type Cgnv6Nat64GlobalInside struct {
    Source Cgnv6Nat64GlobalInsideSource `json:"source"`
}


type Cgnv6Nat64GlobalInsideSource struct {
    ClassList string `json:"class-list"`
}


type Cgnv6Nat64GlobalSamplingEnable struct {
    Counters1 string `json:"counters1"`
    Counters2 string `json:"counters2"`
    Counters3 string `json:"counters3"`
}


type Cgnv6Nat64GlobalTcp struct {
    MssClamp Cgnv6Nat64GlobalTcpMssClamp `json:"mss-clamp"`
    ResetOnError Cgnv6Nat64GlobalTcpResetOnError `json:"reset-on-error"`
}


type Cgnv6Nat64GlobalTcpMssClamp struct {
    MssClampType string `json:"mss-clamp-type" dval:"subtract"`
    MssValue int `json:"mss-value"`
    MssSubtract int `json:"mss-subtract" dval:"20"`
    Min int `json:"min" dval:"476"`
}


type Cgnv6Nat64GlobalTcpResetOnError struct {
    Outbound string `json:"outbound"`
}

func (p *Cgnv6Nat64Global) GetId() string{
    return "1"
}

func (p *Cgnv6Nat64Global) getPath() string{
    return "cgnv6/nat64/global"
}

func (p *Cgnv6Nat64Global) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Nat64Global::Post")
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

func (p *Cgnv6Nat64Global) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Nat64Global::Get")
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
func (p *Cgnv6Nat64Global) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Nat64Global::Put")
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

func (p *Cgnv6Nat64Global) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6Nat64Global::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
