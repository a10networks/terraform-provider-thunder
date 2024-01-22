

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosZoneSrcPortTemplateDns struct {
	Inst struct {

    Name string `json:"name"`
    QueryResolutionCheck DdosZoneSrcPortTemplateDnsQueryResolutionCheck306 `json:"query-resolution-check"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"dns"`
}


type DdosZoneSrcPortTemplateDnsQueryResolutionCheck306 struct {
    SessionTimeoutValue int `json:"session-timeout-value"`
    DomainLockupAction string `json:"domain-lockup-action" dval:"default"`
    BigResponseSize int `json:"big-response-size"`
    BigResponseAction string `json:"big-response-action" dval:"default"`
    Uuid string `json:"uuid"`
}

func (p *DdosZoneSrcPortTemplateDns) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosZoneSrcPortTemplateDns) getPath() string{
    return "ddos/zone-src-port-template/dns"
}

func (p *DdosZoneSrcPortTemplateDns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateDns::Post")
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

func (p *DdosZoneSrcPortTemplateDns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateDns::Get")
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
func (p *DdosZoneSrcPortTemplateDns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateDns::Put")
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

func (p *DdosZoneSrcPortTemplateDns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosZoneSrcPortTemplateDns::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
