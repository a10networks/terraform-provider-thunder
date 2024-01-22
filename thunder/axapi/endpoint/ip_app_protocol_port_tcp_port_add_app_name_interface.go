

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type IpAppProtocolPortTcpPortAddAppNameInterface struct {
	Inst struct {

    EthCfg []IpAppProtocolPortTcpPortAddAppNameInterfaceEthCfg `json:"eth-cfg"`
    Management int `json:"management"`
    Uuid string `json:"uuid"`
    VeCfg []IpAppProtocolPortTcpPortAddAppNameInterfaceVeCfg `json:"ve-cfg"`
    Name string 
    Port string 

	} `json:"interface"`
}


type IpAppProtocolPortTcpPortAddAppNameInterfaceEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type IpAppProtocolPortTcpPortAddAppNameInterfaceVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}

func (p *IpAppProtocolPortTcpPortAddAppNameInterface) GetId() string{
    return "1"
}

func (p *IpAppProtocolPortTcpPortAddAppNameInterface) getPath() string{
    return "ip/app-protocol-port/tcp/port/add/" +p.Inst.Port + "/app-name/" +p.Inst.Name + "/interface"
}

func (p *IpAppProtocolPortTcpPortAddAppNameInterface) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortTcpPortAddAppNameInterface::Post")
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

func (p *IpAppProtocolPortTcpPortAddAppNameInterface) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortTcpPortAddAppNameInterface::Get")
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
func (p *IpAppProtocolPortTcpPortAddAppNameInterface) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortTcpPortAddAppNameInterface::Put")
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

func (p *IpAppProtocolPortTcpPortAddAppNameInterface) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortTcpPortAddAppNameInterface::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
