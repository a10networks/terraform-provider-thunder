

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneTemplateHttpFilter struct {
	Inst struct {

    Dst DdosZoneTemplateHttpFilterDst `json:"dst"`
    HttpAgentCfg DdosZoneTemplateHttpFilterHttpAgentCfg `json:"http-agent-cfg"`
    HttpFilterAction string `json:"http-filter-action"`
    HttpFilterActionListName string `json:"http-filter-action-list-name"`
    HttpFilterName string `json:"http-filter-name"`
    HttpFilterSeq int `json:"http-filter-seq"`
    HttpHeaderCfg DdosZoneTemplateHttpFilterHttpHeaderCfg `json:"http-header-cfg"`
    HttpRefererCfg DdosZoneTemplateHttpFilterHttpRefererCfg `json:"http-referer-cfg"`
    HttpUriCfg DdosZoneTemplateHttpFilterHttpUriCfg `json:"http-uri-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    HttpTmplName string 

	} `json:"filter"`
}


type DdosZoneTemplateHttpFilterDst struct {
    HttpFilterRateLimit int `json:"http-filter-rate-limit"`
}


type DdosZoneTemplateHttpFilterHttpAgentCfg struct {
    AgentEqualsCfg []DdosZoneTemplateHttpFilterHttpAgentCfgAgentEqualsCfg `json:"agent-equals-cfg"`
    AgentContainsCfg []DdosZoneTemplateHttpFilterHttpAgentCfgAgentContainsCfg `json:"agent-contains-cfg"`
    AgentStartsCfg []DdosZoneTemplateHttpFilterHttpAgentCfgAgentStartsCfg `json:"agent-starts-cfg"`
    AgentEndsCfg []DdosZoneTemplateHttpFilterHttpAgentCfgAgentEndsCfg `json:"agent-ends-cfg"`
}


type DdosZoneTemplateHttpFilterHttpAgentCfgAgentEqualsCfg struct {
    HttpFilterAgentEquals string `json:"http-filter-agent-equals"`
}


type DdosZoneTemplateHttpFilterHttpAgentCfgAgentContainsCfg struct {
    HttpFilterAgentContains string `json:"http-filter-agent-contains"`
}


type DdosZoneTemplateHttpFilterHttpAgentCfgAgentStartsCfg struct {
    HttpFilterAgentStartsWith string `json:"http-filter-agent-starts-with"`
}


type DdosZoneTemplateHttpFilterHttpAgentCfgAgentEndsCfg struct {
    HttpFilterAgentEndsWith string `json:"http-filter-agent-ends-with"`
}


type DdosZoneTemplateHttpFilterHttpHeaderCfg struct {
    HttpFilterHeaderRegex string `json:"http-filter-header-regex"`
    HttpFilterHeaderInverseMatch int `json:"http-filter-header-inverse-match"`
}


type DdosZoneTemplateHttpFilterHttpRefererCfg struct {
    RefererEqualsCfg []DdosZoneTemplateHttpFilterHttpRefererCfgRefererEqualsCfg `json:"referer-equals-cfg"`
    RefererContainsCfg []DdosZoneTemplateHttpFilterHttpRefererCfgRefererContainsCfg `json:"referer-contains-cfg"`
    RefererStartsCfg []DdosZoneTemplateHttpFilterHttpRefererCfgRefererStartsCfg `json:"referer-starts-cfg"`
    RefererEndsCfg []DdosZoneTemplateHttpFilterHttpRefererCfgRefererEndsCfg `json:"referer-ends-cfg"`
}


type DdosZoneTemplateHttpFilterHttpRefererCfgRefererEqualsCfg struct {
    HttpFilterRefererEquals string `json:"http-filter-referer-equals"`
}


type DdosZoneTemplateHttpFilterHttpRefererCfgRefererContainsCfg struct {
    HttpFilterRefererContains string `json:"http-filter-referer-contains"`
}


type DdosZoneTemplateHttpFilterHttpRefererCfgRefererStartsCfg struct {
    HttpFilterRefererStartsWith string `json:"http-filter-referer-starts-with"`
}


type DdosZoneTemplateHttpFilterHttpRefererCfgRefererEndsCfg struct {
    HttpFilterRefererEndsWith string `json:"http-filter-referer-ends-with"`
}


type DdosZoneTemplateHttpFilterHttpUriCfg struct {
    UriEqualCfg []DdosZoneTemplateHttpFilterHttpUriCfgUriEqualCfg `json:"uri-equal-cfg"`
    UriContainsCfg []DdosZoneTemplateHttpFilterHttpUriCfgUriContainsCfg `json:"uri-contains-cfg"`
    UriStartsCfg []DdosZoneTemplateHttpFilterHttpUriCfgUriStartsCfg `json:"uri-starts-cfg"`
    UriEndsCfg []DdosZoneTemplateHttpFilterHttpUriCfgUriEndsCfg `json:"uri-ends-cfg"`
}


type DdosZoneTemplateHttpFilterHttpUriCfgUriEqualCfg struct {
    HttpFilterUriEquals string `json:"http-filter-uri-equals"`
}


type DdosZoneTemplateHttpFilterHttpUriCfgUriContainsCfg struct {
    HttpFilterUriContains string `json:"http-filter-uri-contains"`
}


type DdosZoneTemplateHttpFilterHttpUriCfgUriStartsCfg struct {
    HttpFilterUriStartsWith string `json:"http-filter-uri-starts-with"`
}


type DdosZoneTemplateHttpFilterHttpUriCfgUriEndsCfg struct {
    HttpFilterUriEndsWith string `json:"http-filter-uri-ends-with"`
}

func (p *DdosZoneTemplateHttpFilter) GetId() string{
    return url.QueryEscape(p.Inst.HttpFilterName)
}

func (p *DdosZoneTemplateHttpFilter) getPath() string{
    return "ddos/zone-template/http/" +p.Inst.HttpTmplName + "/filter"
}

func (p *DdosZoneTemplateHttpFilter) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateHttpFilter::Post")
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

func (p *DdosZoneTemplateHttpFilter) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateHttpFilter::Get")
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
func (p *DdosZoneTemplateHttpFilter) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateHttpFilter::Put")
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

func (p *DdosZoneTemplateHttpFilter) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneTemplateHttpFilter::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
