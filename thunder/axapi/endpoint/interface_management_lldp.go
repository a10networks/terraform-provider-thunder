

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceManagementLldp struct {
	Inst struct {

    EnableCfg InterfaceManagementLldpEnableCfg `json:"enable-cfg"`
    NotificationCfg InterfaceManagementLldpNotificationCfg `json:"notification-cfg"`
    TxDot1Cfg InterfaceManagementLldpTxDot1Cfg `json:"tx-dot1-cfg"`
    TxTlvsCfg InterfaceManagementLldpTxTlvsCfg `json:"tx-tlvs-cfg"`
    Uuid string `json:"uuid"`

	} `json:"lldp"`
}


type InterfaceManagementLldpEnableCfg struct {
    RtEnable int `json:"rt-enable"`
    Rx int `json:"rx"`
    Tx int `json:"tx"`
}


type InterfaceManagementLldpNotificationCfg struct {
    Notification int `json:"notification"`
    NotifEnable int `json:"notif-enable"`
}


type InterfaceManagementLldpTxDot1Cfg struct {
    TxDot1Tlvs int `json:"tx-dot1-tlvs"`
    LinkAggregation int `json:"link-aggregation"`
    Vlan int `json:"vlan"`
}


type InterfaceManagementLldpTxTlvsCfg struct {
    TxTlvs int `json:"tx-tlvs"`
    Exclude int `json:"exclude"`
    ManagementAddress int `json:"management-address"`
    PortDescription int `json:"port-description"`
    SystemCapabilities int `json:"system-capabilities"`
    SystemDescription int `json:"system-description"`
    SystemName int `json:"system-name"`
}

func (p *InterfaceManagementLldp) GetId() string{
    return "1"
}

func (p *InterfaceManagementLldp) getPath() string{
    return "interface/management/lldp"
}

func (p *InterfaceManagementLldp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceManagementLldp::Post")
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

func (p *InterfaceManagementLldp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceManagementLldp::Get")
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
func (p *InterfaceManagementLldp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceManagementLldp::Put")
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

func (p *InterfaceManagementLldp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceManagementLldp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
