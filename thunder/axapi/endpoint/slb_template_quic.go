

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type SlbTemplateQuic struct {
	Inst struct {

    BurstLen int `json:"burst-len"`
    ConnectionIdLength int `json:"connection-id-length"`
    IdleTimeout int `json:"idle-timeout"`
    InitialWnd int `json:"initial-wnd"`
    KeyUpdateToClient int `json:"key-update-to-client"`
    KeyUpdateToServer int `json:"key-update-to-server"`
    Name string `json:"name"`
    ReceiveBuffer int `json:"receive-buffer"`
    ServerRetry int `json:"server-retry"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"quic"`
}

func (p *SlbTemplateQuic) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *SlbTemplateQuic) getPath() string{
    return "slb/template/quic"
}

func (p *SlbTemplateQuic) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateQuic::Post")
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

func (p *SlbTemplateQuic) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateQuic::Get")
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
func (p *SlbTemplateQuic) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateQuic::Put")
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

func (p *SlbTemplateQuic) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbTemplateQuic::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
