

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type SlbGamingProtocolCompliancePort struct {
	Inst struct {

    Port int `json:"port"`
    Range int `json:"range"`
    Udp int `json:"udp"`
    Uuid string `json:"uuid"`

	} `json:"gaming-protocol-compliance-port"`
}

func (p *SlbGamingProtocolCompliancePort) GetId() string{
    return strconv.Itoa(p.Inst.Port)
}

func (p *SlbGamingProtocolCompliancePort) getPath() string{
    return "slb/gaming-protocol-compliance-port"
}

func (p *SlbGamingProtocolCompliancePort) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbGamingProtocolCompliancePort::Post")
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

func (p *SlbGamingProtocolCompliancePort) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbGamingProtocolCompliancePort::Get")
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
func (p *SlbGamingProtocolCompliancePort) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SlbGamingProtocolCompliancePort::Put")
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

func (p *SlbGamingProtocolCompliancePort) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SlbGamingProtocolCompliancePort::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
