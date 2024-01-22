

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Rename struct {
	Inst struct {

    InstanceName string `json:"instance-name"`
    NewInstanceName string `json:"new-instance-name"`
    Object string `json:"object"`

	} `json:"rename"`
}

func (p *Rename) GetId() string{
    return "1"
}

func (p *Rename) getPath() string{
    return "rename"
}

func (p *Rename) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Rename::Post")
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

func (p *Rename) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Rename::Get")
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
func (p *Rename) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Rename::Put")
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

func (p *Rename) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Rename::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
