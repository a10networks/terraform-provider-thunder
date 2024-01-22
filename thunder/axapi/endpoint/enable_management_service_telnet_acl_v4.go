

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
    "strconv"
)

//based on ACOS 6_0_2_P1-37
type EnableManagementServiceTelnetAclV4 struct {
	Inst struct {

    AclId int `json:"acl-id"`
    AllDataIntf int `json:"all-data-intf"`
    EthCfg []EnableManagementServiceTelnetAclV4EthCfg `json:"eth-cfg"`
    Management int `json:"management"`
    TunnelCfg []EnableManagementServiceTelnetAclV4TunnelCfg `json:"tunnel-cfg"`
    UserTag string `json:"user-tag"`
    Uuid string `json:"uuid"`
    VeCfg []EnableManagementServiceTelnetAclV4VeCfg `json:"ve-cfg"`

	} `json:"acl-v4"`
}


type EnableManagementServiceTelnetAclV4EthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type EnableManagementServiceTelnetAclV4TunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type EnableManagementServiceTelnetAclV4VeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}

func (p *EnableManagementServiceTelnetAclV4) GetId() string{
    return strconv.Itoa(p.Inst.AclId)
}

func (p *EnableManagementServiceTelnetAclV4) getPath() string{
    return "enable-management/service/telnet/acl-v4"
}

func (p *EnableManagementServiceTelnetAclV4) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceTelnetAclV4::Post")
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

func (p *EnableManagementServiceTelnetAclV4) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceTelnetAclV4::Get")
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
func (p *EnableManagementServiceTelnetAclV4) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceTelnetAclV4::Put")
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

func (p *EnableManagementServiceTelnetAclV4) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("EnableManagementServiceTelnetAclV4::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), instId, nil, headers, logger)
    return err
}
