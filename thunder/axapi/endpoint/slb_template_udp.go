

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateUdp struct {
	Inst struct {

    Age int `json:"age"`
    Avp string `json:"avp"`
    DisableClearSession int `json:"disable-clear-session"`
    IdleTimeout int `json:"idle-timeout" dval:"120"`
    Immediate int `json:"immediate"`
    Name string `json:"name"`
    Qos int `json:"qos"`
    RadiusLbMethodHashType string `json:"radius-lb-method-hash-type"`
    ReSelectIfServerDown int `json:"re-select-if-server-down"`
    Short int `json:"short"`
    StatelessConnTimeout int `json:"stateless-conn-timeout" dval:"120"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    V6avp string `json:"v6avp"`

	} `json:"udp"`
}

func (p *SlbTemplateUdp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateUdp) getPath() string{
    return "slb/template/udp"
}

func (p *SlbTemplateUdp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateUdp::Post")
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

func (p *SlbTemplateUdp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateUdp::Get")
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
func (p *SlbTemplateUdp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateUdp::Put")
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

func (p *SlbTemplateUdp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateUdp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
