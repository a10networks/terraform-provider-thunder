

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type InterfaceEthernetLldp struct {
	Inst struct {

    EnableCfg InterfaceEthernetLldpEnableCfg `json:"enable-cfg"`
    NotificationCfg InterfaceEthernetLldpNotificationCfg `json:"notification-cfg"`
    TxDot1Cfg InterfaceEthernetLldpTxDot1Cfg `json:"tx-dot1-cfg"`
    TxTlvsCfg InterfaceEthernetLldpTxTlvsCfg `json:"tx-tlvs-cfg"`
    Uuid string `json:"uuid"`
    Ifnum string 

	} `json:"lldp"`
}


type InterfaceEthernetLldpEnableCfg struct {
    RtEnable int `json:"rt-enable"`
    Rx int `json:"rx"`
    Tx int `json:"tx"`
}


type InterfaceEthernetLldpNotificationCfg struct {
    Notification int `json:"notification"`
    NotifEnable int `json:"notif-enable"`
}


type InterfaceEthernetLldpTxDot1Cfg struct {
    TxDot1Tlvs int `json:"tx-dot1-tlvs"`
    LinkAggregation int `json:"link-aggregation"`
    Vlan int `json:"vlan"`
}


type InterfaceEthernetLldpTxTlvsCfg struct {
    TxTlvs int `json:"tx-tlvs"`
    Exclude int `json:"exclude"`
    ManagementAddress int `json:"management-address"`
    PortDescription int `json:"port-description"`
    SystemCapabilities int `json:"system-capabilities"`
    SystemDescription int `json:"system-description"`
    SystemName int `json:"system-name"`
}

func (p *InterfaceEthernetLldp) GetId() string{
    return "1"
}

func (p *InterfaceEthernetLldp) getPath() string{
    return "interface/ethernet/" +p.Inst.Ifnum + "/lldp"
}

func (p *InterfaceEthernetLldp) Post(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetLldp::Post")
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

func (p *InterfaceEthernetLldp) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetLldp::Get")
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
func (p *InterfaceEthernetLldp) Put(authToken string, host string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetLldp::Put")
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

func (p *InterfaceEthernetLldp) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
    logger.Println("InterfaceEthernetLldp::Delete")
    headers := axapi.GenRequestHeader(authToken)
        _, _, err := axapi.SendDelete(host, p.getPath(), "", nil, headers, logger)
    return err
}
