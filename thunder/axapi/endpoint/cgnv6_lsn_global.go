

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6LsnGlobal struct {
	Inst struct {

    AttemptPortPreservation string `json:"attempt-port-preservation"`
    EnhancedUserTracking int `json:"enhanced-user-tracking"`
    Hairpinning string `json:"hairpinning" dval:"filter-none"`
    HalfCloseTimeout int `json:"half-close-timeout"`
    Icmp Cgnv6LsnGlobalIcmp `json:"icmp"`
    InboundRefresh string `json:"inbound-refresh" dval:"enable"`
    InboundRefreshFullCone string `json:"inbound-refresh-full-cone" dval:"enable"`
    IpSelection string `json:"ip-selection" dval:"random"`
    Logging Cgnv6LsnGlobalLogging `json:"logging"`
    PortBatching Cgnv6LsnGlobalPortBatching `json:"port-batching"`
    SamplingEnable []Cgnv6LsnGlobalSamplingEnable `json:"sampling-enable"`
    StrictlyStickyNat int `json:"strictly-sticky-nat"`
    SynTimeout int `json:"syn-timeout" dval:"4"`
    Uuid string `json:"uuid"`

	} `json:"global"`
}


type Cgnv6LsnGlobalIcmp struct {
    SendOnPortUnavailable string `json:"send-on-port-unavailable" dval:"disable"`
    SendOnUserQuotaExceeded string `json:"send-on-user-quota-exceeded" dval:"admin-filtered"`
}


type Cgnv6LsnGlobalLogging struct {
    DefaultTemplate string `json:"default-template"`
    Pool []Cgnv6LsnGlobalLoggingPool `json:"pool"`
    Shared int `json:"shared"`
    PartitionName string `json:"partition-name"`
}


type Cgnv6LsnGlobalLoggingPool struct {
    PoolName string `json:"pool-name"`
    Template string `json:"template"`
}


type Cgnv6LsnGlobalPortBatching struct {
    Size string `json:"size" dval:"1"`
    TcpTimeWaitInterval int `json:"tcp-time-wait-interval" dval:"2"`
}


type Cgnv6LsnGlobalSamplingEnable struct {
    Counters1 string `json:"counters1"`
    Counters2 string `json:"counters2"`
    Counters3 string `json:"counters3"`
    Counters4 string `json:"counters4"`
}

func (p *Cgnv6LsnGlobal) GetId() string{
    return "1"
}

func (p *Cgnv6LsnGlobal) getPath() string{
    return "cgnv6/lsn/global"
}

func (p *Cgnv6LsnGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnGlobal::Post")
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

func (p *Cgnv6LsnGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnGlobal::Get")
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
func (p *Cgnv6LsnGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnGlobal::Put")
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

func (p *Cgnv6LsnGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6LsnGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
