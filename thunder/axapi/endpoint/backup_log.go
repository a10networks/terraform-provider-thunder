

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type BackupLog struct {
	Inst struct {

    All int `json:"all"`
    Core int `json:"core"`
    Date int `json:"date"`
    Day int `json:"day"`
    Expedite int `json:"expedite"`
    Messages int `json:"messages"`
    Month int `json:"month"`
    Password string `json:"password"`
    Period int `json:"period"`
    RemoteFile string `json:"remote-file"`
    Showtech int `json:"showtech"`
    StatsData int `json:"stats-data"`
    StoreName string `json:"store-name"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Week int `json:"week"`

	} `json:"log"`
}

func (p *BackupLog) GetId() string{
    return "1"
}

func (p *BackupLog) getPath() string{
    return "backup/log"
}

func (p *BackupLog) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("BackupLog::Post")
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

func (p *BackupLog) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("BackupLog::Get")
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
func (p *BackupLog) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("BackupLog::Put")
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

func (p *BackupLog) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("BackupLog::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
