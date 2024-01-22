

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type HealthMonitorHeaderInsert struct {
	Inst struct {

    InsertList []HealthMonitorHeaderInsertInsertList `json:"insert-list"`
    Uuid string `json:"uuid"`
    Name string 

	} `json:"header-insert"`
}


type HealthMonitorHeaderInsertInsertList struct {
    InsertContent string `json:"insert-content"`
}

func (p *HealthMonitorHeaderInsert) GetId() string{
    return "1"
}

func (p *HealthMonitorHeaderInsert) getPath() string{
    return "health/monitor/" +p.Inst.Name + "/header-insert"
}

func (p *HealthMonitorHeaderInsert) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorHeaderInsert::Post")
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

func (p *HealthMonitorHeaderInsert) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorHeaderInsert::Get")
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
func (p *HealthMonitorHeaderInsert) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorHeaderInsert::Put")
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

func (p *HealthMonitorHeaderInsert) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("HealthMonitorHeaderInsert::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
