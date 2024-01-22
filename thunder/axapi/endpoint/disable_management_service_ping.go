

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type DisableManagementServicePing struct {
	Inst struct {

    AllDataIntf int `json:"all-data-intf"`
    EthCfg []DisableManagementServicePingEthCfg `json:"eth-cfg"`
    Management int `json:"management"`
    TunnelCfg []DisableManagementServicePingTunnelCfg `json:"tunnel-cfg"`
    Uuid string `json:"uuid"`
    VeCfg []DisableManagementServicePingVeCfg `json:"ve-cfg"`

	} `json:"ping"`
}


type DisableManagementServicePingEthCfg struct {
    EthernetStart int `json:"ethernet-start"`
    EthernetEnd int `json:"ethernet-end"`
}


type DisableManagementServicePingTunnelCfg struct {
    TunnelStart int `json:"tunnel-start"`
    TunnelEnd int `json:"tunnel-end"`
}


type DisableManagementServicePingVeCfg struct {
    VeStart int `json:"ve-start"`
    VeEnd int `json:"ve-end"`
}

func (p *DisableManagementServicePing) GetId() string{
    return "1"
}

func (p *DisableManagementServicePing) getPath() string{
    return "disable-management/service/ping"
}

func (p *DisableManagementServicePing) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DisableManagementServicePing::Post")
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

func (p *DisableManagementServicePing) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DisableManagementServicePing::Get")
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
func (p *DisableManagementServicePing) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("DisableManagementServicePing::Put")
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

func (p *DisableManagementServicePing) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("DisableManagementServicePing::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
