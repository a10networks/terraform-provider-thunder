

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RemoveUpgradeLock struct {
	Inst struct {

    File int `json:"file"`
    Uuid string `json:"uuid"`
    FileContent []byte `json:"-"`
    FileHandle string `json:"file-handle"`

	} `json:"remove-upgrade-lock"`
}

func (p *RemoveUpgradeLock) GetId() string{
    return "1"
}

func (p *RemoveUpgradeLock) getPath() string{
    return "remove-upgrade-lock"
}

func (p *RemoveUpgradeLock) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RemoveUpgradeLock::Post")
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

func (p *RemoveUpgradeLock) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RemoveUpgradeLock::Get")
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
func (p *RemoveUpgradeLock) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("RemoveUpgradeLock::Put")
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

func (p *RemoveUpgradeLock) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("RemoveUpgradeLock::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
