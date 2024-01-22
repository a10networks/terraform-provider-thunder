

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosSrcPortTemplateDns struct {
	Inst struct {

    Name string `json:"name"`
    QueryResolutionCheck DdosSrcPortTemplateDnsQueryResolutionCheck294 `json:"query-resolution-check"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"dns"`
}


type DdosSrcPortTemplateDnsQueryResolutionCheck294 struct {
    SessionTimeoutValue int `json:"session-timeout-value"`
    DomainLockupAction string `json:"domain-lockup-action" dval:"default"`
    BigResponseSize int `json:"big-response-size"`
    BigResponseAction string `json:"big-response-action" dval:"default"`
    Uuid string `json:"uuid"`
}

func (p *DdosSrcPortTemplateDns) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosSrcPortTemplateDns) getPath() string{
    return "ddos/src-port-template/dns"
}

func (p *DdosSrcPortTemplateDns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcPortTemplateDns::Post")
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

func (p *DdosSrcPortTemplateDns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcPortTemplateDns::Get")
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
func (p *DdosSrcPortTemplateDns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcPortTemplateDns::Put")
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

func (p *DdosSrcPortTemplateDns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosSrcPortTemplateDns::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
