

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementServiceNtp struct {
	Inst struct {

    AclV4List []EnableManagementServiceNtpAclV4List `json:"acl-v4-list"`
    AclV6List []EnableManagementServiceNtpAclV6List `json:"acl-v6-list"`
    Uuid string `json:"uuid"`

	} `json:"ntp"`
}


type EnableManagementServiceNtpAclV4List struct {
    AclId int `json:"acl-id"`
    EthCfg []EnableManagementServiceNtpAclV4ListEthCfg `json:"eth-cfg"`
    VeCfg []EnableManagementServiceNtpAclV4ListVeCfg `json:"ve-cfg"`
    TunnelCfg []EnableManagementServiceNtpAclV4ListTunnelCfg `json:"tunnel-cfg"`
    Management int `json:"management"`
    AllDataIntf int `json:"all-data-intf"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type EnableManagementServiceNtpAclV4ListEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceNtpAclV4ListVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type EnableManagementServiceNtpAclV4ListTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServiceNtpAclV6List struct {
    AclName string `json:"acl-name"`
    EthCfg []EnableManagementServiceNtpAclV6ListEthCfg `json:"eth-cfg"`
    VeCfg []EnableManagementServiceNtpAclV6ListVeCfg `json:"ve-cfg"`
    TunnelCfg []EnableManagementServiceNtpAclV6ListTunnelCfg `json:"tunnel-cfg"`
    Management int `json:"management"`
    AllDataIntf int `json:"all-data-intf"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type EnableManagementServiceNtpAclV6ListEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceNtpAclV6ListVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type EnableManagementServiceNtpAclV6ListTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}

func (p *EnableManagementServiceNtp) GetId() string{
    return "1"
}

func (p *EnableManagementServiceNtp) getPath() string{
    return "enable-management/service/ntp"
}

func (p *EnableManagementServiceNtp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceNtp::Post")
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

func (p *EnableManagementServiceNtp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceNtp::Get")
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
func (p *EnableManagementServiceNtp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceNtp::Put")
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

func (p *EnableManagementServiceNtp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceNtp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
