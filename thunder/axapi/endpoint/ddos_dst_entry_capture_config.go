

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDstEntryCaptureConfig struct {
	Inst struct {

    Mode string `json:"mode"`
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    DstEntryName string 

	} `json:"capture-config"`
}

func (p *DdosDstEntryCaptureConfig) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosDstEntryCaptureConfig) getPath() string{
    return "ddos/dst/entry/" +p.Inst.DstEntryName + "/capture-config"
}

func (p *DdosDstEntryCaptureConfig) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryCaptureConfig::Post")
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

func (p *DdosDstEntryCaptureConfig) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryCaptureConfig::Get")
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
func (p *DdosDstEntryCaptureConfig) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryCaptureConfig::Put")
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

func (p *DdosDstEntryCaptureConfig) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstEntryCaptureConfig::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
