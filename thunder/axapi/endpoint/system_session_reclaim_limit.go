

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemSessionReclaimLimit struct {
	Inst struct {

    NscanLimit int `json:"nscan-limit" dval:"4096"`
    ScanFreq int `json:"scan-freq" dval:"5"`
    Uuid string `json:"uuid"`

	} `json:"session-reclaim-limit"`
}

func (p *SystemSessionReclaimLimit) GetId() string{
    return "1"
}

func (p *SystemSessionReclaimLimit) getPath() string{
    return "system/session-reclaim-limit"
}

func (p *SystemSessionReclaimLimit) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSessionReclaimLimit::Post")
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

func (p *SystemSessionReclaimLimit) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSessionReclaimLimit::Get")
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
func (p *SystemSessionReclaimLimit) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSessionReclaimLimit::Put")
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

func (p *SystemSessionReclaimLimit) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSessionReclaimLimit::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
