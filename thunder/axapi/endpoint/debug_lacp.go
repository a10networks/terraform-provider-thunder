

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugLacp struct {
	Inst struct {

    All int `json:"all"`
    Cli int `json:"cli"`
    Detail int `json:"detail"`
    Event int `json:"event"`
    Ha int `json:"ha"`
    Packet int `json:"packet"`
    Sync int `json:"sync"`
    Timer int `json:"timer"`
    Uuid string `json:"uuid"`

	} `json:"lacp"`
}

func (p *DebugLacp) GetId() string{
    return "1"
}

func (p *DebugLacp) getPath() string{
    return "debug/lacp"
}

func (p *DebugLacp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugLacp::Post")
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

func (p *DebugLacp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugLacp::Get")
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
func (p *DebugLacp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugLacp::Put")
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

func (p *DebugLacp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugLacp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
