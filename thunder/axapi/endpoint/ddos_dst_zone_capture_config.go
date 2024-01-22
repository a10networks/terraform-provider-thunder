

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type DdosDstZoneCaptureConfig struct {
	Inst struct {

    Mode string `json:"mode"`
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    ZoneName string 

	} `json:"capture-config"`
}

func (p *DdosDstZoneCaptureConfig) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *DdosDstZoneCaptureConfig) getPath() string{
    return "ddos/dst/zone/" +p.Inst.ZoneName + "/capture-config"
}

func (p *DdosDstZoneCaptureConfig) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneCaptureConfig::Post")
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

func (p *DdosDstZoneCaptureConfig) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneCaptureConfig::Get")
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
func (p *DdosDstZoneCaptureConfig) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneCaptureConfig::Put")
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

func (p *DdosDstZoneCaptureConfig) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DdosDstZoneCaptureConfig::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
