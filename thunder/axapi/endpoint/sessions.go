

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Sessions struct {
	Inst struct {

    Ext SessionsExt1397 `json:"ext"`
    Smp SessionsSmp1398 `json:"smp"`
    SmpTable SessionsSmpTable1399 `json:"smp-table"`
    Uuid string `json:"uuid"`

	} `json:"sessions"`
}


type SessionsExt1397 struct {
    Uuid string `json:"uuid"`
}


type SessionsSmp1398 struct {
    Uuid string `json:"uuid"`
}


type SessionsSmpTable1399 struct {
    Uuid string `json:"uuid"`
}

func (p *Sessions) GetId() string{
    return "1"
}

func (p *Sessions) getPath() string{
    return "sessions"
}

func (p *Sessions) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Sessions::Post")
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

func (p *Sessions) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Sessions::Get")
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
func (p *Sessions) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Sessions::Put")
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

func (p *Sessions) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Sessions::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
