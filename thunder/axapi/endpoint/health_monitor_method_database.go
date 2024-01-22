

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorMethodDatabase struct {
	Inst struct {

    Database int `json:"database"`
    DatabaseName string `json:"database-name"`
    DbColumn int `json:"db-column"`
    DbColumnInteger int `json:"db-column-integer"`
    DbEncrypted string `json:"db-encrypted"`
    DbName string `json:"db-name"`
    DbPassword int `json:"db-password"`
    DbPasswordStr string `json:"db-password-str"`
    DbReceive string `json:"db-receive"`
    DbReceiveInteger int `json:"db-receive-integer"`
    DbRow int `json:"db-row"`
    DbRowInteger int `json:"db-row-integer"`
    DbSend string `json:"db-send"`
    DbUsername string `json:"db-username"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"database"`
}

func (p *HealthMonitorMethodDatabase) GetId() string{
    return "1"
}

func (p *HealthMonitorMethodDatabase) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/method/database"
}

func (p *HealthMonitorMethodDatabase) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodDatabase::Post")
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

func (p *HealthMonitorMethodDatabase) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodDatabase::Get")
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
func (p *HealthMonitorMethodDatabase) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodDatabase::Put")
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

func (p *HealthMonitorMethodDatabase) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorMethodDatabase::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
