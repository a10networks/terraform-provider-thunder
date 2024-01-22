

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnLid struct {
	Inst struct {

    ConnRateLimit Cgnv6LsnLidConnRateLimit `json:"conn-rate-limit"`
    DsLite Cgnv6LsnLidDsLite `json:"ds-lite"`
    ExtendedUserQuota Cgnv6LsnLidExtendedUserQuota `json:"extended-user-quota"`
    LidNumber int `json:"lid-number"`
    LsnRuleList Cgnv6LsnLidLsnRuleList `json:"lsn-rule-list"`
    Name string `json:"name"`
    Override string `json:"override" dval:"none"`
    RespondToUserMac int `json:"respond-to-user-mac"`
    SourceNatPool Cgnv6LsnLidSourceNatPool `json:"source-nat-pool"`
    UserQuota Cgnv6LsnLidUserQuota `json:"user-quota"`
    UserQuotaPrefixLength int `json:"user-quota-prefix-length"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"lsn-lid"`
}


type Cgnv6LsnLidConnRateLimit struct {
    ConnRateLimitVal int `json:"conn-rate-limit-val"`
}


type Cgnv6LsnLidDsLite struct {
    InsideSrcPermitList string `json:"inside-src-permit-list"`
}


type Cgnv6LsnLidExtendedUserQuota struct {
    Tcp []Cgnv6LsnLidExtendedUserQuotaTcp `json:"tcp"`
    Udp []Cgnv6LsnLidExtendedUserQuotaUdp `json:"udp"`
}


type Cgnv6LsnLidExtendedUserQuotaTcp struct {
    TcpServicePort int `json:"tcp-service-port"`
    TcpSessions int `json:"tcp-sessions"`
}


type Cgnv6LsnLidExtendedUserQuotaUdp struct {
    UdpServicePort int `json:"udp-service-port"`
    UdpSessions int `json:"udp-sessions"`
}


type Cgnv6LsnLidLsnRuleList struct {
    Destination string `json:"destination"`
}


type Cgnv6LsnLidSourceNatPool struct {
    PoolName string `json:"pool-name"`
    Shared int `json:"shared"`
}


type Cgnv6LsnLidUserQuota struct {
    Icmp int `json:"icmp"`
    QuotaUdp Cgnv6LsnLidUserQuotaQuotaUdp `json:"quota-udp"`
    QuotaTcp Cgnv6LsnLidUserQuotaQuotaTcp `json:"quota-tcp"`
    Session int `json:"session"`
}


type Cgnv6LsnLidUserQuotaQuotaUdp struct {
    UdpQuota int `json:"udp-quota"`
    UdpReserve int `json:"udp-reserve"`
}


type Cgnv6LsnLidUserQuotaQuotaTcp struct {
    TcpQuota int `json:"tcp-quota"`
    TcpReserve int `json:"tcp-reserve"`
}

func (p *Cgnv6LsnLid) GetId() string{
    return strconv.Itoa(p.Inst.LidNumber)
}

func (p *Cgnv6LsnLid) getPath() string{
    return "cgnv6/lsn-lid"
}

func (p *Cgnv6LsnLid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnLid::Post")
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

func (p *Cgnv6LsnLid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnLid::Get")
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
func (p *Cgnv6LsnLid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnLid::Put")
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

func (p *Cgnv6LsnLid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnLid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
