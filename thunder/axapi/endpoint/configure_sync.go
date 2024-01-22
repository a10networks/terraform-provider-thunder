

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ConfigureSync struct {
	Inst struct {

    Address string `json:"address"`
    AllPartitions int `json:"all-partitions"`
    AutoAuthentication int `json:"auto-authentication"`
    PartitionName string `json:"partition-name"`
    PrivateKey string `json:"private-key"`
    Pwd string `json:"pwd"`
    PwdEnc string `json:"pwd-enc"`
    Shared int `json:"shared"`
    Type string `json:"type"`
    Usr string `json:"usr"`

	} `json:"sync"`
}

func (p *ConfigureSync) GetId() string{
    return "1"
}

func (p *ConfigureSync) getPath() string{
    return "configure/sync"
}

func (p *ConfigureSync) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ConfigureSync::Post")
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

func (p *ConfigureSync) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ConfigureSync::Get")
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
func (p *ConfigureSync) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ConfigureSync::Put")
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

func (p *ConfigureSync) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ConfigureSync::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
