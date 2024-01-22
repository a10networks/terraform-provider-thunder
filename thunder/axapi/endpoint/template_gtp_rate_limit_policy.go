

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type TemplateGtpRateLimitPolicy struct {
	Inst struct {

    GtpUDownlinkByteRate int `json:"gtp-u-downlink-byte-rate"`
    GtpUDownlinkPacketRate int `json:"gtp-u-downlink-packet-rate"`
    GtpUMaxConcurrentTunnels int `json:"gtp-u-max-concurrent-tunnels"`
    GtpUTotalByteRate int `json:"gtp-u-total-byte-rate"`
    GtpUTotalPacketRate int `json:"gtp-u-total-packet-rate"`
    GtpUTunnelCreateRate int `json:"gtp-u-tunnel-create-rate"`
    GtpUUplinkByteRate int `json:"gtp-u-uplink-byte-rate"`
    GtpUUplinkPacketRate int `json:"gtp-u-uplink-packet-rate"`
    Lockout int `json:"lockout"`
    Name string `json:"name"`
    RateLimitAction string `json:"rate-limit-action" dval:"drop"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    V0AggMessageTypeRate int `json:"v0-agg-message-type-rate"`
    V1AggMessageTypeRate int `json:"v1-agg-message-type-rate"`
    V1CreatePdpRequestRate int `json:"v1-create-pdp-request-rate"`
    V1UpdatePdpRequestRate int `json:"v1-update-pdp-request-rate"`
    V2AggMessageTypeRate int `json:"v2-agg-message-type-rate"`
    V2CreateSessionRequestRate int `json:"v2-create-session-request-rate"`
    V2ModifyBearerRequestRate int `json:"v2-modify-bearer-request-rate"`

	} `json:"rate-limit-policy"`
}

func (p *TemplateGtpRateLimitPolicy) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *TemplateGtpRateLimitPolicy) getPath() string{
    return "template/gtp/rate-limit-policy"
}

func (p *TemplateGtpRateLimitPolicy) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpRateLimitPolicy::Post")
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

func (p *TemplateGtpRateLimitPolicy) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpRateLimitPolicy::Get")
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
func (p *TemplateGtpRateLimitPolicy) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpRateLimitPolicy::Put")
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

func (p *TemplateGtpRateLimitPolicy) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TemplateGtpRateLimitPolicy::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
