

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type NtpServerHostname struct {
	Inst struct {

    Action string `json:"action" dval:"enable"`
    HostServername string `json:"host-servername"`
    Key int `json:"key"`
    Prefer int `json:"prefer"`
    Uuid string `json:"uuid"`

	} `json:"hostname"`
}

func (p *NtpServerHostname) GetId() string{
    return p.Inst.HostServername
}

func (p *NtpServerHostname) getPath() string{
    return "ntp/server/hostname"
}

func (p *NtpServerHostname) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NtpServerHostname::Post")
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

func (p *NtpServerHostname) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NtpServerHostname::Get")
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
func (p *NtpServerHostname) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("NtpServerHostname::Put")
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

func (p *NtpServerHostname) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("NtpServerHostname::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
