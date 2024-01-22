

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Accounting struct {
	Inst struct {

    Commands int `json:"commands"`
    Debug int `json:"debug"`
    Exec AccountingExec40 `json:"exec"`
    StopOnly int `json:"stop-only"`
    Tacplus int `json:"tacplus"`
    Uuid string `json:"uuid"`

	} `json:"accounting"`
}


type AccountingExec40 struct {
    AccountingExecType string `json:"accounting-exec-type"`
    AccountingExecMethod string `json:"accounting-exec-method"`
    Uuid string `json:"uuid"`
}

func (p *Accounting) GetId() string{
    return "1"
}

func (p *Accounting) getPath() string{
    return "accounting"
}

func (p *Accounting) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Accounting::Post")
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

func (p *Accounting) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Accounting::Get")
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
func (p *Accounting) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Accounting::Put")
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

func (p *Accounting) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Accounting::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
