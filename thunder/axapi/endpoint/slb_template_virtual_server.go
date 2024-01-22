

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateVirtualServer struct {
	Inst struct {

    ConnLimit int `json:"conn-limit" dval:"64000000"`
    ConnLimitNoLogging int `json:"conn-limit-no-logging"`
    ConnLimitReset int `json:"conn-limit-reset"`
    ConnRateLimit int `json:"conn-rate-limit"`
    ConnRateLimitNoLogging int `json:"conn-rate-limit-no-logging"`
    ConnRateLimitReset int `json:"conn-rate-limit-reset"`
    DisableWhenAllPortsDown int `json:"disable-when-all-ports-down"`
    DisableWhenAnyPortDown int `json:"disable-when-any-port-down"`
    IcmpLockup int `json:"icmp-lockup"`
    IcmpLockupPeriod int `json:"icmp-lockup-period"`
    IcmpRateLimit int `json:"icmp-rate-limit"`
    Icmpv6Lockup int `json:"icmpv6-lockup"`
    Icmpv6LockupPeriod int `json:"icmpv6-lockup-period"`
    Icmpv6RateLimit int `json:"icmpv6-rate-limit"`
    Name string `json:"name"`
    RateInterval string `json:"rate-interval" dval:"second"`
    SubnetGratuitousArp int `json:"subnet-gratuitous-arp"`
    TcpStackTfoActiveConnLimit int `json:"tcp-stack-tfo-active-conn-limit"`
    TcpStackTfoBackoffTime int `json:"tcp-stack-tfo-backoff-time" dval:"600"`
    TcpStackTfoCookieTimeLimit int `json:"tcp-stack-tfo-cookie-time-limit" dval:"60"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"virtual-server"`
}

func (p *SlbTemplateVirtualServer) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateVirtualServer) getPath() string{
    return "slb/template/virtual-server"
}

func (p *SlbTemplateVirtualServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateVirtualServer::Post")
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

func (p *SlbTemplateVirtualServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateVirtualServer::Get")
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
func (p *SlbTemplateVirtualServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateVirtualServer::Put")
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

func (p *SlbTemplateVirtualServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateVirtualServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
