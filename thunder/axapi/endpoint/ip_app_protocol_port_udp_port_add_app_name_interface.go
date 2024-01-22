

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpAppProtocolPortUdpPortAddAppNameInterface struct {
	Inst struct {

    EthCfg []IpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg `json:"eth-cfg"`
    Management int `json:"management"`
    Uuid string `json:"uuid"`
    VeCfg []IpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg `json:"ve-cfg"`
    Name string 
    Port string 

	} `json:"interface"`
}


type IpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type IpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}

func (p *IpAppProtocolPortUdpPortAddAppNameInterface) GetId() string{
    return "1"
}

func (p *IpAppProtocolPortUdpPortAddAppNameInterface) getPath() string{
    return "ip/app-protocol-port/udp/port/add/" +p.Inst.Port + "/app-name/" +p.Inst.Name + "/interface"
}

func (p *IpAppProtocolPortUdpPortAddAppNameInterface) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortAddAppNameInterface::Post")
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

func (p *IpAppProtocolPortUdpPortAddAppNameInterface) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortAddAppNameInterface::Get")
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
func (p *IpAppProtocolPortUdpPortAddAppNameInterface) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortAddAppNameInterface::Put")
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

func (p *IpAppProtocolPortUdpPortAddAppNameInterface) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortAddAppNameInterface::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
