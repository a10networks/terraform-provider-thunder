

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementServiceTelnet struct {
	Inst struct {

    AclV4List []EnableManagementServiceTelnetAclV4List `json:"acl-v4-list"`
    AclV6List []EnableManagementServiceTelnetAclV6List `json:"acl-v6-list"`
    AllDataIntf int `json:"all-data-intf"`
    EthCfg []EnableManagementServiceTelnetEthCfg `json:"eth-cfg"`
    Management int `json:"management"`
    TunnelCfg []EnableManagementServiceTelnetTunnelCfg `json:"tunnel-cfg"`
    Uuid string `json:"uuid"`
    VeCfg []EnableManagementServiceTelnetVeCfg `json:"ve-cfg"`

	} `json:"telnet"`
}


type EnableManagementServiceTelnetAclV4List struct {
    AclId int `json:"acl-id"`
    EthCfg []EnableManagementServiceTelnetAclV4ListEthCfg `json:"eth-cfg"`
    VeCfg []EnableManagementServiceTelnetAclV4ListVeCfg `json:"ve-cfg"`
    TunnelCfg []EnableManagementServiceTelnetAclV4ListTunnelCfg `json:"tunnel-cfg"`
    Management int `json:"management"`
    AllDataIntf int `json:"all-data-intf"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type EnableManagementServiceTelnetAclV4ListEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceTelnetAclV4ListVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type EnableManagementServiceTelnetAclV4ListTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServiceTelnetAclV6List struct {
    AclName string `json:"acl-name"`
    EthCfg []EnableManagementServiceTelnetAclV6ListEthCfg `json:"eth-cfg"`
    VeCfg []EnableManagementServiceTelnetAclV6ListVeCfg `json:"ve-cfg"`
    TunnelCfg []EnableManagementServiceTelnetAclV6ListTunnelCfg `json:"tunnel-cfg"`
    Management int `json:"management"`
    AllDataIntf int `json:"all-data-intf"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type EnableManagementServiceTelnetAclV6ListEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceTelnetAclV6ListVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type EnableManagementServiceTelnetAclV6ListTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServiceTelnetEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceTelnetTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServiceTelnetVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}

func (p *EnableManagementServiceTelnet) GetId() string{
    return "1"
}

func (p *EnableManagementServiceTelnet) getPath() string{
    return "enable-management/service/telnet"
}

func (p *EnableManagementServiceTelnet) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceTelnet::Post")
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

func (p *EnableManagementServiceTelnet) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceTelnet::Get")
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
func (p *EnableManagementServiceTelnet) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceTelnet::Put")
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

func (p *EnableManagementServiceTelnet) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceTelnet::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
