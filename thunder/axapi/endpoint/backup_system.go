

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type BackupSystem struct {
	Inst struct {

    Encrypt int `json:"encrypt"`
    Password string `json:"password"`
    RemoteFile string `json:"remote-file"`
    StoreName string `json:"store-name"`
    UseMgmtPort int `json:"use-mgmt-port"`

	} `json:"system"`
}

func (p *BackupSystem) GetId() string{
    return "1"
}

func (p *BackupSystem) getPath() string{
    return "backup/system"
}

func (p *BackupSystem) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("BackupSystem::Post")
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

func (p *BackupSystem) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("BackupSystem::Get")
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
func (p *BackupSystem) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("BackupSystem::Put")
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

func (p *BackupSystem) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("BackupSystem::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
