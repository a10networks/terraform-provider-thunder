

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateServer struct {
	Inst struct {

    Add int `json:"add"`
    BwRateLimit int `json:"bw-rate-limit"`
    BwRateLimitAcct string `json:"bw-rate-limit-acct" dval:"all"`
    BwRateLimitDuration int `json:"bw-rate-limit-duration"`
    BwRateLimitNoLogging int `json:"bw-rate-limit-no-logging"`
    BwRateLimitResume int `json:"bw-rate-limit-resume"`
    ConnLimit int `json:"conn-limit" dval:"64000000"`
    ConnLimitNoLogging int `json:"conn-limit-no-logging"`
    ConnRateLimit int `json:"conn-rate-limit"`
    ConnRateLimitNoLogging int `json:"conn-rate-limit-no-logging"`
    DnsFailInterval int `json:"dns-fail-interval" dval:"30"`
    DnsQueryInterval int `json:"dns-query-interval" dval:"10"`
    DynamicServerPrefix string `json:"dynamic-server-prefix" dval:"DRS"`
    Every int `json:"every" dval:"10"`
    ExtendedStats int `json:"extended-stats"`
    HealthCheck string `json:"health-check"`
    HealthCheckDisable int `json:"health-check-disable"`
    InitialSlowStart int `json:"initial-slow-start" dval:"128"`
    LogSelectionFailure int `json:"log-selection-failure"`
    MaxDynamicServer int `json:"max-dynamic-server" dval:"255"`
    MinTtlRatio int `json:"min-ttl-ratio" dval:"2"`
    Name string `json:"name"`
    RateInterval string `json:"rate-interval" dval:"second"`
    Resume int `json:"resume"`
    SlowStart int `json:"slow-start"`
    SpoofingCache int `json:"spoofing-cache"`
    StatsDataAction string `json:"stats-data-action" dval:"stats-data-enable"`
    Till int `json:"till" dval:"4096"`
    Times int `json:"times" dval:"2"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Weight int `json:"weight" dval:"1"`

	} `json:"server"`
}

func (p *SlbTemplateServer) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateServer) getPath() string{
    return "slb/template/server"
}

func (p *SlbTemplateServer) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateServer::Post")
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

func (p *SlbTemplateServer) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateServer::Get")
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
func (p *SlbTemplateServer) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateServer::Put")
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

func (p *SlbTemplateServer) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateServer::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
