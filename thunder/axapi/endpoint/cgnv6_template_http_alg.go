

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type Cgnv6TemplateHttpAlg struct {
	Inst struct {

    HeaderNameClientIp string `json:"header-name-client-ip" dval:"X-Forwarded-For"`
    HeaderNameMsisdn string `json:"header-name-msisdn" dval:"X-MSISDN"`
    IncludeTunnelIp int `json:"include-tunnel-ip"`
    Method string `json:"method" dval:"append"`
    Name string `json:"name"`
    RequestInsertClientIp int `json:"request-insert-client-ip"`
    RequestInsertMsisdn int `json:"request-insert-msisdn"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"http-alg"`
}

func (p *Cgnv6TemplateHttpAlg) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *Cgnv6TemplateHttpAlg) getPath() string{
    return "cgnv6/template/http-alg"
}

func (p *Cgnv6TemplateHttpAlg) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateHttpAlg::Post")
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

func (p *Cgnv6TemplateHttpAlg) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateHttpAlg::Get")
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
func (p *Cgnv6TemplateHttpAlg) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateHttpAlg::Put")
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

func (p *Cgnv6TemplateHttpAlg) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Cgnv6TemplateHttpAlg::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
