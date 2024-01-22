

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NtpNtpGlobal struct {
	Inst struct {

    AllowDataPorts int `json:"allow-data-ports"`
    Uuid string `json:"uuid"`

	} `json:"ntp-global"`
}

func (p *NtpNtpGlobal) GetId() string{
    return "1"
}

func (p *NtpNtpGlobal) getPath() string{
    return "ntp/ntp-global"
}

func (p *NtpNtpGlobal) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NtpNtpGlobal::Post")
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

func (p *NtpNtpGlobal) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NtpNtpGlobal::Get")
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
func (p *NtpNtpGlobal) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NtpNtpGlobal::Put")
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

func (p *NtpNtpGlobal) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NtpNtpGlobal::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
