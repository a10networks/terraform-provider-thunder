

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type TechreportPriorityPartition struct {
	Inst struct {

    PartName string `json:"part-name"`
    Uuid string `json:"uuid"`

	} `json:"priority-partition"`
}

func (p *TechreportPriorityPartition) GetId() string{
    return p.Inst.PartName
}

func (p *TechreportPriorityPartition) getPath() string{
    return "techreport/priority-partition"
}

func (p *TechreportPriorityPartition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TechreportPriorityPartition::Post")
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

func (p *TechreportPriorityPartition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TechreportPriorityPartition::Get")
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
func (p *TechreportPriorityPartition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("TechreportPriorityPartition::Put")
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

func (p *TechreportPriorityPartition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("TechreportPriorityPartition::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
