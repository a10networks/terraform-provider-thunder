

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LinkStartupConfig struct {
	Inst struct {

    AllPartitions int `json:"all-partitions"`
    Default int `json:"default"`
    FileName string `json:"file-name"`
    Primary int `json:"primary"`
    Secondary int `json:"secondary"`

	} `json:"startup-config"`
}

func (p *LinkStartupConfig) GetId() string{
    return "1"
}

func (p *LinkStartupConfig) getPath() string{
    return "link/startup-config"
}

func (p *LinkStartupConfig) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LinkStartupConfig::Post")
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

func (p *LinkStartupConfig) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LinkStartupConfig::Get")
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
func (p *LinkStartupConfig) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LinkStartupConfig::Put")
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

func (p *LinkStartupConfig) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LinkStartupConfig::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
