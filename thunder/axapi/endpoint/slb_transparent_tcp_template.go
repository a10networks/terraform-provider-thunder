

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbTransparentTcpTemplate struct {
	Inst struct {

    Name string `json:"name" dval:"default"`
    Uuid string `json:"uuid"`

	} `json:"transparent-tcp-template"`
}

func (p *SlbTransparentTcpTemplate) GetId() string{
    return "1"
}

func (p *SlbTransparentTcpTemplate) getPath() string{
    return "slb/transparent-tcp-template"
}

func (p *SlbTransparentTcpTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTransparentTcpTemplate::Post")
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

func (p *SlbTransparentTcpTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTransparentTcpTemplate::Get")
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
func (p *SlbTransparentTcpTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTransparentTcpTemplate::Put")
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

func (p *SlbTransparentTcpTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTransparentTcpTemplate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
