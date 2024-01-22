

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck struct {
	Inst struct {

    GwhcNsCacheLookup string `json:"gwhc-ns-cache-lookup" dval:"disabled"`
    Interval int `json:"interval" dval:"10"`
    NumQueryType int `json:"num-query-type"`
    QueryName string `json:"query-name" dval:"a10networks.com"`
    Retry int `json:"retry" dval:"6"`
    RetryMulti int `json:"retry-multi" dval:"1"`
    StrQueryType string `json:"str-query-type" dval:"A"`
    Timeout int `json:"timeout" dval:"5"`
    UpRetry int `json:"up-retry" dval:"1"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"gateway-health-check"`
}

func (p *SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck) GetId() string{
    return "1"
}

func (p *SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck) getPath() string{
    return "slb/template/dns/" +p.Inst.Name + "/recursive-dns-resolution/gateway-health-check"
}

func (p *SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck::Post")
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

func (p *SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck::Get")
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
func (p *SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck::Put")
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

func (p *SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateDnsRecursiveDnsResolutionGatewayHealthCheck::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
