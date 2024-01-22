

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementServiceSnmp struct {
	Inst struct {

    AclV4List []EnableManagementServiceSnmpAclV4List `json:"acl-v4-list"`
    AclV6List []EnableManagementServiceSnmpAclV6List `json:"acl-v6-list"`
    AllDataIntf int `json:"all-data-intf"`
    EthCfg []EnableManagementServiceSnmpEthCfg `json:"eth-cfg"`
    TunnelCfg []EnableManagementServiceSnmpTunnelCfg `json:"tunnel-cfg"`
    Uuid string `json:"uuid"`
    VeCfg []EnableManagementServiceSnmpVeCfg `json:"ve-cfg"`

	} `json:"snmp"`
}


type EnableManagementServiceSnmpAclV4List struct {
    AclId int `json:"acl-id"`
    EthCfg []EnableManagementServiceSnmpAclV4ListEthCfg `json:"eth-cfg"`
    VeCfg []EnableManagementServiceSnmpAclV4ListVeCfg `json:"ve-cfg"`
    TunnelCfg []EnableManagementServiceSnmpAclV4ListTunnelCfg `json:"tunnel-cfg"`
    Management int `json:"management"`
    AllDataIntf int `json:"all-data-intf"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type EnableManagementServiceSnmpAclV4ListEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceSnmpAclV4ListVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type EnableManagementServiceSnmpAclV4ListTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServiceSnmpAclV6List struct {
    AclName string `json:"acl-name"`
    EthCfg []EnableManagementServiceSnmpAclV6ListEthCfg `json:"eth-cfg"`
    VeCfg []EnableManagementServiceSnmpAclV6ListVeCfg `json:"ve-cfg"`
    TunnelCfg []EnableManagementServiceSnmpAclV6ListTunnelCfg `json:"tunnel-cfg"`
    Management int `json:"management"`
    AllDataIntf int `json:"all-data-intf"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type EnableManagementServiceSnmpAclV6ListEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceSnmpAclV6ListVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type EnableManagementServiceSnmpAclV6ListTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServiceSnmpEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceSnmpTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServiceSnmpVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}

func (p *EnableManagementServiceSnmp) GetId() string{
    return "1"
}

func (p *EnableManagementServiceSnmp) getPath() string{
    return "enable-management/service/snmp"
}

func (p *EnableManagementServiceSnmp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceSnmp::Post")
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

func (p *EnableManagementServiceSnmp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceSnmp::Get")
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
func (p *EnableManagementServiceSnmp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceSnmp::Put")
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

func (p *EnableManagementServiceSnmp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceSnmp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
