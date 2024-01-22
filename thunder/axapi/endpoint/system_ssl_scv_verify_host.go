

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemSslScvVerifyHost struct {
	Inst struct {

    Disable int `json:"disable"`
    Uuid string `json:"uuid"`

	} `json:"ssl-scv-verify-host"`
}

func (p *SystemSslScvVerifyHost) GetId() string{
    return "1"
}

func (p *SystemSslScvVerifyHost) getPath() string{
    return "system/ssl-scv-verify-host"
}

func (p *SystemSslScvVerifyHost) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSslScvVerifyHost::Post")
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

func (p *SystemSslScvVerifyHost) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSslScvVerifyHost::Get")
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
func (p *SystemSslScvVerifyHost) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSslScvVerifyHost::Put")
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

func (p *SystemSslScvVerifyHost) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemSslScvVerifyHost::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
