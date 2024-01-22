

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type LoggingHostIpv4addr struct {
	Inst struct {

    HostIpv4 string `json:"host-ipv4"`
    OverTls int `json:"over-tls"`
    Port int `json:"port" dval:"514"`
    Tcp int `json:"tcp"`
    UseMgmtPort int `json:"use-mgmt-port"`
    Uuid string `json:"uuid"`

	} `json:"ipv4addr"`
}

func (p *LoggingHostIpv4addr) GetId() string{
    return p.Inst.HostIpv4
}

func (p *LoggingHostIpv4addr) getPath() string{
    return "logging/host/ipv4addr"
}

func (p *LoggingHostIpv4addr) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingHostIpv4addr::Post")
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

func (p *LoggingHostIpv4addr) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingHostIpv4addr::Get")
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
func (p *LoggingHostIpv4addr) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingHostIpv4addr::Put")
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

func (p *LoggingHostIpv4addr) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("LoggingHostIpv4addr::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
