

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type NetflowTemplate struct {
	Inst struct {

    InformationElementBlk []NetflowTemplateInformationElementBlk `json:"information-element-blk"`
    IpfixTemplateId int `json:"ipfix-template-id"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"template"`
}


type NetflowTemplateInformationElementBlk struct {
    InformationElement string `json:"information-element"`
}

func (p *NetflowTemplate) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *NetflowTemplate) getPath() string{
    return "netflow/template"
}

func (p *NetflowTemplate) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowTemplate::Post")
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

func (p *NetflowTemplate) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowTemplate::Get")
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
func (p *NetflowTemplate) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowTemplate::Put")
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

func (p *NetflowTemplate) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowTemplate::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
