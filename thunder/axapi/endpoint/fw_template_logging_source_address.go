

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type FwTemplateLoggingSourceAddress struct {
	Inst struct {

    Ip string `json:"ip"`
    Ipv6 string `json:"ipv6"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"source-address"`
}

func (p *FwTemplateLoggingSourceAddress) GetId() string{
    return "1"
}

func (p *FwTemplateLoggingSourceAddress) getPath() string{
    return "fw/template/logging/" +p.Inst.Name + "/source-address"
}

func (p *FwTemplateLoggingSourceAddress) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwTemplateLoggingSourceAddress::Post")
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

func (p *FwTemplateLoggingSourceAddress) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwTemplateLoggingSourceAddress::Get")
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
func (p *FwTemplateLoggingSourceAddress) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("FwTemplateLoggingSourceAddress::Put")
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

func (p *FwTemplateLoggingSourceAddress) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("FwTemplateLoggingSourceAddress::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
