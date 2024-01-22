

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DdosTemplateSslL4SslTrafficCheck struct {
	Inst struct {

    CheckResumedConnection int `json:"check-resumed-connection"`
    HeaderAction string `json:"header-action"`
    HeaderInspection int `json:"header-inspection"`
    Uuid string `json:"uuid"`
    SslL4TmplName string 

	} `json:"ssl-traffic-check"`
}

func (p *DdosTemplateSslL4SslTrafficCheck) GetId() string{
    return "1"
}

func (p *DdosTemplateSslL4SslTrafficCheck) getPath() string{
    return "ddos/template/ssl-l4/" +p.Inst.SslL4TmplName + "/ssl-traffic-check"
}

func (p *DdosTemplateSslL4SslTrafficCheck) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSslL4SslTrafficCheck::Post")
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

func (p *DdosTemplateSslL4SslTrafficCheck) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSslL4SslTrafficCheck::Get")
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
func (p *DdosTemplateSslL4SslTrafficCheck) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSslL4SslTrafficCheck::Put")
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

func (p *DdosTemplateSslL4SslTrafficCheck) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosTemplateSslL4SslTrafficCheck::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
