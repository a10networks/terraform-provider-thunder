

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbProtocolLimit struct {
	Inst struct {

    ArdtQuery int `json:"ardt-query" dval:"200"`
    ArdtResponse int `json:"ardt-response" dval:"1000"`
    ArdtSession int `json:"ardt-session" dval:"32768"`
    ConnResponse int `json:"conn-response"`
    Message int `json:"message" dval:"10000"`
    Response int `json:"response" dval:"3600"`
    Uuid string `json:"uuid"`

	} `json:"limit"`
}

func (p *GslbProtocolLimit) GetId() string{
    return "1"
}

func (p *GslbProtocolLimit) getPath() string{
    return "gslb/protocol/limit"
}

func (p *GslbProtocolLimit) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbProtocolLimit::Post")
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

func (p *GslbProtocolLimit) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbProtocolLimit::Get")
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
func (p *GslbProtocolLimit) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbProtocolLimit::Put")
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

func (p *GslbProtocolLimit) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbProtocolLimit::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
