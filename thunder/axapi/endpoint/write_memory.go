

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type WriteMemory struct {
	Inst struct {

    Destination string `json:"destination"`
    Partition string `json:"partition"`
    Profile string `json:"profile"`
    SpecifiedPartition string `json:"specified-partition"`

	} `json:"memory"`
}

func (p *WriteMemory) GetId() string{
    return "1"
}

func (p *WriteMemory) getPath() string{
    return "write/memory"
}

func (p *WriteMemory) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WriteMemory::Post")
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

func (p *WriteMemory) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WriteMemory::Get")
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
func (p *WriteMemory) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("WriteMemory::Put")
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

func (p *WriteMemory) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("WriteMemory::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
