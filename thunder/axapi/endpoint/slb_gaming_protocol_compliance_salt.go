

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SlbGamingProtocolComplianceSalt struct {
	Inst struct {

    CurrentSalt int `json:"current-salt"`
    PreviousSalt int `json:"previous-salt"`
    Uuid string `json:"uuid"`

	} `json:"gaming-protocol-compliance-salt"`
}

func (p *SlbGamingProtocolComplianceSalt) GetId() string{
    return "1"
}

func (p *SlbGamingProtocolComplianceSalt) getPath() string{
    return "slb/gaming-protocol-compliance-salt"
}

func (p *SlbGamingProtocolComplianceSalt) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbGamingProtocolComplianceSalt::Post")
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

func (p *SlbGamingProtocolComplianceSalt) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbGamingProtocolComplianceSalt::Get")
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
func (p *SlbGamingProtocolComplianceSalt) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbGamingProtocolComplianceSalt::Put")
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

func (p *SlbGamingProtocolComplianceSalt) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbGamingProtocolComplianceSalt::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
