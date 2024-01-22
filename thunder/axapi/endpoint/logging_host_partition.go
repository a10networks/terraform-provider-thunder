

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingHostPartition struct {
	Inst struct {

    PartitionName string `json:"partition-name"`
    Shared int `json:"shared"`
    Uuid string `json:"uuid"`

	} `json:"partition"`
}

func (p *LoggingHostPartition) GetId() string{
    return "1"
}

func (p *LoggingHostPartition) getPath() string{
    return "logging/host/partition"
}

func (p *LoggingHostPartition) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingHostPartition::Post")
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

func (p *LoggingHostPartition) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingHostPartition::Get")
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
func (p *LoggingHostPartition) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingHostPartition::Put")
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

func (p *LoggingHostPartition) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingHostPartition::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
