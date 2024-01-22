

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateConnectionReuse struct {
	Inst struct {

    AddHeader int `json:"add-header"`
    KeepAliveConn int `json:"keep-alive-conn"`
    LimitPerServer int `json:"limit-per-server" dval:"1000"`
    Name string `json:"name"`
    NumConnPerPort int `json:"num-conn-per-port" dval:"100"`
    Preopen int `json:"preopen"`
    Timeout int `json:"timeout" dval:"2400"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"connection-reuse"`
}

func (p *SlbTemplateConnectionReuse) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateConnectionReuse) getPath() string{
    return "slb/template/connection-reuse"
}

func (p *SlbTemplateConnectionReuse) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateConnectionReuse::Post")
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

func (p *SlbTemplateConnectionReuse) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateConnectionReuse::Get")
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
func (p *SlbTemplateConnectionReuse) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateConnectionReuse::Put")
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

func (p *SlbTemplateConnectionReuse) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateConnectionReuse::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
