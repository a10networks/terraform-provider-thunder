

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerHostIpv4Host struct {
	Inst struct {

    Ipv4Addr string `json:"ipv4-addr"`
    UdpPort int `json:"udp-port" dval:"162"`
    User string `json:"user"`
    Uuid string `json:"uuid"`
    V1V2cComm string `json:"v1-v2c-comm"`
    Version string `json:"version"`

	} `json:"ipv4-host"`
}

func (p *SnmpServerHostIpv4Host) GetId() string{
    return p.Inst.Ipv4Addr+"+"+p.Inst.Version
}

func (p *SnmpServerHostIpv4Host) getPath() string{
    return "snmp-server/host/ipv4-host"
}

func (p *SnmpServerHostIpv4Host) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerHostIpv4Host::Post")
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

func (p *SnmpServerHostIpv4Host) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerHostIpv4Host::Get")
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
func (p *SnmpServerHostIpv4Host) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerHostIpv4Host::Put")
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

func (p *SnmpServerHostIpv4Host) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerHostIpv4Host::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
