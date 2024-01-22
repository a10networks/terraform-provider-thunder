

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementServiceHttp struct {
	Inst struct {

    AclV4List []EnableManagementServiceHttpAclV4List `json:"acl-v4-list"`
    AclV6List []EnableManagementServiceHttpAclV6List `json:"acl-v6-list"`
    AllDataIntf int `json:"all-data-intf"`
    EthCfg []EnableManagementServiceHttpEthCfg `json:"eth-cfg"`
    TunnelCfg []EnableManagementServiceHttpTunnelCfg `json:"tunnel-cfg"`
    Uuid string `json:"uuid"`
    VeCfg []EnableManagementServiceHttpVeCfg `json:"ve-cfg"`

	} `json:"http"`
}


type EnableManagementServiceHttpAclV4List struct {
    AclId int `json:"acl-id"`
    EthCfg []EnableManagementServiceHttpAclV4ListEthCfg `json:"eth-cfg"`
    VeCfg []EnableManagementServiceHttpAclV4ListVeCfg `json:"ve-cfg"`
    TunnelCfg []EnableManagementServiceHttpAclV4ListTunnelCfg `json:"tunnel-cfg"`
    Management int `json:"management"`
    AllDataIntf int `json:"all-data-intf"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type EnableManagementServiceHttpAclV4ListEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceHttpAclV4ListVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type EnableManagementServiceHttpAclV4ListTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServiceHttpAclV6List struct {
    AclName string `json:"acl-name"`
    EthCfg []EnableManagementServiceHttpAclV6ListEthCfg `json:"eth-cfg"`
    VeCfg []EnableManagementServiceHttpAclV6ListVeCfg `json:"ve-cfg"`
    TunnelCfg []EnableManagementServiceHttpAclV6ListTunnelCfg `json:"tunnel-cfg"`
    Management int `json:"management"`
    AllDataIntf int `json:"all-data-intf"`
    Uuid string `json:"uuid"`
    UserTag string `json:"user-tag"`
}


type EnableManagementServiceHttpAclV6ListEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceHttpAclV6ListVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}


type EnableManagementServiceHttpAclV6ListTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServiceHttpEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceHttpTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServiceHttpVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}

func (p *EnableManagementServiceHttp) GetId() string{
    return "1"
}

func (p *EnableManagementServiceHttp) getPath() string{
    return "enable-management/service/http"
}

func (p *EnableManagementServiceHttp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceHttp::Post")
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

func (p *EnableManagementServiceHttp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceHttp::Get")
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
func (p *EnableManagementServiceHttp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceHttp::Put")
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

func (p *EnableManagementServiceHttp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceHttp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
