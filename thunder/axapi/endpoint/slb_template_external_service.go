

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateExternalService struct {
	Inst struct {

    Action string `json:"action" dval:"continue"`
    BypassIpCfg []SlbTemplateExternalServiceBypassIpCfg `json:"bypass-ip-cfg"`
    FailureAction string `json:"failure-action" dval:"continue"`
    Name string `json:"name"`
    RequestHeaderForwardList []SlbTemplateExternalServiceRequestHeaderForwardList `json:"request-header-forward-list"`
    ServiceGroup string `json:"service-group"`
    SharedPartitionPersistSourceIpTemplate int `json:"shared-partition-persist-source-ip-template"`
    SharedPartitionTcpProxyTemplate int `json:"shared-partition-tcp-proxy-template"`
    SourceIp string `json:"source-ip"`
    TcpProxy string `json:"tcp-proxy"`
    TemplatePersistSourceIpShared string `json:"template-persist-source-ip-shared"`
    TemplateTcpProxyShared string `json:"template-tcp-proxy-shared"`
    Timeout int `json:"timeout" dval:"5"`
    Type string `json:"type" dval:"url-filter"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"external-service"`
}


type SlbTemplateExternalServiceBypassIpCfg struct {
    BypassIp string `json:"bypass-ip"`
    Mask string `json:"mask"`
}


type SlbTemplateExternalServiceRequestHeaderForwardList struct {
    RequestHeaderForward string `json:"request-header-forward"`
}

func (p *SlbTemplateExternalService) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateExternalService) getPath() string{
    return "slb/template/external-service"
}

func (p *SlbTemplateExternalService) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateExternalService::Post")
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

func (p *SlbTemplateExternalService) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateExternalService::Get")
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
func (p *SlbTemplateExternalService) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateExternalService::Put")
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

func (p *SlbTemplateExternalService) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateExternalService::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
