

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type ImportPeriodicClassList struct {
	Inst struct {

    Background int `json:"background"`
    ClassList string `json:"class-list"`
    Period int `json:"period"`
    RemoteFile string `json:"remote-file"`
    UseMgmtPort int `json:"use-mgmt-port"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"class-list"`
}

func (p *ImportPeriodicClassList) GetId() string{
    return p.Inst.ClassList
}

func (p *ImportPeriodicClassList) getPath() string{
    return "import-periodic/class-list"
}

func (p *ImportPeriodicClassList) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicClassList::Post")
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

func (p *ImportPeriodicClassList) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicClassList::Get")
    headers := axapi.GenRequestHeader(authToken)
        _, axResp, err := axapi.SendGet(host, p.getPath(), instId, nil, headers, logger)
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
func (p *ImportPeriodicClassList) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicClassList::Put")
    headers := axapi.GenRequestHeader(authToken)
    payloadBytes, err := axapi.SerializeToJson(p)
    if err != nil {
        logger.Println("Failed to serialize struct as json", err)
        return err
    }
    logger.Println("payload: " + string(payloadBytes))
    _, _, err = axapi.SendPut(host, p.getPath(), p.GetId(), payloadBytes, headers, logger)
    return err
}

func (p *ImportPeriodicClassList) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("ImportPeriodicClassList::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
