

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplatePort struct {
	Inst struct {

    Add int `json:"add"`
    BwRateLimit int `json:"bw-rate-limit"`
    BwRateLimitDuration int `json:"bw-rate-limit-duration"`
    BwRateLimitNoLogging int `json:"bw-rate-limit-no-logging"`
    BwRateLimitResume int `json:"bw-rate-limit-resume"`
    ConnLimit int `json:"conn-limit" dval:"64000000"`
    ConnLimitNoLogging int `json:"conn-limit-no-logging"`
    ConnRateLimit int `json:"conn-rate-limit"`
    ConnRateLimitNoLogging int `json:"conn-rate-limit-no-logging"`
    DampeningFlaps int `json:"dampening-flaps"`
    Decrement int `json:"decrement"`
    DelSessionOnServerDown int `json:"del-session-on-server-down"`
    DestNat int `json:"dest-nat"`
    DownGracePeriod int `json:"down-grace-period"`
    DownTimer int `json:"down-timer"`
    Dscp int `json:"dscp"`
    DynamicMemberPriority int `json:"dynamic-member-priority" dval:"16"`
    Every int `json:"every" dval:"10"`
    ExtendedStats int `json:"extended-stats"`
    FlapPeriod int `json:"flap-period"`
    HealthCheck string `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    InbandHealthCheck int `json:"inband-health-check"`
    InitialSlowStart int `json:"initial-slow-start" dval:"128"`
    Name string `json:"name"`
    NoSsl int `json:"no-ssl"`
    RateInterval string `json:"rate-interval" dval:"second"`
    Reassign int `json:"reassign" dval:"25"`
    RequestRateInterval string `json:"request-rate-interval" dval:"second"`
    RequestRateLimit int `json:"request-rate-limit"`
    RequestRateNoLogging int `json:"request-rate-no-logging"`
    ReselOnReset int `json:"resel-on-reset"`
    Reset int `json:"reset"`
    RestoreSvcTime int `json:"restore-svc-time"`
    Resume int `json:"resume"`
    Retry int `json:"retry" dval:"2"`
    SharedPartitionPool int `json:"shared-partition-pool"`
    SlowStart int `json:"slow-start"`
    SourceNat string `json:"source-nat"`
    StatsDataAction string `json:"stats-data-action" dval:"stats-data-enable"`
    SubGroup int `json:"sub-group"`
    TemplatePortPoolShared string `json:"template-port-pool-shared"`
    Till int `json:"till" dval:"4096"`
    Times int `json:"times" dval:"2"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Weight int `json:"weight" dval:"1"`

	} `json:"port"`
}

func (p *SlbTemplatePort) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplatePort) getPath() string{
    return "slb/template/port"
}

func (p *SlbTemplatePort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePort::Post")
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

func (p *SlbTemplatePort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePort::Get")
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
func (p *SlbTemplatePort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePort::Put")
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

func (p *SlbTemplatePort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplatePort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
