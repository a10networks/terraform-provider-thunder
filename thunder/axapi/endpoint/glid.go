

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Glid struct {
	Inst struct {

    BitRateLimit int `json:"bit-rate-limit"`
    ConnLimit int `json:"conn-limit"`
    ConnRateLimit int `json:"conn-rate-limit"`
    ConnRateLimitInterval int `json:"conn-rate-limit-interval"`
    Description string `json:"description"`
    Dns GlidDns `json:"dns"`
    Dns64 GlidDns64 `json:"dns64"`
    FragPktRateLimit int `json:"frag-pkt-rate-limit"`
    Name string `json:"name"`
    OverLimitCfg GlidOverLimitCfg `json:"over-limit-cfg"`
    PktRateLimit int `json:"pkt-rate-limit"`
    RateUnit string `json:"rate-unit" dval:"system-global-setting"`
    RequestLimit int `json:"request-limit"`
    RequestRateLimit int `json:"request-rate-limit"`
    RequestRateLimitInterval int `json:"request-rate-limit-interval"`
    SynCookieThr int `json:"syn-cookie-thr"`
    UseNatPool string `json:"use-nat-pool"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"glid"`
}


type GlidDns struct {
    Action string `json:"action" dval:"cache-disable"`
    Weight int `json:"weight"`
    Ttl int `json:"ttl"`
}


type GlidDns64 struct {
    Disable int `json:"disable"`
    ExclusiveAnswer int `json:"exclusive-answer"`
    Prefix string `json:"prefix"`
}


type GlidOverLimitCfg struct {
    OverLimitAction int `json:"over-limit-action"`
    ActionType string `json:"action-type"`
    BlacklistSrcMin int `json:"blacklist-src-min"`
    ActionValue string `json:"action-value"`
    Lockout int `json:"lockout"`
    Log int `json:"log"`
    LogInterval int `json:"log-interval"`
}

func (p *Glid) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Glid) getPath() string{
    return "glid"
}

func (p *Glid) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Glid::Post")
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

func (p *Glid) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Glid::Get")
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
func (p *Glid) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Glid::Put")
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

func (p *Glid) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Glid::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
