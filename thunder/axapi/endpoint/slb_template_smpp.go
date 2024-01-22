

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateSmpp struct {
	Inst struct {

    ClientEnquireLink int `json:"client-enquire-link"`
    Name string `json:"name"`
    Password string `json:"password"`
    ServerEnquireLink int `json:"server-enquire-link"`
    ServerEnquireLinkVal int `json:"server-enquire-link-val" dval:"30"`
    ServerSelectionPerRequest int `json:"server-selection-per-request"`
    User string `json:"user"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"smpp"`
}

func (p *SlbTemplateSmpp) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateSmpp) getPath() string{
    return "slb/template/smpp"
}

func (p *SlbTemplateSmpp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateSmpp::Post")
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

func (p *SlbTemplateSmpp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateSmpp::Get")
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
func (p *SlbTemplateSmpp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateSmpp::Put")
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

func (p *SlbTemplateSmpp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateSmpp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
