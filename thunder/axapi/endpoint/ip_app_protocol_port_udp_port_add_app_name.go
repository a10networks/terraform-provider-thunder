

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "net/url"
)

//based on ACOS 6_0_2_P1-37
type IpAppProtocolPortUdpPortAddAppName struct {
	Inst struct {

    Interface IpAppProtocolPortUdpPortAddAppNameInterface1036 `json:"interface"`
    Name string `json:"name"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    Port string 

	} `json:"app-name"`
}


type IpAppProtocolPortUdpPortAddAppNameInterface1036 struct {
    Management int `json:"management"`
    VeCfg []IpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg1037 `json:"ve-cfg"`
    EthCfg []IpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg1038 `json:"eth-cfg"`
    Uuid string `json:"uuid"`
}


type IpAppProtocolPortUdpPortAddAppNameInterfaceVeCfg1037 struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type IpAppProtocolPortUdpPortAddAppNameInterfaceEthCfg1038 struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}

func (p *IpAppProtocolPortUdpPortAddAppName) GetId() string{
    return url.QueryEscape(p.Inst.Name)
}

func (p *IpAppProtocolPortUdpPortAddAppName) getPath() string{
    return "ip/app-protocol-port/udp/port/add/" +p.Inst.Port + "/app-name"
}

func (p *IpAppProtocolPortUdpPortAddAppName) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortAddAppName::Post")
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

func (p *IpAppProtocolPortUdpPortAddAppName) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortAddAppName::Get")
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
func (p *IpAppProtocolPortUdpPortAddAppName) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortAddAppName::Put")
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

func (p *IpAppProtocolPortUdpPortAddAppName) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("IpAppProtocolPortUdpPortAddAppName::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
