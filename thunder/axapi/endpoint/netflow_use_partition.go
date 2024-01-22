

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NetflowUsePartition struct {
	Inst struct {

    PartName string `json:"part-name"`
    Uuid string `json:"uuid"`

	} `json:"use-partition"`
}

func (p *NetflowUsePartition) GetId() string{
    return "1"
}

func (p *NetflowUsePartition) getPath() string{
    return "netflow/use-partition"
}

func (p *NetflowUsePartition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowUsePartition::Post")
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

func (p *NetflowUsePartition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowUsePartition::Get")
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
func (p *NetflowUsePartition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowUsePartition::Put")
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

func (p *NetflowUsePartition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NetflowUsePartition::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
