

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateCache struct {
	Inst struct {

    AcceptReloadReq int `json:"accept-reload-req"`
    Age int `json:"age" dval:"3600"`
    DefaultPolicyNocache int `json:"default-policy-nocache"`
    DisableInsertAge int `json:"disable-insert-age"`
    DisableInsertVia int `json:"disable-insert-via"`
    LocalUriPolicy []SlbTemplateCacheLocalUriPolicy `json:"local-uri-policy"`
    Logging string `json:"logging"`
    MaxCacheSize int `json:"max-cache-size" dval:"80"`
    MaxContentSize int `json:"max-content-size" dval:"81920"`
    MinContentSize int `json:"min-content-size" dval:"512"`
    Name string `json:"name"`
    PacketCaptureTemplate string `json:"packet-capture-template"`
    RemoveCookies int `json:"remove-cookies"`
    ReplacementPolicy string `json:"replacement-policy" dval:"LFU"`
    SamplingEnable []SlbTemplateCacheSamplingEnable `json:"sampling-enable"`
    UriPolicy []SlbTemplateCacheUriPolicy `json:"uri-policy"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VerifyHost int `json:"verify-host"`

	} `json:"cache"`
}


type SlbTemplateCacheLocalUriPolicy struct {
    LocalUri string `json:"local-uri"`
}


type SlbTemplateCacheSamplingEnable struct {
    Counters1 string `json:"counters1"`
}


type SlbTemplateCacheUriPolicy struct {
    Uri string `json:"uri"`
    CacheAction string `json:"cache-action"`
    CacheValue int `json:"cache-value"`
    Invalidate string `json:"invalidate"`
}

func (p *SlbTemplateCache) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateCache) getPath() string{
    return "slb/template/cache"
}

func (p *SlbTemplateCache) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateCache::Post")
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

func (p *SlbTemplateCache) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateCache::Get")
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
func (p *SlbTemplateCache) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateCache::Put")
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

func (p *SlbTemplateCache) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateCache::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
