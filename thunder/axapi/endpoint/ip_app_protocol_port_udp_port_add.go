

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type IpAppProtocolPortUdpPortAdd struct {
	Inst struct {

    AppNameList []IpAppProtocolPortUdpPortAddAppNameList `json:"app-name-list"`
    Port int `json:"port"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`

	} `json:"add"`
}


type IpAppProtocolPortUdpPortAddAppNameList struct {
    Name string `json:"name"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
    Interface IpAppProtocolPortUdpPortAddAppNameListInterface `json:"interface"`
}


type IpAppProtocolPortUdpPortAddAppNameListInterface struct {
    Management int `json:"management"`
    VeCfg []IpAppProtocolPortUdpPortAddAppNameListInterfaceVeCfg `json:"ve-cfg"`
    EthCfg []IpAppProtocolPortUdpPortAddAppNameListInterfaceEthCfg `json:"eth-cfg"`
    Uuid string `json:"uuid"`
}


type IpAppProtocolPortUdpPortAddAppNameListInterfaceVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type IpAppProtocolPortUdpPortAddAppNameListInterfaceEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}

func (p *IpAppProtocolPortUdpPortAdd) GetId() string{
    return strconv.Itoa(p.Inst.Port)
}

func (p *IpAppProtocolPortUdpPortAdd) getPath() string{
    return "ip/app-protocol-port/udp/port/add"
}

func (p *IpAppProtocolPortUdpPortAdd) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortAdd::Post")
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

func (p *IpAppProtocolPortUdpPortAdd) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortAdd::Get")
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
func (p *IpAppProtocolPortUdpPortAdd) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortAdd::Put")
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

func (p *IpAppProtocolPortUdpPortAdd) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortAdd::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
