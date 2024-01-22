

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DebugQuic struct {
	Inst struct {

    Dummy int `json:"dummy"`
    Level int `json:"level"`
    Tls int `json:"tls"`
    TlsLevel int `json:"tls-level"`
    Uuid string `json:"uuid"`

	} `json:"quic"`
}

func (p *DebugQuic) GetId() string{
    return "1"
}

func (p *DebugQuic) getPath() string{
    return "debug/quic"
}

func (p *DebugQuic) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugQuic::Post")
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

func (p *DebugQuic) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugQuic::Get")
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
func (p *DebugQuic) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DebugQuic::Put")
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

func (p *DebugQuic) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DebugQuic::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
