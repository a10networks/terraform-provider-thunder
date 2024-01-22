

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type MirrorPort struct {
	Inst struct {

    Ethernet int `json:"ethernet"`
    MirrorDir string `json:"mirror-dir" dval:"both"`
    MirrorIndex int `json:"mirror-index"`
    Uuid string `json:"uuid"`

	} `json:"mirror-port"`
}

func (p *MirrorPort) GetId() string{
    return strconv.Itoa(p.Inst.MirrorIndex)
}

func (p *MirrorPort) getPath() string{
    return "mirror-port"
}

func (p *MirrorPort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("MirrorPort::Post")
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

func (p *MirrorPort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("MirrorPort::Get")
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
func (p *MirrorPort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("MirrorPort::Put")
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

func (p *MirrorPort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("MirrorPort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
