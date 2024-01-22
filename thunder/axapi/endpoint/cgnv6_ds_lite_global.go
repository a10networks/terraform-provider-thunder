

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6DsLiteGlobal struct {
	Inst struct {

    Icmp Cgnv6DsLiteGlobalIcmp `json:"icmp"`
    Inside Cgnv6DsLiteGlobalInside `json:"inside"`
    IpChecksumError string `json:"ip-checksum-error" dval:"fix"`
    L4ChecksumError string `json:"l4-checksum-error" dval:"propagate"`
    SamplingEnable []Cgnv6DsLiteGlobalSamplingEnable `json:"sampling-enable"`
    Tcp Cgnv6DsLiteGlobalTcp `json:"tcp"`
    UserQuotaPrefixLength int `json:"user-quota-prefix-length" dval:"128"`
    Uuid string `json:"uuid"`

	} `json:"global"`
}


type Cgnv6DsLiteGlobalIcmp struct {
    SendOnPortUnavailable string `json:"send-on-port-unavailable" dval:"disable"`
    SendOnUserQuotaExceeded string `json:"send-on-user-quota-exceeded" dval:"admin-filtered"`
}


type Cgnv6DsLiteGlobalInside struct {
    Source Cgnv6DsLiteGlobalInsideSource `json:"source"`
}


type Cgnv6DsLiteGlobalInsideSource struct {
    ClassList string `json:"class-list"`
}


type Cgnv6DsLiteGlobalSamplingEnable struct {
    Counters1 string `json:"counters1"`
    Counters2 string `json:"counters2"`
    Counters3 string `json:"counters3"`
}


type Cgnv6DsLiteGlobalTcp struct {
    MssClamp Cgnv6DsLiteGlobalTcpMssClamp `json:"mss-clamp"`
    ResetOnError Cgnv6DsLiteGlobalTcpResetOnError `json:"reset-on-error"`
}


type Cgnv6DsLiteGlobalTcpMssClamp struct {
    MssClampType string `json:"mss-clamp-type" dval:"subtract"`
    MssValue int `json:"mss-value"`
    MssSubtract int `json:"mss-subtract" dval:"40"`
    Min int `json:"min" dval:"416"`
}


type Cgnv6DsLiteGlobalTcpResetOnError struct {
    Outbound string `json:"outbound"`
}

func (p *Cgnv6DsLiteGlobal) GetId() string{
    return "1"
}

func (p *Cgnv6DsLiteGlobal) getPath() string{
    return "cgnv6/ds-lite/global"
}

func (p *Cgnv6DsLiteGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6DsLiteGlobal::Post")
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

func (p *Cgnv6DsLiteGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6DsLiteGlobal::Get")
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
func (p *Cgnv6DsLiteGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6DsLiteGlobal::Put")
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

func (p *Cgnv6DsLiteGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6DsLiteGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
