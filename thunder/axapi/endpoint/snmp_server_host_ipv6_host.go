

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type SnmpServerHostIpv6Host struct {
	Inst struct {

    Ipv6Addr string `json:"ipv6-addr"`
    UdpPort int `json:"udp-port" dval:"162"`
    User string `json:"user"`
    Uuid string `json:"uuid"`
    V1V2cComm string `json:"v1-v2c-comm"`
    Version string `json:"version"`

	} `json:"ipv6-host"`
}

func (p *SnmpServerHostIpv6Host) GetId() string{
    return p.Inst.Ipv6Addr+"+"+p.Inst.Version
}

func (p *SnmpServerHostIpv6Host) getPath() string{
    return "snmp-server/host/ipv6-host"
}

func (p *SnmpServerHostIpv6Host) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerHostIpv6Host::Post")
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

func (p *SnmpServerHostIpv6Host) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerHostIpv6Host::Get")
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
func (p *SnmpServerHostIpv6Host) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerHostIpv6Host::Put")
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

func (p *SnmpServerHostIpv6Host) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("SnmpServerHostIpv6Host::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
