

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementServicePingAclV4 struct {
	Inst struct {

    AclId int `json:"acl-id"`
    AllDataIntf int `json:"all-data-intf"`
    EthCfg []EnableManagementServicePingAclV4EthCfg `json:"eth-cfg"`
    Management int `json:"management"`
    TunnelCfg []EnableManagementServicePingAclV4TunnelCfg `json:"tunnel-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VeCfg []EnableManagementServicePingAclV4VeCfg `json:"ve-cfg"`

	} `json:"acl-v4"`
}


type EnableManagementServicePingAclV4EthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServicePingAclV4TunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServicePingAclV4VeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}

func (p *EnableManagementServicePingAclV4) GetId() string{
    return strconv.Itoa(p.Inst.AclId)
}

func (p *EnableManagementServicePingAclV4) getPath() string{
    return "enable-management/service/ping/acl-v4"
}

func (p *EnableManagementServicePingAclV4) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServicePingAclV4::Post")
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

func (p *EnableManagementServicePingAclV4) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServicePingAclV4::Get")
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
func (p *EnableManagementServicePingAclV4) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServicePingAclV4::Put")
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

func (p *EnableManagementServicePingAclV4) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServicePingAclV4::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
