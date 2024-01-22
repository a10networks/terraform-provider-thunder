

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Hostname struct {
	Inst struct {

    Uuid string `json:"uuid"`
    Value string `json:"value"`

	} `json:"hostname"`
}

func (p *Hostname) GetId() string{
    return "1"
}

func (p *Hostname) getPath() string{
    return "hostname"
}

func (p *Hostname) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Hostname::Post")
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

func (p *Hostname) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Hostname::Get")
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
func (p *Hostname) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Hostname::Put")
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

func (p *Hostname) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Hostname::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
