

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type GslbDns struct {
	Inst struct {

    Action string `json:"action" dval:"none"`
    Logging string `json:"logging" dval:"none"`
    SamplingEnable []GslbDnsSamplingEnable `json:"sampling-enable"`
    Template string `json:"template"`
    Uuid string `json:"uuid"`

	} `json:"dns"`
}


type GslbDnsSamplingEnable struct {
    Counters1 string `json:"counters1"`
}

func (p *GslbDns) GetId() string{
    return "1"
}

func (p *GslbDns) getPath() string{
    return "gslb/dns"
}

func (p *GslbDns) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbDns::Post")
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

func (p *GslbDns) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbDns::Get")
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
func (p *GslbDns) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("GslbDns::Put")
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

func (p *GslbDns) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("GslbDns::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
