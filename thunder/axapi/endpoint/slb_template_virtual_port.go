

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateVirtualPort struct {
	Inst struct {

    Aflow int `json:"aflow"`
    AllowSynOtherflags int `json:"allow-syn-otherflags"`
    AllowVipToRportMapping int `json:"allow-vip-to-rport-mapping"`
    ConnLimit int `json:"conn-limit" dval:"64000000"`
    ConnLimitNoLogging int `json:"conn-limit-no-logging"`
    ConnLimitReset int `json:"conn-limit-reset"`
    ConnRateLimit int `json:"conn-rate-limit"`
    ConnRateLimitNoLogging int `json:"conn-rate-limit-no-logging"`
    ConnRateLimitReset int `json:"conn-rate-limit-reset"`
    DropUnknownConn int `json:"drop-unknown-conn"`
    Dscp int `json:"dscp"`
    IgnoreTcpMsl int `json:"ignore-tcp-msl"`
    LogOptions string `json:"log-options"`
    Name string `json:"name"`
    NonSynInitiation int `json:"non-syn-initiation"`
    PktRateInterval string `json:"pkt-rate-interval" dval:"second"`
    PktRateLimitReset int `json:"pkt-rate-limit-reset"`
    PktRateType string `json:"pkt-rate-type"`
    Rate int `json:"rate"`
    RateInterval string `json:"rate-interval" dval:"second"`
    ResetL7OnFailover int `json:"reset-l7-on-failover"`
    ResetUnknownConn int `json:"reset-unknown-conn"`
    SnatMsl int `json:"snat-msl"`
    SnatPortPreserve int `json:"snat-port-preserve"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    WhenRrEnable int `json:"when-rr-enable"`

	} `json:"virtual-port"`
}

func (p *SlbTemplateVirtualPort) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateVirtualPort) getPath() string{
    return "slb/template/virtual-port"
}

func (p *SlbTemplateVirtualPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateVirtualPort::Post")
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

func (p *SlbTemplateVirtualPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateVirtualPort::Get")
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
func (p *SlbTemplateVirtualPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateVirtualPort::Put")
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

func (p *SlbTemplateVirtualPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateVirtualPort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
