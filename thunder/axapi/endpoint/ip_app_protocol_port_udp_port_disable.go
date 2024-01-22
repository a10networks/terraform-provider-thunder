

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type IpAppProtocolPortUdpPortDisable struct {
	Inst struct {

    Interface IpAppProtocolPortUdpPortDisableInterface1039 `json:"interface"`
    Port int `json:"port"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"disable"`
}


type IpAppProtocolPortUdpPortDisableInterface1039 struct {
    Management int `json:"management"`
    VeCfg []IpAppProtocolPortUdpPortDisableInterfaceVeCfg1040 `json:"ve-cfg"`
    EthCfg []IpAppProtocolPortUdpPortDisableInterfaceEthCfg1041 `json:"eth-cfg"`
    Uuid string `json:"uuid"`
}


type IpAppProtocolPortUdpPortDisableInterfaceVeCfg1040 struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type IpAppProtocolPortUdpPortDisableInterfaceEthCfg1041 struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}

func (p *IpAppProtocolPortUdpPortDisable) GetId() string{
    return strconv.Itoa(p.Inst.Port)
}

func (p *IpAppProtocolPortUdpPortDisable) getPath() string{
    return "ip/app-protocol-port/udp/port/disable"
}

func (p *IpAppProtocolPortUdpPortDisable) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortDisable::Post")
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

func (p *IpAppProtocolPortUdpPortDisable) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortDisable::Get")
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
func (p *IpAppProtocolPortUdpPortDisable) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortDisable::Put")
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

func (p *IpAppProtocolPortUdpPortDisable) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortDisable::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
