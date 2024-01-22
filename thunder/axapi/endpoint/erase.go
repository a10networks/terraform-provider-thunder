

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Erase struct {
	Inst struct {

    AllPartitions int `json:"all-partitions"`
    Grubconfig int `json:"grubconfig"`
    PreserveAccounts int `json:"preserve-accounts"`
    PreserveManagement int `json:"preserve-management"`
    Reload int `json:"reload"`
    ServiceConfig int `json:"service-config"`

	} `json:"erase"`
}

func (p *Erase) GetId() string{
    return "1"
}

func (p *Erase) getPath() string{
    return "erase"
}

func (p *Erase) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Erase::Post")
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

func (p *Erase) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Erase::Get")
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
func (p *Erase) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Erase::Put")
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

func (p *Erase) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Erase::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
