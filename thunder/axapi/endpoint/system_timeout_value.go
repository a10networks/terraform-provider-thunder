

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemTimeoutValue struct {
	Inst struct {

    Ftp int `json:"ftp" dval:"120"`
    Http int `json:"http"`
    Https int `json:"https"`
    Scp int `json:"scp" dval:"300"`
    Sftp int `json:"sftp"`
    Tftp int `json:"tftp" dval:"300"`
    Uuid string `json:"uuid"`

	} `json:"timeout-value"`
}

func (p *SystemTimeoutValue) GetId() string{
    return "1"
}

func (p *SystemTimeoutValue) getPath() string{
    return "system/timeout-value"
}

func (p *SystemTimeoutValue) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTimeoutValue::Post")
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

func (p *SystemTimeoutValue) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTimeoutValue::Get")
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
func (p *SystemTimeoutValue) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTimeoutValue::Put")
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

func (p *SystemTimeoutValue) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemTimeoutValue::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
