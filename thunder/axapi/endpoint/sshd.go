

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type Sshd struct {
	Inst struct {

    FileUrl string `json:"file-url"`
    Generate int `json:"generate"`
    Load int `json:"load"`
    ReAddRsa string `json:"re-add-rsa"`
    Regenerate int `json:"regenerate"`
    Restart int `json:"restart"`
    Size string `json:"size"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Wipe int `json:"wipe"`

	} `json:"sshd"`
}

func (p *Sshd) GetId() string{
    return "1"
}

func (p *Sshd) getPath() string{
    return "sshd"
}

func (p *Sshd) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Sshd::Post")
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

func (p *Sshd) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Sshd::Get")
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
func (p *Sshd) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("Sshd::Put")
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

func (p *Sshd) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("Sshd::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
