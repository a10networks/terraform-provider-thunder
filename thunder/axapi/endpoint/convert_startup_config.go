

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ConvertStartupConfig struct {
	Inst struct {

    DestProfile string `json:"dest-profile"`
    Profile string `json:"profile"`
    Type string `json:"type"`

	} `json:"convert-startup-config"`
}

func (p *ConvertStartupConfig) GetId() string{
    return "1"
}

func (p *ConvertStartupConfig) getPath() string{
    return "convert-startup-config"
}

func (p *ConvertStartupConfig) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ConvertStartupConfig::Post")
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

func (p *ConvertStartupConfig) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ConvertStartupConfig::Get")
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
func (p *ConvertStartupConfig) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ConvertStartupConfig::Put")
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

func (p *ConvertStartupConfig) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ConvertStartupConfig::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
