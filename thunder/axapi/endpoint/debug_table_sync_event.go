

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugTableSyncEvent struct {
	Inst struct {

    All int `json:"all"`
    Arp int `json:"arp"`
    Fibv4 int `json:"fibv4"`
    Fibv6 int `json:"fibv6"`
    Mac int `json:"mac"`
    Nd6 int `json:"nd6"`
    Uuid string `json:"uuid"`

	} `json:"table-sync-event"`
}

func (p *DebugTableSyncEvent) GetId() string{
    return "1"
}

func (p *DebugTableSyncEvent) getPath() string{
    return "debug/table-sync-event"
}

func (p *DebugTableSyncEvent) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugTableSyncEvent::Post")
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

func (p *DebugTableSyncEvent) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugTableSyncEvent::Get")
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
func (p *DebugTableSyncEvent) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugTableSyncEvent::Put")
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

func (p *DebugTableSyncEvent) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugTableSyncEvent::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
