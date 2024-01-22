

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6DdosProtection struct {
	Inst struct {

    DisableNatIpByBgp Cgnv6DdosProtectionDisableNatIpByBgp81 `json:"disable-nat-ip-by-bgp"`
    EnableAction string `json:"enable-action" dval:"local"`
    IpEntries Cgnv6DdosProtectionIpEntries82 `json:"ip-entries"`
    L4Entries Cgnv6DdosProtectionL4Entries83 `json:"l4-entries"`
    LoggingAction string `json:"logging-action" dval:"enable"`
    MaxHwEntries int `json:"max-hw-entries" dval:"262144"`
    PacketsPerSecond Cgnv6DdosProtectionPacketsPerSecond `json:"packets-per-second"`
    SamplingEnable []Cgnv6DdosProtectionSamplingEnable `json:"sampling-enable"`
    SynCookie Cgnv6DdosProtectionSynCookie `json:"syn-cookie"`
    Toggle string `json:"toggle" dval:"enable"`
    Uuid string `json:"uuid"`
    Zone string `json:"zone"`

	} `json:"ddos-protection"`
}


type Cgnv6DdosProtectionDisableNatIpByBgp81 struct {
    Uuid string `json:"uuid"`
}


type Cgnv6DdosProtectionIpEntries82 struct {
    Uuid string `json:"uuid"`
}


type Cgnv6DdosProtectionL4Entries83 struct {
    Uuid string `json:"uuid"`
}


type Cgnv6DdosProtectionPacketsPerSecond struct {
    Ip int `json:"ip" dval:"3000000"`
    Action Cgnv6DdosProtectionPacketsPerSecondAction `json:"action"`
    Tcp int `json:"tcp" dval:"3000"`
    TcpAction Cgnv6DdosProtectionPacketsPerSecondTcpAction `json:"tcp-action"`
    Udp int `json:"udp" dval:"3000"`
    UdpAction Cgnv6DdosProtectionPacketsPerSecondUdpAction `json:"udp-action"`
    Other int `json:"other" dval:"10000"`
    OtherAction Cgnv6DdosProtectionPacketsPerSecondOtherAction `json:"other-action"`
    IncludeExistingSession int `json:"include-existing-session"`
}


type Cgnv6DdosProtectionPacketsPerSecondAction struct {
    ActionType string `json:"action-type" dval:"drop"`
    RouteMap string `json:"route-map"`
    Expiration int `json:"expiration" dval:"3600"`
    ExpirationRoute int `json:"expiration-route" dval:"3600"`
    TimerMultiplyMax int `json:"timer-multiply-max" dval:"6"`
    RemoveWaitTimer int `json:"remove-wait-timer" dval:"300"`
}


type Cgnv6DdosProtectionPacketsPerSecondTcpAction struct {
    TcpActionType string `json:"tcp-action-type" dval:"drop"`
    TcpExpiration int `json:"tcp-expiration" dval:"30"`
}


type Cgnv6DdosProtectionPacketsPerSecondUdpAction struct {
    UdpActionType string `json:"udp-action-type" dval:"drop"`
    UdpExpiration int `json:"udp-expiration" dval:"30"`
}


type Cgnv6DdosProtectionPacketsPerSecondOtherAction struct {
    OtherActionType string `json:"other-action-type" dval:"drop"`
    OtherExpiration int `json:"other-expiration" dval:"30"`
}


type Cgnv6DdosProtectionSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type Cgnv6DdosProtectionSynCookie struct {
    SynCookieEnable int `json:"syn-cookie-enable"`
    SynCookieOnThreshold int `json:"syn-cookie-on-threshold"`
    SynCookieOnTimeout int `json:"syn-cookie-on-timeout" dval:"120"`
}

func (p *Cgnv6DdosProtection) GetId() string{
    return "1"
}

func (p *Cgnv6DdosProtection) getPath() string{
    return "cgnv6/ddos-protection"
}

func (p *Cgnv6DdosProtection) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6DdosProtection::Post")
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

func (p *Cgnv6DdosProtection) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6DdosProtection::Get")
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
func (p *Cgnv6DdosProtection) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6DdosProtection::Put")
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

func (p *Cgnv6DdosProtection) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6DdosProtection::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
