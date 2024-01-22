

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementServicePing struct {
	Inst struct {

    AclV4List []EnableManagementServicePingAclV4List `json:"acl-v4-list"`
    AclV6List []EnableManagementServicePingAclV6List `json:"acl-v6-list"`
    Uuid string `json:"uuid"`

	} `json:"ping"`
}


type EnableManagementServicePingAclV4List struct {
    AclId int `json:"acl-id"`
    EthCfg []EnableManagementServicePingAclV4ListEthCfg `json:"eth-cfg"`
    VeCfg []EnableManagementServicePingAclV4ListVeCfg `json:"ve-cfg"`
    TunnelCfg []EnableManagementServicePingAclV4ListTunnelCfg `json:"tunnel-cfg"`
    Management int `json:"management"`
    AllDataIntf int `json:"all-data-intf"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type EnableManagementServicePingAclV4ListEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServicePingAclV4ListVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type EnableManagementServicePingAclV4ListTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServicePingAclV6List struct {
    AclName string `json:"acl-name"`
    EthCfg []EnableManagementServicePingAclV6ListEthCfg `json:"eth-cfg"`
    VeCfg []EnableManagementServicePingAclV6ListVeCfg `json:"ve-cfg"`
    TunnelCfg []EnableManagementServicePingAclV6ListTunnelCfg `json:"tunnel-cfg"`
    Management int `json:"management"`
    AllDataIntf int `json:"all-data-intf"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type EnableManagementServicePingAclV6ListEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServicePingAclV6ListVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type EnableManagementServicePingAclV6ListTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}

func (p *EnableManagementServicePing) GetId() string{
    return "1"
}

func (p *EnableManagementServicePing) getPath() string{
    return "enable-management/service/ping"
}

func (p *EnableManagementServicePing) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServicePing::Post")
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

func (p *EnableManagementServicePing) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServicePing::Get")
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
func (p *EnableManagementServicePing) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServicePing::Put")
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

func (p *EnableManagementServicePing) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServicePing::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
