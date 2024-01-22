

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbActiveRdt struct {
	Inst struct {

    Domain string `json:"domain"`
    Icmp int `json:"icmp"`
    Interval int `json:"interval" dval:"1"`
    Port int `json:"port"`
    Retry int `json:"retry" dval:"3"`
    Sleep int `json:"sleep" dval:"3"`
    Timeout int `json:"timeout" dval:"3000"`
    Track int `json:"track" dval:"60"`
    Uuid string `json:"uuid"`

	} `json:"active-rdt"`
}

func (p *GslbActiveRdt) GetId() string{
    return "1"
}

func (p *GslbActiveRdt) getPath() string{
    return "gslb/active-rdt"
}

func (p *GslbActiveRdt) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbActiveRdt::Post")
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

func (p *GslbActiveRdt) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbActiveRdt::Get")
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
func (p *GslbActiveRdt) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbActiveRdt::Put")
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

func (p *GslbActiveRdt) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbActiveRdt::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
