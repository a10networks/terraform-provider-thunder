

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type IpAppProtocolPortTcpPortDisable struct {
	Inst struct {

    Interface IpAppProtocolPortTcpPortDisableInterface1033 `json:"interface"`
    Port int `json:"port"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"disable"`
}


type IpAppProtocolPortTcpPortDisableInterface1033 struct {
    Management int `json:"management"`
    VeCfg []IpAppProtocolPortTcpPortDisableInterfaceVeCfg1034 `json:"ve-cfg"`
    EthCfg []IpAppProtocolPortTcpPortDisableInterfaceEthCfg1035 `json:"eth-cfg"`
    Uuid string `json:"uuid"`
}


type IpAppProtocolPortTcpPortDisableInterfaceVeCfg1034 struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type IpAppProtocolPortTcpPortDisableInterfaceEthCfg1035 struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}

func (p *IpAppProtocolPortTcpPortDisable) GetId() string{
    return strconv.Itoa(p.Inst.Port)
}

func (p *IpAppProtocolPortTcpPortDisable) getPath() string{
    return "ip/app-protocol-port/tcp/port/disable"
}

func (p *IpAppProtocolPortTcpPortDisable) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortTcpPortDisable::Post")
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

func (p *IpAppProtocolPortTcpPortDisable) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortTcpPortDisable::Get")
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
func (p *IpAppProtocolPortTcpPortDisable) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortTcpPortDisable::Put")
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

func (p *IpAppProtocolPortTcpPortDisable) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortTcpPortDisable::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
