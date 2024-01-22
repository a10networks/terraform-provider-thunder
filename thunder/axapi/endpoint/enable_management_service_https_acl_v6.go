

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementServiceHttpsAclV6 struct {
	Inst struct {

    AclName string `json:"acl-name"`
    AllDataIntf int `json:"all-data-intf"`
    EthCfg []EnableManagementServiceHttpsAclV6EthCfg `json:"eth-cfg"`
    Management int `json:"management"`
    TunnelCfg []EnableManagementServiceHttpsAclV6TunnelCfg `json:"tunnel-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VeCfg []EnableManagementServiceHttpsAclV6VeCfg `json:"ve-cfg"`

	} `json:"acl-v6"`
}


type EnableManagementServiceHttpsAclV6EthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceHttpsAclV6TunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServiceHttpsAclV6VeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}

func (p *EnableManagementServiceHttpsAclV6) GetId() string{
    return p.Inst.AclName
}

func (p *EnableManagementServiceHttpsAclV6) getPath() string{
    return "enable-management/service/https/acl-v6"
}

func (p *EnableManagementServiceHttpsAclV6) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceHttpsAclV6::Post")
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

func (p *EnableManagementServiceHttpsAclV6) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceHttpsAclV6::Get")
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
func (p *EnableManagementServiceHttpsAclV6) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceHttpsAclV6::Put")
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

func (p *EnableManagementServiceHttpsAclV6) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceHttpsAclV6::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
