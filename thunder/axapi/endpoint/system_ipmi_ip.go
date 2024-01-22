

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SystemIpmiIp struct {
	Inst struct {

    DefaultGateway string `json:"default-gateway"`
    Ipv4Address string `json:"ipv4-address"`
    Ipv4Netmask string `json:"ipv4-netmask"`

	} `json:"ip"`
}

func (p *SystemIpmiIp) GetId() string{
    return "1"
}

func (p *SystemIpmiIp) getPath() string{
    return "system/ipmi/ip"
}

func (p *SystemIpmiIp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpmiIp::Post")
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

func (p *SystemIpmiIp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpmiIp::Get")
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
func (p *SystemIpmiIp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpmiIp::Put")
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

func (p *SystemIpmiIp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SystemIpmiIp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
